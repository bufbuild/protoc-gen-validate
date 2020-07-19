using System;
using System.Text;

namespace Envoyproxy.Validator
{
    // A URL represents a parsed URL (technically, a URI reference).
    //
    // The general form represented is:
    //
    //	[scheme:][//[userinfo@]hostname[:port]][/]path[?query][#fragment]
    //
    // URLs that do not start with a slash after the scheme are interpreted as:
    //
    //	scheme:opaque[?query][#fragment]
    //
    // Note that the Path field is stored in decoded form: /%47%6f%2f becomes /Go/.
    // A consequence is that it is impossible to tell which slashes in the Path were
    // slashes in the raw URL and which were %2f. This distinction is rarely important,
    // but when it is, the code should use RawPath, an optional field which only gets
    // set if the default encoding is different from Path.
    //
    // URL's String method uses the EscapedPath method to obtain the path. See the
    // EscapedPath method for more details.
    public class Url
    {
        public string Scheme { get; set; }

        // encoded opaque data
        public string Opaque { get; set; }

        // username and password information
        public UserInfo User { get; set; }

        public string Hostname { get; set; }

        public string Port { get; set; }

        // path (relative paths may omit leading slash)
        public string Path { get; set; }

        // encoded query values, without '?'
        public string RawQuery { get; set; }

        // fragment for references, without '#'
        public string Fragment { get; set; }
    }

    public class UserInfo
    {
        public string Username { get; set; }
        public string Password { get; set; }
    }

    /// Re-Implementation of Golang's net/url/Parse
    /// (https://golang.org/src/net/url/url.go)
    public static class UrlParser
    {
        public static Url Parse(string value)
        {
            return Parse(value.AsSpan());
        }

        public static Url Parse(ReadOnlySpan<char> value)
        {
            var url = new Url();
            int i;

            // Cut off #fragment
            i = value.IndexOf('#');
            if (i >= 0)
            {
                url.Fragment = Unescape(value.Slice(i + 1), Mode.Fragment);
                value = value.Slice(0, i);
            }
            // else value remains unchanged

            if (ContainsCTLByte(value))
                throw new ArgumentException("invalid control character in URL");
            if (value.IsEmpty)
                throw new ArgumentException("empty URL");

            if (value.Length == 1 && value[0] == '*')
            {
                url.Path = "*";
                return url;
            }

            // Split off possible leading "http:", "mailto:", etc.
        	// Cannot contain escaped characters.
            url.Scheme = GetScheme(ref value);

            // Cut off ?query
            i = value.IndexOf('?');
            if (i >= 0)
            {
                url.RawQuery = value.Slice(i + 1).ToString();
                value = value.Slice(0, i);
            }
            // else value remains unchanged

            if (value.Length != 0 && value[0] != '/')
            {
                if (url.Scheme != null)
                {
                    // We consider rootless paths per RFC 3986 as opaque.
                    url.Opaque = value.ToString();
                    return url;
                }

                // Avoid confusion with malformed schemes, like cache_object:foo/bar.
                // See golang.org/issue/16822.
                //
                // RFC 3986, §3.3:
                // In addition, a URI reference (Section 4.1) may be a relative-path reference,
                // in which case the first path segment cannot contain a colon (":") character.
                var colon = value.IndexOf(':');
                var slash = value.IndexOf('/');

                if (colon >= 0 && (slash < 0 || colon < slash))
                    throw new ArgumentException("first path segment in URL cannot contain colon");
            }

            if ((url.Scheme != null || !value.StartsWith("///".AsSpan())) && value.StartsWith("//".AsSpan()))
            {
                value = value.Slice(2);

                ReadOnlySpan<char> authority;

                i = value.IndexOf('/');
                if (i >= 0)
                {
                    authority = value.Slice(0, i);
                    value = value.Slice(i);
                }
                else
                {
                    authority = value;
                    value = ReadOnlySpan<char>.Empty;
                }

                if (authority.Length > 0)
                    (url.User, url.Hostname, url.Port) = ParseAuthority(authority);
            }

            if (value.Length > 0)
                url.Path = Unescape(value, Mode.Path);

            return url;
        }

