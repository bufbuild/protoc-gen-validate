using System;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using Envoyproxy.Validator;

namespace Validator.Tests
{
    [TestClass]
    public class WordDecoderTest
    {
        [TestMethod]
        [DataRow("=?UTF-8?A?Test?=")]
        [DataRow("=?UTF-8?Q?A=B?=")]
        [DataRow("=?UTF-8?Q?=A?=")]
        [DataRow("=?UTF-8?A?A?=")]
        [DataRow("=????=")]
        [DataRow("=?UTF-8???=")]
        public void CheckError(string value)
        {
            Assert.ThrowsException<ArgumentException>(() => WordDecoder.Decode(value));
        }

        [TestMethod]
        [DataRow("=?UTF-8?Q?=C2=A1Hola,_se=C3=B1or!?=", "¡Hola, señor!")]
        [DataRow("=?UTF-8?Q?Fran=C3=A7ois-J=C3=A9r=C3=B4me?=", "François-Jérôme")]
        [DataRow("=?UTF-8?q?ascii?=", "ascii")]
        [DataRow("=?utf-8?B?QW5kcsOp?=", "André")]
        [DataRow("=?ISO-8859-1?Q?Rapha=EBl_Dupont?=", "Raphaël Dupont")]
        [DataRow("=?utf-8?b?IkFudG9uaW8gSm9zw6kiIDxqb3NlQGV4YW1wbGUub3JnPg==?=", "\"Antonio José\" <jose@example.org>")]
        [DataRow("=?UTF-8?Q??=", "")]
        public void CheckResult(string value, string expected)
        {
            Assert.AreEqual(expected, WordDecoder.Decode(value));
        }
    }
}
