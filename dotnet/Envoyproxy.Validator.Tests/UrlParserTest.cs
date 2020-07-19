using Microsoft.VisualStudio.TestTools.UnitTesting;
using System;
using System.Collections.Generic;
using Envoyproxy.Validator;

namespace Validator.Tests
{
    [TestClass]
    public class UrlParserTest
    {
        [TestMethod]
        [DynamicData(nameof(Data))]
        public void Parse(string value, Url expected)
        {
            var actual = UrlParser.Parse(value);

            Assert.AreEqual(expected.Scheme, actual.Scheme);
            Assert.AreEqual(expected.Opaque, actual.Opaque);
            Assert.AreEqual(expected.User == null, actual.User == null);
            Assert.AreEqual(expected.User?.Username, actual.User?.Username);
            Assert.AreEqual(expected.User?.Password, actual.User?.Password);
            Assert.AreEqual(expected.Hostname, actual.Hostname);
            Assert.AreEqual(expected.Port, actual.Port);
            Assert.AreEqual(expected.Path, actual.Path);
            Assert.AreEqual(expected.RawQuery, actual.RawQuery);
            Assert.AreEqual(expected.Fragment, actual.Fragment);
        }

        [TestMethod]
        [DataRow("http://[::1]")]
        [DataRow("http://[::1]:80")]
        [DataRow("http://[::1]/")]
        [DataRow("http://[::1%25en0]")]  // valid zone id
        [DataRow("http://[::1]:")]  // colon, but no port OK
        [DataRow("http://x:")]  // colon, but no port OK
        [DataRow("http://[::1%25%41]")]  // RFC 6874 allows over-escaping in zone
        [DataRow("http://[::1]/%48")]  // %xx in path is fine
        [DataRow("ahttp://foo.com")]  // valid schema characters
        [DataRow("cache_object/:foo/bar")]
        public void Valid(string value)
        {
            UrlParser.Parse(value);

            // Assert "not throws"
        }

        [TestMethod]
        [DataRow("http://[::1]:namedport")]  // rfc3986 3.2.3
        [DataRow("http://x:namedport")]  // rfc3986 3.2.3
        [DataRow("http://[::1]a")]
        [DataRow("http://[::1]%23")]
        [DataRow("http://[::1]:%38%30")]  // not allowed: % encoding only for non-ASCII
        [DataRow("http://[%10::1]")]  // no %xx escapes in IP address
        [DataRow("http://%41:8080/")]  // not allowed: % encoding only for non-ASCII
        [DataRow("mysql://x@y(z:123)/foo")]  // not well-formed per RFC 3986, golang.org/issue/33646
        [DataRow("mysql://x@y(1.2.3.4:123)/foo")]
        [DataRow(" http://foo.com")]  // invalid character in schema
        [DataRow("ht tp://foo.com")]  // invalid character in schema
        [DataRow("1http://foo.com")]  // invalid character in schema
        [DataRow("http://[]%20%48%54%54%50%2f%31%2e%31%0a%4d%79%48%65%61%64%65%72%3a%20%31%32%33%0a%0a/")]  // golang.org/issue/11208
        [DataRow("http://a b.com/")]  // no space in host name please
        [DataRow("cache_object://foo")]  // scheme cannot have _, relative path cannot have : in first segment
        [DataRow("cache_object:foo")]
        [DataRow("cache_object:foo/bar")]
        public void Errors(string value)
        {
            Assert.ThrowsException<ArgumentException>(() => UrlParser.Parse(value));
        }

        [TestMethod]
        [DataRow("http://foo.com:80", "foo.com", "80")]
        [DataRow("http://foo.com", "foo.com", null)]
        [DataRow("http://foo.com:", "foo.com", "")]
        [DataRow("http://FOO.COM", "FOO.COM", null)]  // no canonicalization
        [DataRow("http://1.2.3.4", "1.2.3.4", null)]
        [DataRow("http://1.2.3.4:80", "1.2.3.4", "80")]
        [DataRow("http://[1:2:3:4]", "1:2:3:4", null)]
        [DataRow("http://[1:2:3:4]:80", "1:2:3:4", "80")]
        [DataRow("http://[::1]:80", "::1", "80")]
        [DataRow("http://[::1]", "::1", null)]
        [DataRow("http://[::1]:", "::1", "")]
        [DataRow("http://localhost", "localhost", null)]
        [DataRow("http://localhost:443", "localhost", "443")]
        [DataRow("http://some.super.long.domain.example.org:8080", "some.super.long.domain.example.org", "8080")]
        [DataRow("http://[2001:0db8:85a3:0000:0000:8a2e:0370:7334]:17000", "2001:0db8:85a3:0000:0000:8a2e:0370:7334", "17000")]
        [DataRow("http://[2001:0db8:85a3:0000:0000:8a2e:0370:7334]", "2001:0db8:85a3:0000:0000:8a2e:0370:7334", null)]
        public void HostnameAndPort(string value, string expectedHostname, string expectedPort)
        {
            var url = UrlParser.Parse(value);

            Assert.AreEqual(expectedHostname, url.Hostname);
            Assert.AreEqual(expectedPort, url.Port);
        }