        // Maybe rawurl is of the form scheme:path.
        // (Scheme must be [a-zA-Z][a-zA-Z0-9+-.]*)
        // If so, return scheme (and modify _value); else null.
        private static string GetScheme(ref ReadOnlySpan<char> value)
        {
            for (int i = 0; i < value.Length; ++i)
            {
                var c = value[i];

                if ('a' <= c && c <= 'z' || 'A' <= c && c <= 'Z')
                {
                    // do nothing
                }
                else if ('0' <= c && c <= '9' || c == '+' || c == '-' || c == '.')
                {
                    if (i == 0)
                        return null;
                }
                else if (c == ':')
                {
                    if (i == 0)
                        throw new ArgumentException("missing protocol scheme");

                    var scheme = value.Slice(0, i).ToString().ToLower();
                    value = value.Slice(i + 1);
                    return scheme;
                }
                else
                {
                    // we have encountered an invalid character,
                    // so there is no valid scheme
                    return null;
                }
            }

            return null;
        }

        private static (UserInfo, string, string) ParseAuthority(ReadOnlySpan<char> value)
        {
            string host, port;
            int i;

            i = value.LastIndexOf('@');
            if (i < 0)
            {
                (host, port) = ParseHost(value);
                return (null, host, port);
            }

            (host, port) = ParseHost(value.Slice(i + 1));
            value = value.Slice(0, i);

            if (!ValidUserInfo(value))
                throw new ArgumentException("invalid userinfo");

            UserInfo user;

            i = value.IndexOf(':');
            if (i < 0)
            {
                user = new UserInfo
                {
                    Username = Unescape(value, Mode.UsernamePassword)
                };
            }
            else
            {
                user = new UserInfo
                {
                    Username = Unescape(value.Slice(0, i), Mode.UsernamePassword),
                    Password = Unescape(value.Slice(i + 1), Mode.UsernamePassword)
                };
            }

            return (user, host, port);
        }

        // ParseHost parses host as an authority without user
        // information. That is, as (host, port) from host[:port].
        private static (string, string) ParseHost(ReadOnlySpan<char> value)
        {
            string host;
            string port;

            int i;

            if (value.Length != 0 && value[0] == '[')
            {
                // Parse an IP-Literal in RFC 3986 and RFC 6874.
                // E.g., "[fe80::1]", "[fe80::1%25en0]", "[fe80::1]:80".
                i = value.LastIndexOf(']');
                if (i < 0)
                    throw new ArgumentException("missing ']' in host");

                var colonPort = value.Slice(i + 1);
                if (!ValidOptionalPort(colonPort))
                    throw new ArgumentException($"invalid port \"{colonPort.ToString()}\" after host");


                // RFC 6874 defines that %25 (%-encoded percent) introduces
                // the zone identifier, and the zone identifier can use basically
                // any %-encoding it likes. That's different from the host, which
                // can only %-encode non-ASCII bytes.
                // We do impose some restrictions on the zone, to avoid stupidity
                // like newlines.
                var zone = value.IndexOf("%25".AsSpan());
                if (zone >= 0)
                {
                    var host1 = Unescape(value.Slice(1, zone - 1), Mode.Host);
                    var host2 = Unescape(value.Slice(zone, i - zone), Mode.Zone);
                    host = $"{host1}{host2}";
                }
                else
                {
                    host = Unescape(value.Slice(1, i - 1), Mode.Host);
                }

                if (value.Length >= i + 2)
                    port = value.Slice(i + 2).ToString();  // already validated (contains only digits)
                else
                    port = null;
            }
            else
            {
                i = value.LastIndexOf(':');
                if (i >= 0)
                {
                    var colonPort = value.Slice(i);
                    if (!ValidOptionalPort(colonPort))
                        throw new ArgumentException($"invalid port \"{colonPort.ToString()}\" after host");

                    host = Unescape(value.Slice(0, i), Mode.Host);
                    port = colonPort.Slice(1).ToString();  // already validated (contains only digits)
                }
                else
                {
                    host = Unescape(value, Mode.Host);
                    port = null;
                }
            }

            return (host, port);
        }

        // ValidOptionalPort reports whether port is either an empty string
        // or matches /^:\d*$/
        private static bool ValidOptionalPort(ReadOnlySpan<char> value)
        {
            if (value.IsEmpty)
                return true;
            if (value[0] != ':')
                return false;

            for (var i = 1; i < value.Length; ++i)
            {
                if (value[i] < '0' || value[i] > '9')
                    return false;
            }

            return true;
        }

