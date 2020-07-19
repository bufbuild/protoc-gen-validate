using System;

namespace Envoyproxy.Validator
{
    public static class StringValidation
    {
        public static bool ValidateAddress(string value)
        {
            return ValidateHostname(value) || ValidateIP(value);
        }

        public static bool ValidateEmail(string value)
        {
            string addr;

            try
            {
                var a = EmailAddressParser.Parse(value.AsSpan());
                addr = a.Address;
            }
            catch (ArgumentException)
            {
                return false;
            }

            if (addr.Length > 254)
                return false;

            var parts = addr.Split('@');

            if (parts[0].Length > 64)
                return false;

            return ValidateHostname(parts[1]);
        }

        public static bool ValidateHostname(string value)
        {
            if (value.Length == 0 || value.Length > 253)
                return false;

            if (value[value.Length - 1] == '.')
                value = value.Substring(0, value.Length - 1);

            foreach (var part in value.Split('.'))
            {
                if (part.Length == 0 || part.Length > 63)
                    return false;

                if (part[0] == '-' || part[part.Length - 1] == '-')
                    return false;

                foreach (var c in part)
                {
                    if ((c < 'A' || c > 'Z') && (c < 'a' || c > 'z') && (c < '0' || c > '9') && (c != '-'))
                        return false;
                }
            }

            return true;
        }

        public static bool ValidateIP(string value)
        {
            var result = Uri.CheckHostName(value);
            return result == UriHostNameType.IPv4 || result == UriHostNameType.IPv6;
        }

        public static bool ValidateIPv4(string value)
        {
            return Uri.CheckHostName(value) == UriHostNameType.IPv4;
        }

        public static bool ValidateIPv6(string value)
        {
            return Uri.CheckHostName(value) == UriHostNameType.IPv6;
        }

        public static bool ValidateUri(string value)
        {
            try
            {
                var url = UrlParser.Parse(value);
                return url.Scheme != null && url.Hostname != null && url.Path != null;
            }
            catch (ArgumentException)
            {
                return false;
            }
        }

        public static bool ValidateUriRef(string value)
        {
            try
            {
                var url = UrlParser.Parse(value);
                return (url.Scheme != null && url.Path != null) || url.Fragment == null;
            }
            catch (ArgumentException)
            {
                return false;
            }
        }

        public static bool ValidateUuid(string value)
        {
            return Guid.TryParseExact(value, "D", out Guid _);
        }
    }
}