        [TestMethod]
        public void NullUser()
        {
            var url = UrlParser.Parse("http://foo.com/");

            Assert.IsNull(url.User);
        }


        [TestMethod]
        public void InvalidUserPassword()
        {
            ArgumentException actual = null;

            try
            {
                UrlParser.Parse("http://user^:passwo^rd@foo.com/");
            }
            catch (ArgumentException ex)
            {
                actual = ex;
            }

            Assert.IsNotNull(actual);
            Console.WriteLine(actual.Message);
            Assert.IsTrue(actual.Message.Contains("invalid userinfo"));
        }

        [TestMethod]
        [DataRow("http://foo.com/?foo\nbar")]
        [DataRow("http\r://foo.com/")]
        [DataRow("http://foo\x7f.com/")]
        public void RejectControlCharacters(string value)
        {
            ArgumentException actual = null;

            try
            {
                UrlParser.Parse(value);
            }
            catch (ArgumentException ex)
            {
                actual = ex;
            }

            Assert.IsNotNull(actual);
            Assert.IsTrue(actual.Message.Contains("invalid control character in URL"));
        }

        static IEnumerable<object[]> Data
        {
            get
            {
                return new[]
                {
                    // no path
                    new object[]
                    {
                        "http://www.google.com",
                        new Url
                        {
                            Scheme = "http",
                            Hostname = "www.google.com"
                        }
                    },

	                // path
                    new object[]
                    {
                        "http://www.google.com/",
                        new Url
                        {
                            Scheme = "http",
                            Hostname = "www.google.com",
                            Path = "/"
                        }
                    },

	                // path with hex escaping
                    new object[]
                    {
                        "http://www.google.com/file%20one%26two",
                        new Url
                        {
                            Scheme = "http",
                            Hostname = "www.google.com",
                            Path = "/file one&two"
                        }
                    },

	                // user
                    new object[]
                    {
                        "ftp://webmaster@www.google.com/",
                        new Url
                        {
                            Scheme = "ftp",
                            User = new UserInfo
                            {
                                Username = "webmaster"
                            },
                            Hostname = "www.google.com",
			                Path = "/"
                        }
                    },

	                // escape sequence in username
                    new object[]
                    {
                        "ftp://john%20doe@www.google.com/",
                        new Url
                        {
                            Scheme = "ftp",
                            User = new UserInfo
                            {
                                Username = "john doe"
                            },
                            Hostname = "www.google.com",
			                Path = "/"
                        }
                    },

                    // empty query
                    new object[]
                    {
                        "http://www.google.com/?",
                        new Url
                        {
                            Scheme = "http",
                            Hostname = "www.google.com",
			                Path = "/",
                            RawQuery = ""
                        }
                    },

                    // query ending in question mark (Issue 14573)
                    new object[]
                    {
                        "http://www.google.com/?foo=bar?",
                        new Url
                        {
                            Scheme = "http",
                            Hostname = "www.google.com",
			                Path = "/",
                            RawQuery = "foo=bar?"
                        }
                    },

                    // query
                    new object[]
                    {
                        "http://www.google.com/?q=go+language",
                        new Url
                        {
                            Scheme = "http",
                            Hostname = "www.google.com",
			                Path = "/",
                            RawQuery = "q=go+language"
                        }
                    },

                    // query with hex escaping: NOT parsed
                    new object[]
                    {
                        "http://www.google.com/?q=go%20language",
                        new Url
                        {
                            Scheme = "http",
                            Hostname = "www.google.com",
			                Path = "/",
                            RawQuery = "q=go%20language"
                        }
                    },

                    // %20 outside query
                    new object[]
                    {
                        "http://www.google.com/a%20b?q=c+d",
                        new Url
                        {
                            Scheme = "http",
                            Hostname = "www.google.com",
			                Path = "/a b",
                            RawQuery = "q=c+d"
                        }
                    },

                    // path without leading /, so no parsing
                    new object[]
                    {
                        "http:www.google.com/?q=go+language",
                        new Url
                        {
                            Scheme = "http",
                            Opaque = "www.google.com/",
                            RawQuery = "q=go+language"
                        }
                    },

                    // path without leading /, so no parsing
                    new object[]
                    {
                        "http:%2f%2fwww.google.com/?q=go+language",
                        new Url
                        {
                            Scheme = "http",
                            Opaque = "%2f%2fwww.google.com/",
                            RawQuery = "q=go+language"
                        }
                    },

                    // non-authority with path
                    new object[]
                    {
                        "mailto:/webmaster@golang.org",
                        new Url
                        {
                            Scheme = "mailto",
                            Path = "/webmaster@golang.org"
                        }
                    },

                    // non-authority
                    new object[]
                    {
                        "mailto:webmaster@golang.org",
                        new Url
                        {
                            Scheme = "mailto",
                            Opaque = "webmaster@golang.org"
                        }
                    },

                    // unescaped :// in query should not create a scheme
                    new object[]
                    {
                        "/foo?query=http://bad",
                        new Url
                        {
                            Path = "/foo",
                            RawQuery = "query=http://bad"
                        }
                    },

                    // leading // without scheme should create an authority
                    new object[]
                    {
                        "//foo",
                        new Url
                        {
                            Hostname = "foo"
                        }
                    },

                    // leading // without scheme, with userinfo, path, and query
                    new object[]
                    {
                        "//user@foo/path?a=b",
                        new Url
                        {
                            User = new UserInfo
                            {
                                Username = "user"
                            },
                            Hostname = "foo",
                            Path = "/path",
                            RawQuery = "a=b"
                        }
                    },

                    // Three leading slashes isn't an authority, but doesn't return an error.
                    // (We can't return an error, as this code is also used via
                    // ServeHTTP -> ReadRequest -> Parse, which is arguably a
                    // different URL parsing context, but currently shares the
                    // same codepath)
                    new object[]
                    {
                        "///threeslashes",
                        new Url
                        {
                            Path = "///threeslashes"
                        }
                    },

                    new object[]
                    {
                        "http://user:password@google.com",
                        new Url
                        {
                            Scheme = "http",
                            User = new UserInfo
                            {
                                Username = "user",
                                Password = "password"
                            },
                            Hostname = "google.com"
                        }
                    },

                    // unescaped @ in username should not confuse host
                    new object[]
                    {
                        "http://j@ne:password@google.com",
                        new Url
                        {
                            Scheme = "http",
                            User = new UserInfo
                            {
                                Username = "j@ne",
                                Password = "password"
                            },
                            Hostname = "google.com"
                        }
                    },

                    // unescaped @ in password should not confuse host
                    new object[]
                    {
                        "http://jane:p@ssword@google.com",
                        new Url
                        {
                            Scheme = "http",
                            User = new UserInfo
                            {
                                Username = "jane",
                                Password = "p@ssword"
                            },
                            Hostname = "google.com"
                        }
                    },

                    new object[]
                    {
                        "http://j@ne:password@google.com/p@th?q=@go",
                        new Url
                        {
                            Scheme = "http",
                            User = new UserInfo
                            {
                                Username = "j@ne",
                                Password = "password"
                            },
                            Hostname = "google.com",
                            Path = "/p@th",
                            RawQuery = "q=@go"
                        }
                    },

                    new object[]
                    {
                        "http://www.google.com/?q=go+language#foo",
                        new Url
                        {
                            Scheme = "http",
                            Hostname = "www.google.com",
                            Path = "/",
                            RawQuery = "q=go+language",
                            Fragment = "foo"
                        }
                    },

                    new object[]
                    {
                        "http://www.google.com/?q=go+language#foo%26bar",
                        new Url
                        {
                            Scheme = "http",
                            Hostname = "www.google.com",
                            Path = "/",
                            RawQuery = "q=go+language",
                            Fragment = "foo&bar"
                        }
                    },

                    new object[]
                    {
                        "file:///home/adg/rabbits",
                        new Url
                        {
                            Scheme = "file",
                            Path = "/home/adg/rabbits"
                        }
                    },

                    // "Windows" paths are no exception to the rule.
                    // See golang.org/issue/6027, especially comment #9.
                    new object[]
                    {
                        "file:///C:/FooBar/Baz.txt",
                        new Url
                        {
                            Scheme = "file",
                            Path = "/C:/FooBar/Baz.txt"
                        }
                    },

                    // case-insensitive scheme
                    new object[]
                    {
                        "MaIlTo:webmaster@golang.org",
                        new Url
                        {
                            Scheme = "mailto",
                            Opaque = "webmaster@golang.org"
                        }
                    },

                    // Relative path
                    new object[]
                    {
                        "a/b/c",
                        new Url
                        {
                            Path = "a/b/c"
                        }
                    },

                    // escaped '?' in username and password
                    new object[]
                    {
                        "http://%3Fam:pa%3Fsword@google.com",
                        new Url
                        {
                            Scheme = "http",
                            User = new UserInfo
                            {
                                Username = "?am",
                                Password = "pa?sword"
                            },
                            Hostname = "google.com"
                        }
                    },

                    // host subcomponent; IPv4 address in RFC 3986
                    new object[]
                    {
                        "http://192.168.0.1/",
                        new Url
                        {
                            Scheme = "http",
                            Hostname = "192.168.0.1",
                            Path = "/"
                        }
                    },

                    // host and port subcomponents; IPv4 address in RFC 3986
                    new object[]
                    {
                        "http://192.168.0.1:8080/",
                        new Url
                        {
                            Scheme = "http",
                            Hostname = "192.168.0.1",
                            Port = "8080",
                            Path = "/"
                        }
                    },

                    // host subcomponent; IPv6 address in RFC 3986
                    new object[]
                    {
                        "http://[fe80::1]/",
                        new Url
                        {
                            Scheme = "http",
                            Hostname = "fe80::1",
                            Path = "/"
                        }
                    },

                    // host and port subcomponent; IPv6 address in RFC 3986
                    new object[]
                    {
                        "http://[fe80::1]:8080/",
                        new Url
                        {
                            Scheme = "http",
                            Hostname = "fe80::1",
                            Port = "8080",
                            Path = "/"
                        }
                    },

                    // host subcomponent; IPv6 address with zone identifier in RFC 6874
                    new object[]
                    {
                        "http://[fe80::1%25en0]/",  // alphanum zone identifier
                        new Url
                        {
                            Scheme = "http",
                            Hostname = "fe80::1%en0",
                            Path = "/"
                        }
                    },

                    // host and port subcomponents; IPv6 address with zone identifier in RFC 6874
                    new object[]
                    {
                        "http://[fe80::1%25en0]:8080/",  // alphanum zone identifier
                        new Url
                        {
                            Scheme = "http",
                            Hostname = "fe80::1%en0",
                            Port = "8080",
                            Path = "/"
                        }
                    },

                    // host subcomponent; IPv6 address with zone identifier in RFC 6874
                    new object[]
                    {
                        "http://[fe80::1%25%65%6e%301-._~]/",  // percent-encoded+unreserved zone identifier
                        new Url
                        {
                            Scheme = "http",
                            Hostname = "fe80::1%en01-._~",
                            Path = "/"
                        }
                    },

                    // host and port subcomponents; IPv6 address with zone identifier in RFC 6874
                    new object[]
                    {
                        "http://[fe80::1%25%65%6e%301-._~]:8080/",  // percent-encoded+unreserved zone identifier
                        new Url
                        {
                            Scheme = "http",
                            Hostname = "fe80::1%en01-._~",
                            Port = "8080",
                            Path = "/"
                        }
                    },

                    // alternate escapings of path survive round trip
                    new object[]
                    {
                        "http://rest.rsc.io/foo%2fbar/baz%2Fquux?alt=media",
                        new Url
                        {
                            Scheme = "http",
                            Hostname = "rest.rsc.io",
                            Path = "/foo/bar/baz/quux",
                            RawQuery = "alt=media"
                        }
                    },

                    // issue 12036
                    new object[]
                    {
                        "mysql://a,b,c/bar",
                        new Url
                        {
                            Scheme = "mysql",
                            Hostname = "a,b,c",
                            Path = "/bar"
                        }
                    },

                    // worst case host, still round trips
                    new object[]
                    {
                        "scheme://!$&'()*+,;=hello!:1/path",
                        new Url
                        {
                            Scheme = "scheme",
                            Hostname = "!$&'()*+,;=hello!",
                            Port = "1",
                            Path = "/path"
                        }
                    },

                    // worst case path, still round trips
                    new object[]
                    {
                        "http://host/!$&'()*+,;=:@[hello]",
                        new Url
                        {
                            Scheme = "http",
                            Hostname = "host",
                            Path = "/!$&'()*+,;=:@[hello]"
                        }
                    },

                    // golang.org/issue/5684
                    new object[]
                    {
                        "http://example.com/oid/[order_id]",
                        new Url
                        {
                            Scheme = "http",
                            Hostname = "example.com",
                            Path = "/oid/[order_id]"
                        }
                    },

                    // golang.org/issue/12200 (colon with empty port)
                    new object[]
                    {
                        "http://192.168.0.2:/foo",
                        new Url
                        {
                            Scheme = "http",
                            Hostname = "192.168.0.2",
                            Port = "",
                            Path = "/foo"
                        }
                    },

                    // Malformed IPv6 but still accepted.
                    new object[]
                    {
                        "http://2b01:e34:ef40:7730:8e70:5aff:fefe:edac:8080/foo",
                        new Url
                        {
                            Scheme = "http",
                            Hostname = "2b01:e34:ef40:7730:8e70:5aff:fefe:edac",
                            Port = "8080",
                            Path = "/foo"
                        }
                    },

                    // Malformed IPv6 but still accepted.
                    new object[]
                    {
                        "http://2b01:e34:ef40:7730:8e70:5aff:fefe:edac:/foo",
                        new Url
                        {
                            Scheme = "http",
                            Hostname = "2b01:e34:ef40:7730:8e70:5aff:fefe:edac",
                            Port = "",
                            Path = "/foo"
                        }
                    },

                    new object[]
                    {
                        "http://[2b01:e34:ef40:7730:8e70:5aff:fefe:edac]:8080/foo",
                        new Url
                        {
                            Scheme = "http",
                            Hostname = "2b01:e34:ef40:7730:8e70:5aff:fefe:edac",
                            Port = "8080",
                            Path = "/foo"
                        }
                    },

                    new object[]
                    {
                        "http://[2b01:e34:ef40:7730:8e70:5aff:fefe:edac]:/foo",
                        new Url
                        {
                            Scheme = "http",
                            Hostname = "2b01:e34:ef40:7730:8e70:5aff:fefe:edac",
                            Port = "",
                            Path = "/foo"
                        }
                    },

                    // golang.org/issue/7991 and golang.org/issue/12719 (non-ascii %-encoded in host)
                    new object[]
                    {
                        "http://hello.世界.com/foo",
                        new Url
                        {
                            Scheme = "http",
                            Hostname = "hello.世界.com",
                            Path = "/foo"
                        }
                    },

                    new object[]
                    {
                        "http://hello.%e4%b8%96%e7%95%8c.com/foo",
                        new Url
                        {
                            Scheme = "http",
                            Hostname = "hello.世界.com",
                            Path = "/foo"
                        }
                    },

                    new object[]
                    {
                        "http://hello.%E4%B8%96%E7%95%8C.com/foo",
                        new Url
                        {
                            Scheme = "http",
                            Hostname = "hello.世界.com",
                            Path = "/foo"
                        }
                    },

                    // golang.org/issue/10433 (path beginning with //)
                    new object[]
                    {
                        "http://example.com//foo",
                        new Url
                        {
                            Scheme = "http",
                            Hostname = "example.com",
                            Path = "//foo"
                        }
                    },

                    // test that we can reparse the host names we accept.
                    new object[]
                    {
                        "myscheme://authority<\"hi\">/foo",
                        new Url
                        {
                            Scheme = "myscheme",
                            Hostname = "authority<\"hi\">",
                            Path = "/foo"
                        }
                    },

                    // spaces in hosts are disallowed but escaped spaces in IPv6 scope IDs are grudgingly OK.
                    // This happens on Windows.
                    // golang.org/issue/14002
                    new object[]
                    {
                        "tcp://[2020::2020:20:2020:2020%25Windows%20Loves%20Spaces]:2020",
                        new Url
                        {
                            Scheme = "tcp",
                            Hostname = "2020::2020:20:2020:2020%Windows Loves Spaces",
                            Port = "2020"
                        }
                    },

                    // test we can roundtrip magnet url
                    // fix issue https://golang.org/issue/20054
                    new object[]
                    {
                        "magnet:?xt=urn:btih:c12fe1c06bba254a9dc9f519b335aa7c1367a88a&dn",
                        new Url
                        {
                            Scheme = "magnet",
                            RawQuery = "xt=urn:btih:c12fe1c06bba254a9dc9f519b335aa7c1367a88a&dn"
                        }
                    },

                    new object[]
                    {
                        "mailto:?subject=hi",
                        new Url
                        {
                            Scheme = "mailto",
                            RawQuery = "subject=hi"
                        }
                    }
                };
            }
        }
    }
}