        // ValidUserInfo reports whether s is a valid userinfo string per RFC 3986
        // Section 3.2.1:
        //     userinfo    = *( unreserved / pct-encoded / sub-delims / ":" )
        //     unreserved  = ALPHA / DIGIT / "-" / "." / "_" / "~"
        //     sub-delims  = "!" / "$" / "&" / "'" / "(" / ")"
        //                   / "*" / "+" / "," / ";" / "="
        //
        // It doesn't validate pct-encoded. The caller does that via func unescape.
        private static bool ValidUserInfo(ReadOnlySpan<char> value)
        {
            foreach (var c in value)
            {
                if (('A' <= c && c <= 'Z' ) || ('a' <= c && c <= 'z') || ('0' <= c && c <= '9'))
                    continue;

                switch (c)
                {
                    case '-':
                    case '.':
                    case '_':
                    case ':':
                    case '~':
                    case '!':
                    case '$':
                    case '&':
                    case '\'':
                    case '(':
                    case ')':
                    case '*':
                    case '+':
                    case ',':
                    case ';':
                    case '=':
                    case '%':
                    case '@':
                        continue;

                    default:
                        return false;
                }
            }

            return true;
        }

        private static bool ContainsCTLByte(ReadOnlySpan<char> value)
        {
            foreach (var c in value)
            {
                if (c < ' ' || c == '\x7f')
                    return true;
            }

            return false;
        }

        // Unescape unescapes a string; the mode specifies
        // which section of the URL string is being unescaped.
        private static string Unescape(ReadOnlySpan<char> value, Mode mode)
        {
            // Count %, check that they're well-formed.
            var n = 0;
            //var hasPlus = false;

            for (var i = 0; i < value.Length; ++i)
            {
                switch (value[i])
                {
                    case '%':
                        ++n;

                        if (i + 2 >= value.Length)
                            throw new ArgumentException($"invalid URL escape \"{value.Slice(i, value.Length - i).ToString()}\"");
                        if (!IsHex(value[i + 1]) || !IsHex(value[i + 2]))
                            throw new ArgumentException($"invalid URL escape \"{value.Slice(i, 3).ToString()}\"");

                        // Per https://tools.ietf.org/html/rfc3986#page-21
                        // in the host component %-encoding can only be used
                        // for non-ASCII bytes.
                        // But https://tools.ietf.org/html/rfc6874#section-2
                        // introduces %25 being allowed to escape a percent sign
                        // in IPv6 scoped-address literals. Yay.
                        if (mode == Mode.Host && Unhex(value[i + 1]) < 8 && !(value[i + 1] == '2' && value[i + 1] == '5'))
                            throw new ArgumentException($"invalid URL escape \"{value.Slice(i, 3).ToString()}\"");

                        // RFC 6874 says basically "anything goes" for zone identifiers
                        // and that even non-ASCII can be redundantly escaped,
                        // but it seems prudent to restrict %-escaped bytes here to those
                        // that are valid host name bytes in their unescaped form.
                        // That is, you can use escaping in the zone identifier but not
                        // to introduce bytes you couldn't just write directly.
                        // But Windows puts spaces here! Yay.
                        if (mode == Mode.Zone)
                        {
                            var v = (Unhex(value[i + 1]) << 4) | Unhex(value[i + 2]);
                            if (v != '%' && v != ' ' && ShouldEscape(v, Mode.Host))
                                throw new ArgumentException($"invalid URL escape \"{value.Slice(i, 3).ToString()}\"");
                        }

                        i += 2;
                        break;

                    //case '+':
                    //    hasPlus = Mode.Query;
                    //    break;

                    default:
                        if ((mode == Mode.Host || mode == Mode.Zone) && value[i] < 0x80 && ShouldEscape(value[i], mode))
                            throw new ArgumentException($"invalid character \"{value.Slice(i, 1).ToString()}\" in host name");
                        break;
                }
            }

            if (n == 0)
                return value.ToString();
            //if (!hasPlus)
            //    return value.ToString();

            var bytes = new byte[value.Length - 2 * n];
            var j = 0;

            for (var i = 0; i < value.Length; ++i)
            {
                switch (value[i])
                {
                    case '%':
                        bytes[j] = Convert.ToByte((Unhex(value[i + 1]) << 4) | Unhex(value[i + 2]));
                        i += 2;
                        break;

                    //case '+':
                    //    builder.Append(mode == Mode.Query ? ' ' : '+');
                    //    break;

                    default:
                        bytes[j] = Convert.ToByte(value[i]);
                        break;
                }

                ++j;
            }

            return Encoding.UTF8.GetString(bytes);
        }

