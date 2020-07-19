using System;
using System.Collections.Generic;
using System.Text;

namespace Envoyproxy.Validator
{
    /// RFC 2047 word decoder
    /// Re-Implementation of Golang's mime/Decode
    /// (https://golang.org/src/mime/encodedword.go)
    public static class WordDecoder
    {
        public static string Decode(string value)
        {
            return Decode(value.AsSpan());
        }

        public static string Decode(ReadOnlySpan<char> value)
        {
            // See https://tools.ietf.org/html/rfc2047#section-2 for details.
            // Our decoder is permissive, we accept empty encoded-text.
            if (value.Length < 8 || !value.StartsWith("=?".AsSpan()) || !value.EndsWith("?=".AsSpan()) || Count(value, '?') != 4)
                throw new ArgumentException("invalid RFC 2047 encoded-word");

            value = value.Slice(2, value.Length - 4);

            // split delimits the first 2 fields
            // NOTE: guaranteed to find '?' as we just checked that there are 4 occurrences
            var split = value.IndexOf('?');

	        // split word "UTF-8?q?ascii" into "UTF-8", 'q', and "ascii"
	        var charset = value.Slice(0, split);
            if (charset.IsEmpty || value.Length < split + 3)
                throw new ArgumentException("invalid RFC 2047 encoded-word");

	        var encoding = value[split + 1];

            // the field after split must only be one byte
            if (value[split + 2] != '?')
                throw new ArgumentException("invalid RFC 2047 encoded-word");

            var text = value.Slice(split + 3);

            byte[] content;

            switch (encoding)
            {
                case 'B':
                case 'b':
                    // Converter from ReadOnlySpan<char> requires netstandard2.1
                    content = Convert.FromBase64String(text.ToString());
                    break;

                case 'Q':
                case 'q':
                    content = QDecode(text);
                    break;

                default:
                    throw new ArgumentException("invalid RFC 2047 encoded-word");
            }

            Encoding enc;

            try
            {
                enc = Encoding.GetEncoding(charset.ToString());
            }
            catch (ArgumentException)
            {
                throw new NotSupportedException();
            }

            return enc.GetString(content);
        }

        // QDecode decodes a Q encoded string.
        private static byte[] QDecode(ReadOnlySpan<char> text)
        {
            var decoded = new List<byte>();

            for (int i = 0; i < text.Length; ++i)
            {
                var c = text[i];

                if (c == '_')
                {
                    decoded.Add(Convert.ToByte(' '));
                }
                else if (c == '=')
                {
                    if (i + 2 >= text.Length)
                        throw new ArgumentException("invalid RFC 2047 encoded-word");

                    // Int32.Parse with ReadOnlySpan<char> requires netstandard2.1
                    decoded.Add(Convert.ToByte(ReadHexByte(text[i + 1], text[i + 2])));
                    i += 2;
                }
                else if ((c <= '~' && c >= ' ') || c == '\n' || c == '\r' || c == '\t')
                {
                    decoded.Add(Convert.ToByte(c));
                }
                else
                {
                    throw new ArgumentException("invalid RFC 2047 encoded-word");
                }
            }

            return decoded.ToArray();
        }

        private static int Count(ReadOnlySpan<char> span, char value)
        {
            var count = 0;

            foreach (var v in span)
            {
                if (v == value)
                    ++count;
            }

            return count;
        }

        private static int ReadHexByte(char a, char b)
        {
            return (FromHex(a) << 4) | FromHex(b);
        }

        private static int FromHex(char c)
        {
            if (c >= '0' && c <= '9')
                return c - '0';

            if (c >= 'A' && c <= 'F')
                return c - 'A' + 10;

            // Accept badly encoded bytes.
            if (c >= 'a' && c <= 'f')
                return c - 'a' - 10;

            throw new ArgumentException($"invalid hex byte {Convert.ToByte(c).ToString("X2")}");
        }
    }
}
