using Microsoft.VisualStudio.TestTools.UnitTesting;
using System;
using System.Collections.Generic;
using Envoyproxy.Validator;

namespace Validator.Tests
{
    [TestClass]
    public class EmailParserTest
    {
        [TestMethod]
        [DynamicData(nameof(ErrorData))]
        public void CheckError(string value, string expected)
        {
            // Act
            ArgumentException actual = null;

            try
            {
                EmailAddressParser.Parse(value);
            }
            catch (ArgumentException ex)
            {
                actual = ex;
            }

            // Assert
            Assert.IsNotNull(actual);
            Assert.IsTrue(actual.Message.Contains(expected), "expected {0} got {1}", expected, actual);
        }

        [TestMethod]
        [DynamicData(nameof(TestData))]
        public void CheckResult(string value, EmailAddress[] expected)
        {
            // Act
            IList<EmailAddress> actual;

            if (expected.Length == 1)
            {
                actual = new List<EmailAddress>
                {
                    EmailAddressParser.Parse(value)
                };
            }
            else
            {
                actual = EmailAddressParser.ParseList(value);
            }

            // Assert
            Assert.IsNotNull(actual);
            Assert.AreEqual(expected.Length, actual.Count);

            for (int i = 0; i < expected.Length; ++i)
            {
                Assert.IsNotNull(actual[i]);
                Assert.AreEqual(expected[i].Address, actual[i].Address);
                Assert.AreEqual(expected[i].Name, actual[i].Name);
            }
        }

        static IEnumerable<object[]> ErrorData
        {
            get
            {
                return new[]
                {
                    new object[] { "=?iso-8859-2?Q?Bogl=E1rka_Tak=E1cs?= <unknown@gmail.com>", "charset not supported" },
                    new object[] { "a@gmail.com b@gmail.com", "expected single address" },
                    // TODO: How would these scenarios have to be defined???
                    //new object[] { BytesToString(new byte[] { 0xed, 0xa0, 0x80 }) + " <micro@example.net>", "invalid utf-8 in address" },
                    //new object[] { "\"" + BytesToString(new byte[] { 0xed, 0xa0, 0x80 }) + " <half-surrogate@example.com>", "invalid utf-8 in quoted-string" },
                    //new object[] { "\"\\" + BytesToString(new byte[] { 0x80 }) + " <escaped-invalid-unicode@example.net>", "invalid utf-8 in quoted-string" },
                    new object[] { "\"\x00\" <null@example.net>", "bad character in quoted-string" },
                    new object[] { "\"\\\x00\" <escaped-null@example.net>", "bad character in quoted-string" },
                    new object[] { "John Doe", "no angle-addr" },
                    new object[] { "<jdoe#machine.example>", "missing @ in addr-spec" },
                    new object[] { "John <middle> Doe <jdoe@machine.example>", "missing @ in addr-spec" },
                    new object[] { "cfws@example.com (", "misformatted parenthetical comment" },
                    new object[] { "empty group: ;", "empty group" },
                    new object[] { "root group: embed group: null@example.com;", "no angle-addr" },
                    new object[] { "group not closed: null@example.com", "expected comma" },
                    new object[] { "group: first@example.com, second@example.com;", "group with multiple addresses" },
                    new object[] { "john.doe", "missing '@' or angle-addr" },
                    new object[] { "john.doe@", "no angle-addr" },
                    new object[] { "John Doe@foo.bar", "no angle-addr" },
                };
            }
        }

