using System;
using System.Collections.Generic;
using System.Text;

namespace Envoyproxy.Validator
{
    public class EmailAddress
    {
        public string Name { get; set; }
        public string Address { get; set; }
    }

    /// RFC 5322 address parser
    /// Re-Implementation of Golang's net/mail/ParseAddress
    /// (https://golang.org/src/net/mail/message.go)
    public ref struct EmailAddressParser
    {
        private ReadOnlySpan<char> _value;

        private EmailAddressParser(ReadOnlySpan<char> value)
        {
            _value = value;
        }

        public static EmailAddress Parse(string value)
        {
            return Parse(value.AsSpan());
        }

        public static EmailAddress Parse(ReadOnlySpan<char> value)
        {
            var parser = new EmailAddressParser(value);
            return parser.ParseSingleAddress();
        }

        public static IList<EmailAddress> ParseList(string value)
        {
            return ParseList(value.AsSpan());
        }

        public static IList<EmailAddress> ParseList(ReadOnlySpan<char> value)
        {
            var parser = new EmailAddressParser(value);
            return parser.ParseAddressList();
        }

        private EmailAddress ParseSingleAddress()
        {
            var addrs = ParseAddress(true);

            if (!SkipCFWS())
                throw new ArgumentException("misformatted parenthetical comment");
            if (!_value.IsEmpty)
                throw new ArgumentException("expected single address");
            if (addrs.Count == 0)
                throw new ArgumentException("empty group");
            if (addrs.Count > 1)
                throw new ArgumentException("group with multiple addresses");

            return addrs[0];
        }

        private IList<EmailAddress> ParseAddressList()
        {
            var list = new List<EmailAddress>();

            for (;;)
            {
                SkipSpace();

                var addrs = ParseAddress(true);
                list.AddRange(addrs);

                if (!SkipCFWS())
                    throw new ArgumentException("misformatted parenthetical comment");
                if (_value.IsEmpty)
                    break;
                if (!Consume(','))
                    throw new ArgumentException("expected comma");
            }

            return list;
        }

        private IList<EmailAddress> ParseAddress(bool handleGroup)
        {
            SkipSpace();
            if (_value.IsEmpty)
                throw new ArgumentException("no address");

            string displayName = "";
            string spec = "";

            // address = mailbox / group
            // mailbox = name-addr / addr-spec
            // group = display-name ":" [group-list] ";" [CFWS]

            // addr-spec has a more restricted grammar than name-addr,
            // so try parsing it first, and fallback to name-addr.
            try
            {
                spec = ConsumeAddrSpec();
            }
            catch (ArgumentException)
            {
                spec = null;
            }

            if (spec != null)
            {
                SkipSpace();
                if (!_value.IsEmpty && _value[0] == '(')
                    displayName = ConsumeDisplayNameComment();

                return new List<EmailAddress>
                {
                    new EmailAddress
                    {
                        Name = displayName,
                        Address = spec
                    }
                };
            }

	        // display-name
            if (_value[0] != '<')
                displayName = ConsumePhrase();

            SkipSpace();
            if (handleGroup)
            {
                if (Consume(':'))
                    return ConsumeGroupList();
            }

	        // angle-addr = "<" addr-spec ">"
            if (!Consume('<'))
            {
                bool atext = true;
                foreach (var c in displayName)
                {
                    if (!IsAtext(c, true, false))
                    {
                        atext = false;
                        break;
                    }
                }

                if (atext)
                {
                    // The input is like "foo.bar"; it's possible the input
                    // meant to be "foo.bar@domain", or "foo.bar <...>".
                    throw new ArgumentException("missing '@' or angle-addr");
                }
                else
                {
                    // The input is like "Full Name", which couldn't possibly be a
                    // valid email address if followed by "@domain"; the input
                    // likely meant to be "Full Name <...>".
                    throw new ArgumentException("no angle-addr");
                }
            }

            spec = ConsumeAddrSpec();

            if (!Consume('>'))
                throw new ArgumentException("unclosed angle-addr");

            return new List<EmailAddress>
            {
                new EmailAddress
                {
                    Name = displayName,
                    Address = spec
                }
            };
        }

        private IList<EmailAddress> ConsumeGroupList()
        {
            var group = new List<EmailAddress>();

            // handle empty group.
            SkipSpace();
            if (Consume(';'))
            {
                SkipCFWS();
                return group;
            }

            for (;;)
            {
                SkipSpace();

                // embedded groups not allowed.
                var addrs = ParseAddress(false);
                group.AddRange(addrs);

                if (!SkipCFWS())
                    throw new ArgumentException("misformatted parenthetical comment");

                if (Consume(';'))
                {
                    SkipCFWS();
                    return group;
                }

                if (!Consume(','))
                    throw new ArgumentException("expected comma");
            }
        }

        // ConsumeAddrSpec parses a single RFC 5322 addr-spec at the start of _value
        private string ConsumeAddrSpec()
        {
            var orig = _value;

            try
            {
                // local-part = dot-atom / quoted-string
                string localPart;

                SkipSpace();
                if (_value.IsEmpty)
                    throw new ArgumentException("no addr-spec");

                if (_value[0] == '"')
                {
                    // quoted-string
                    localPart = ConsumeQuotedString();
                    if (localPart == "")
                        throw new ArgumentException("empty quoted string in addr-spec");
                }
                else
                {
                    // dot-atom
                    localPart = ConsumeAtom(true, false);
                }

                if (!Consume('@'))
                    throw new ArgumentException("missing @ in addr-spec");

                // domain = dot-atom / domain-literal
                string domain;

                SkipSpace();
                if (_value.IsEmpty)
                    throw new ArgumentException("no domain in addr-spec");

                domain = ConsumeAtom(true, false);

                return $"{localPart}@{domain}";
            }
            catch (ArgumentException)
            {
                _value = orig;
                throw;
            }
        }

        // ConsumePhrase parses the RFC 5322 phrase at the start of _value
        private string ConsumePhrase()
        {
            var words = new List<string>();
            var isPrevEncoded = false;

            try
            {
                for (;;)
                {
                    // word = atom / quoted-string
                    string word;

                    SkipSpace();
                    if (_value.IsEmpty)
                        break;

                    var isEncoded = false;

                    if (_value[0] == '"')
                    {
                        // quoted-string
                        word = ConsumeQuotedString();
                    }
                    else
                    {
                        // atom
                        // We actually parse dot-atom here to be more permissive
                        // than what RFC 5322 specifies.
                        word = ConsumeAtom(true, true);
                        (word, isEncoded) = DecodeRFC2047Word(word);
                    }

                    if (isPrevEncoded && isEncoded)
                    {
                        words[words.Count - 1] = words[words.Count - 1] + word;
                    }
                    else
                    {
                        words.Add(word);
                    }

                    isPrevEncoded = isEncoded;
                }
            }
            catch (ArgumentException ex)
            {
                // Ignore any error if we got at least one word.
                if (words.Count == 0)
                    throw new ArgumentException("missing word in phrase: " + ex.Message, ex);
            }

            return string.Join(" ", words);
        }

        // ConsumeQuotedString parses the quoted string at the start of _value
        private string ConsumeQuotedString()
        {
            // Assume first byte is '"'.
            var i = 1;
            var escaped = false;
            var quoted = new StringBuilder();

            for (; i < _value.Length; ++i)
            {
                if (escaped)
                {
                    //  quoted-pair = ("\" (VCHAR / WSP))

                    if (!IsVchar(_value[i]) && !IsWSP(_value[i]))
                        throw new ArgumentException("bad character in quoted-string");

                    quoted.Append(_value[i]);
                    escaped = false;
                }
                else if (IsQtext(_value[i]) || IsWSP(_value[i]))
                {
                    quoted.Append(_value[i]);
                }
                else if (_value[i] == '"')
                {
                    break;
                }
                else if (_value[i] == '\\')
                {
                    escaped = true;
                }
                else
                {
                    throw new ArgumentException("bad character in quoted-string");
                }
            }

            if (i == _value.Length)
                throw new ArgumentException("unclosed quoted-string");

            _value = _value.Slice(i + 1);

            return quoted.ToString();
        }

        // consumeAtom parses an RFC 5322 atom at the start of p.
        // If dot is true, consumeAtom parses an RFC 5322 dot-atom instead.
        // If permissive is true, consumeAtom will not fail on:
        // - leading/trailing/double dots in the atom (see golang.org/issue/4938)
        // - special characters (RFC 5322 3.2.3) except '<', '>', ':' and '"' (see golang.org/issue/21018)
        private string ConsumeAtom(bool dot, bool permissive)
        {
            var i = 0;

            for (; i < _value.Length; ++i)
            {
                if (!IsAtext(_value[i], dot, permissive))
                    break;
            }

            if (i == 0)
                throw new ArgumentException("invalid string");

            var atom = _value.Slice(0, i).ToString();
            _value = _value.Slice(i);

            if (!permissive)
            {
                if (atom.StartsWith("."))
                    throw new ArgumentException("leading dot in atom");
                if (atom.Contains(".."))
                    throw new ArgumentException("double dot in atom");
                if (atom.EndsWith("."))
                    throw new ArgumentException("trailing dot in atom");
            }

            return atom;
        }

        private string ConsumeDisplayNameComment()
        {
            if (!Consume('('))
                throw new ArgumentException("comment does not start with (");

            var comment = ConsumeComment();
            var words = comment.Split(" \t".ToCharArray(), StringSplitOptions.RemoveEmptyEntries);

            for (var i = 0; i < words.Length; ++i)
            {
                (var decoded, var isEncoded) = DecodeRFC2047Word(words[i]);

                if (isEncoded)
                    words[i] = decoded;
            }

            return string.Join(" ", words);
        }

        private bool Consume(char c)
        {
            if (_value.IsEmpty || _value[0] != c)
                return false;

            _value = _value.Slice(1);
            return true;
        }

        private string ConsumeComment()
        {
            // '(' already consumed.
            var i = 0;
            var depth = 1;
            var comment = new StringBuilder();

            for (; i < _value.Length && depth > 0; ++i)
            {
                if (_value[i] == '\\' && i + 1 < _value.Length)
                    ++i;
                else if (_value[i] == '(')
                    ++depth;
                else if (_value[i] == ')')
                    --depth;

                if (depth > 0)
                    comment.Append(_value[i]);
            }

            if (depth != 0)
                throw new ArgumentException("misformatted parenthetical comment");

            _value = _value.Slice(i);

            return comment.ToString();
        }

        // SkipSpace skips the leading space and tab characters.
        private void SkipSpace()
        {
            _value = _value.TrimStart(" \t".AsSpan());
        }

        // SkipCFWS skips CFWS as defined in RFC5322.
        private bool SkipCFWS()
        {
            SkipSpace();

            for (;;)
            {
                if (!Consume('('))
                    break;

                try
                {
                    ConsumeComment();
                }
                catch (ArgumentException)
                {
                    return false;
                }

                SkipSpace();
            }

            return true;
        }

        private static (string, bool) DecodeRFC2047Word(string word)
        {
            try
            {
                var decoded = WordDecoder.Decode(word);
                return (decoded, true);
            }
            catch (NotSupportedException)
            {
                throw new ArgumentException("charset not supported");
            }
            catch (ArgumentException)
            {
                // Ignore invalid RFC 2047 encoded-word errors.
                return (word, false);
            }
        }

        // IsAtext reports whether c is an RFC 5322 atext character.
        // If dot is true, period is included.
        // If permissive is true, RFC 5322 3.2.3 specials is included,
        // except '<', '>', ':' and '"'.
        private static bool IsAtext(char c, bool dot, bool permissive)
        {
            switch (c)
            {
                case '.':
                    return dot;

	            // RFC 5322 3.2.3. specials
                case '(':
                case ')':
                case '[':
                case ']':
                case ';':
                case '@':
                case '\\':
                case ',':
                    return permissive;

                case '<':
                case '>':
                case '"':
                case ':':
                    return false;
            }

            return IsVchar(c);
        }

        // IsQtext reports whether c is an RFC 5322 qtext character.
        private static bool IsQtext(char c)
        {
	        // Printable US-ASCII, excluding backslash or quote.
            if (c == '\\' || c == '"')
                return false;

	        return IsVchar(c);
        }

        // IsVchar reports whether c is an RFC 5322 VCHAR character.
        private static bool IsVchar(char c)
        {
            // Visible (printing) characters.
	        return '!' <= c && c <= '~' || IsMultibyte(c);
        }

        // IsMultibyte reports whether c is a multi-byte UTF-8 character
        // as supported by RFC 6532
        private static bool IsMultibyte(char c)
        {
            return Encoding.UTF8.GetByteCount(new char[] { c }, 0, 1) != 1;
        }

        // IsWSP reports whether r is a WSP (white space).
        // WSP is a space or horizontal tab (RFC 5234 Appendix B).
        private static bool IsWSP(char c)
        {
            return c == ' ' || c == '\t';
        }
    }
}