        // Return true if the specified character should be escaped when
        // appearing in a URL string, according to RFC 3986.
        //
        // Please be informed that for now shouldEscape does not check all
        // reserved characters correctly. See golang.org/issue/5684.
        private static bool ShouldEscape(int c, Mode mode)
        {
            // §2.3 Unreserved characters (alphanum)
	        if ('a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || '0' <= c && c <= '9')
		        return false;

            if (mode == Mode.Host || mode == Mode.Zone)
            {
                // §3.2.2 Host allows
                //	sub-delims = "!" / "$" / "&" / "'" / "(" / ")" / "*" / "+" / "," / ";" / "="
                // as part of reg-name.
                // We add : because we include :port as part of host.
                // We add [ ] because we include [ipv6]:port as part of host.
                // We add < > because they're the only characters left that
                // we could possibly allow, and Parse will reject them if we
                // escape them (because hosts can't use %-encoding for
                // ASCII bytes).
                if (c == '!' || c == '$' || c == '&' || c == '\'' || c == '(' || c ==  ')'
                    || c == '*' || c == '+' || c == ',' || c == ';' || c == '=' || c == ':'
                    || c == '[' || c == ']' || c == '<' || c == '>' || c == '"')
                    return false;
            }

            switch (c)
            {
                case '-':
                case '_':
                case '.':
                case '~':
                    // §2.3 Unreserved characters (mark)
                    return false;

	            case '$':
                case '&':
                case '+':
                case ',':
                case '/':
                case ':':
                case ';':
                case '=':
                case '?':
                case '@':
                    // §2.2 Reserved characters (reserved)
                    // Different sections of the URL allow a few of
                    // the reserved characters to appear unescaped.
                    switch (mode) {
                        case Mode.Path:
                            // §3.3
                            // The RFC allows : @ & = + $ but saves / ; , for assigning
                            // meaning to individual path segments. This package
                            // only manipulates the path as a whole, so we allow those
                            // last three as well. That leaves only ? to escape.
                            return c == '?';

                        case Mode.UsernamePassword:
                             // §3.2.1
                            // The RFC allows ';', ':', '&', '=', '+', '$', and ',' in
                            // userinfo, so we must escape only '@', '/', and '?'.
                            // The parsing of userinfo treats ':' as special so we must escape
                            // that too.
                            return c == '@' || c == '/' || c == '?' || c == ':';

                        //case Mode.Query:
                        //    // §3.4
                        //    // The RFC reserves (so we must escape) everything.
                        //    return true;

                        case Mode.Fragment: // §4.1
                            // The RFC text is silent but the grammar allows
                            // everything, so escape nothing.
                            return false;
                    }
                    break;
            }

            if (mode == Mode.Fragment)
            {
                // RFC 3986 §2.2 allows not escaping sub-delims. A subset of sub-delims are
                // included in reserved from RFC 2396 §2.2. The remaining sub-delims do not
                // need to be escaped. To minimize potential breakage, we apply two restrictions:
                // (1) we always escape sub-delims outside of the fragment, and (2) we always
                // escape single quote to avoid breaking callers that had previously assumed that
                // single quotes would be escaped. See issue #19917.
                if (c ==  '!' || c ==  '(' || c ==  ')' || c ==  '*')
                    return false;
            }

            // Everything else must be escaped.
            return true;
        }

        private static bool IsHex(char c)
        {
            return (('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F'));
        }

        private static int Unhex(char c)
        {
            if (c >= '0' && c <= '9')
                return c - '0';

            if (c >= 'A' && c <= 'F')
                return c - 'A' + 10;

            if (c >= 'a' && c <= 'f')
                return c - 'a' + 10;

            return 0;
        }

        private enum Mode
        {
            UsernamePassword,
            Host,
            Zone,
            Path,
            Fragment
        }
    }
}