        static IEnumerable<object[]> TestData
        {
            get
            {
                return new[]
                {
                    // Bare address
                    new object[]
                    {
                        "jdoe@machine.example",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "",
                                Address = "jdoe@machine.example"
                            }
                        }
                    },
                    // RFC 5322, Appendix A.1.1
                    new object[]
                    {
                        "John Doe <jdoe@machine.example>",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "John Doe",
                                Address = "jdoe@machine.example"
                            }
                        }
                    },
                    // RFC 5322, Appendix A.1.2
                    new object[]
                    {
                        "\"Joe Q. Public\" <john.q.public@example.com>",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "Joe Q. Public",
                                Address = "john.q.public@example.com"
                            }
                        }
                    },
                    new object[]
                    {
			            "\"John (middle) Doe\" <jdoe@machine.example>",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "John (middle) Doe",
                                Address = "jdoe@machine.example"
                            }
                        }
		            },
                    new object[]
                    {
                        "John (middle) Doe <jdoe@machine.example>",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "John (middle) Doe",
                                Address = "jdoe@machine.example"
                            }
                        }
		            },
                    new object[]
                    {
                        "John !@M@! Doe <jdoe@machine.example>",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                				Name = "John !@M@! Doe",
                                Address = "jdoe@machine.example"
                            }
                        }
            		},
                    new object[]
		            {
                        "\"John <middle> Doe\" <jdoe@machine.example>",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "John <middle> Doe",
                                Address = "jdoe@machine.example"
                            }
                        }
		            },
                    new object[]
                    {
                        "Mary Smith <mary@x.test>, jdoe@example.org, Who? <one@y.test>",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "Mary Smith",
                                Address = "mary@x.test",
                            },
                            new EmailAddress
                            {
                                Name = "",
                                Address = "jdoe@example.org"
                            },
                            new EmailAddress
                            {
                                Name = "Who?",
                                Address = "one@y.test",
                            }
			            }
		            },
                    new object[]
                    {
                        "<boss@nil.test>, \"Giant; \\\"Big\\\" Box\" <sysservices@example.net>",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "",
					            Address = "boss@nil.test"
				            },
                            new EmailAddress
            				{
					            Name = "Giant; \"Big\" Box",
            					Address = "sysservices@example.net"
				            }
			            }
		            },
                    // RFC 5322, Appendix A.6.1
                    new object[]
                    {
                        "Joe Q. Public <john.q.public@example.com>",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "Joe Q. Public",
                                Address = "john.q.public@example.com"
                            }
                        }
		            },
		            // RFC 5322, Appendix A.1.3
                    new object[]
		            {
                        "group1: groupaddr1@example.com;",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "",
                                Address = "groupaddr1@example.com"
                            }
                        }
                    },
                    new object[]
            		{
			            "empty group: ;",
                        new EmailAddress[] {}
                    },
                    new object[]
                    {
			            "A Group:Ed Jones <c@a.test>,joe@where.test,John <jdoe@one.test>;",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "Ed Jones",
                                Address = "c@a.test"
                            },
                            new EmailAddress
                            {
                                Name = "",
                                Address = "joe@where.test"
                            },
                            new EmailAddress
                            {
                                Name = "John",
                                Address = "jdoe@one.test"
                            }
			            }
            		},
                    new object[]
                    {
            			"Group1: <addr1@example.com>;, Group 2: addr2@example.com;, John <addr3@example.com>",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "",
                                Address = "addr1@example.com"
                            },
                            new EmailAddress
                            {
                                Name = "",
                                Address = "addr2@example.com"
                            },
                            new EmailAddress
                            {
                                Name = "John",
                                Address = "addr3@example.com"
                            }
                        }
                    },
                    // RFC 2047 "Q"-encoded ISO-8859-1 address.
                    new object[]
                    {
                        "=?iso-8859-1?q?J=F6rg_Doe?= <joerg@example.com>",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "Jörg Doe",
                                Address = "joerg@example.com"
                            }
                        }
                    },
		            // RFC 2047 "Q"-encoded US-ASCII address. Dumb but legal.
                    new object[]
                    {
                        "=?us-ascii?q?J=6Frg_Doe?= <joerg@example.com>",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "Jorg Doe",
                                Address = "joerg@example.com"
                            }
                        }
                    },
		            // RFC 2047 "Q"-encoded UTF-8 address.
                    new object[]
                    {
                        "=?utf-8?q?J=C3=B6rg_Doe?= <joerg@example.com>",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "Jörg Doe",
                                Address = "joerg@example.com"
                            }
                        }
                    },
		            // RFC 2047 "Q"-encoded UTF-8 address with multiple encoded-words.
                    new object[]
                    {
			            "=?utf-8?q?J=C3=B6rg?=  =?utf-8?q?Doe?= <joerg@example.com>",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
					            Name = "JörgDoe",
					            Address = "joerg@example.com"
				            }
                        }
                    },
		            // RFC 2047, Section 8.
                    new object[]
                    {
			            "=?ISO-8859-1?Q?Andr=E9?= Pirard <PIRARD@vm1.ulg.ac.be>",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "André Pirard",
                                Address = "PIRARD@vm1.ulg.ac.be"
                            }
                        }
                    },
		            // Custom example of RFC 2047 "B"-encoded ISO-8859-1 address.
                    new object[]
                    {
            			"=?ISO-8859-1?B?SvZyZw==?= <joerg@example.com>",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "Jörg",
                                Address = "joerg@example.com"
                            }
                        }
                    },
            		// Custom example of RFC 2047 "B"-encoded UTF-8 address.
                    new object[]
                    {
			            "=?UTF-8?B?SsO2cmc=?= <joerg@example.com>",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "Jörg",
                                Address = "joerg@example.com"
                            }
                        }
                    },
		            // Custom example with "." in name. For issue 4938
                    new object[]
                    {
			            "Asem H. <noreply@example.com>",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "Asem H.",
                                Address = "noreply@example.com"
                            }
                        }
                    },
		            // RFC 6532 3.2.3, qtext /= UTF8-non-ascii
                    new object[]
                    {
			            "\"Gø Pher\" <gopher@example.com>",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "Gø Pher",
                                Address = "gopher@example.com"
                            }
                        }
                    },
		            // RFC 6532 3.2, atext /= UTF8-non-ascii
                    new object[]
                    {
			            "µ <micro@example.com>",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "µ",
                                Address = "micro@example.com"
                            }
                        }
                    },
		            // RFC 6532 3.2.2, local address parts allow UTF-8
                    new object[]
                    {
                        "Micro <µ@example.com>",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "Micro",
                                Address = "µ@example.com"
                            }
                        }
                    },
		            // RFC 6532 3.2.4, domains parts allow UTF-8
                    new object[]
                    {
			            "Micro <micro@µ.example.com>",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "Micro",
                                Address = "micro@µ.example.com"
                            }
                        }
                    },
		            // Issue 14866
                    new object[]
                    {
			            "\"\" <emptystring@example.com>",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "",
                                Address = "emptystring@example.com"
                            }
                        }
                    },
		            // CFWS
                    new object[]
                    {
                        "<cfws@example.com> (CFWS (cfws))  (another comment)",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "",
                                Address = "cfws@example.com",
                            }
                        }
                    },
                    new object[]
                    {
                        "<cfws@example.com> ()  (another comment), <cfws2@example.com> (another)",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "",
                                Address = "cfws@example.com"
                            },
                            new EmailAddress
                            {
                                Name = "",
                                Address = "cfws2@example.com"
                            }
                        }
                    },
		            // Comment as display name
                    new object[]
                    {
			            "john@example.com (John Doe)",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "John Doe",
                                Address = "john@example.com"
                            }
                        }
                    },
		            // Comment and display name
                    new object[]
                    {
			            "John Doe <john@example.com> (Joey)",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "John Doe",
                                Address = "john@example.com"
                            }
                        }
                    },
		            // Comment as display name, no space
                    new object[]
                    {
			            "john@example.com(John Doe)",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "John Doe",
                                Address = "john@example.com"
                            }
                        }
                    },
		            // Comment as display name, Q-encoded
                    new object[]
                    {
                        "asjo@example.com (Adam =?utf-8?Q?Sj=C3=B8gren?=)",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "Adam Sjøgren",
                                Address = "asjo@example.com"
                            }
                        }
                    },
                    // Comment as display name, Q-encoded and tab-separated
                    new object[]
                    {
                        "asjo@example.com (Adam	=?utf-8?Q?Sj=C3=B8gren?=)",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "Adam Sjøgren",
                                Address = "asjo@example.com"
                            }
                        }
                    },
                    // Nested comment as display name, Q-encoded
                    new object[]
                    {
                        "asjo@example.com (Adam =?utf-8?Q?Sj=C3=B8gren?= (Debian))",
                        new EmailAddress[]
                        {
                            new EmailAddress
                            {
                                Name = "Adam Sjøgren (Debian)",
                                Address = "asjo@example.com"
                            }
                        }
                    }
                };
            }
        }
    }
}
