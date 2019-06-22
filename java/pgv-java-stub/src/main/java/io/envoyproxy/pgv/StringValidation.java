package io.envoyproxy.pgv;

import com.google.re2j.Matcher;
import com.google.re2j.Pattern;
import org.apache.commons.validator.routines.DomainValidator;
import org.apache.commons.validator.routines.EmailValidator;
import org.apache.commons.validator.routines.InetAddressValidator;

import java.net.URI;
import java.net.URISyntaxException;
import java.nio.charset.StandardCharsets;

/**
 * {@code StringValidation} implements PGV validation for protobuf {@code String} fields.
 */
@SuppressWarnings("WeakerAccess")
public final class StringValidation {
    private StringValidation() {
    }

    public static void length(final String field, final String value, final int expected) throws ValidationException {
        final int actual = value.codePointCount(0, value.length());
        if (actual != expected) {
            throw new ValidationException(field, enquote(value), "length must be " + expected + " but got: " + actual);
        }
    }

    public static void minLength(final String field, final String value, final int expected) throws ValidationException {
        final int actual = value.codePointCount(0, value.length());
        if (actual < expected) {
            throw new ValidationException(field, enquote(value), "length must be " + expected + " but got: " + actual);
        }
    }

    public static void maxLength(final String field, final String value, final int expected) throws ValidationException {
        final int actual = value.codePointCount(0, value.length());
        if (actual > expected) {
            throw new ValidationException(field, enquote(value), "length must be " + expected + " but got: " + actual);
        }
    }

    public static void lenBytes(String field, String value, int expected) throws ValidationException {
        if (value.getBytes(StandardCharsets.UTF_8).length != expected) {
            throw new ValidationException(field, enquote(value), "bytes length must be " + expected);
        }
    }

    public static void minBytes(String field, String value, int expected) throws ValidationException {
        if (value.getBytes(StandardCharsets.UTF_8).length < expected) {
            throw new ValidationException(field, enquote(value), "bytes length must be at least " + expected);
        }
    }

    public static void maxBytes(String field, String value, int expected) throws ValidationException {
        if (value.getBytes(StandardCharsets.UTF_8).length > expected) {
            throw new ValidationException(field, enquote(value), "bytes length must be at maximum " + expected);
        }
    }

    public static void pattern(String field, String value, Pattern p) throws ValidationException {
        if (!p.matches(value)) {
            throw new ValidationException(field, enquote(value), "must match pattern " + p.pattern());
        }
    }

    public static void prefix(String field, String value, String prefix) throws ValidationException {
        if (!value.startsWith(prefix)) {
            throw new ValidationException(field, enquote(value), "should start with " + prefix);
        }
    }

    public static void contains(String field, String value, String contains) throws ValidationException {
        if (!value.contains(contains)) {
            throw new ValidationException(field, enquote(value), "should contain " + contains);
        }
    }

    public static void suffix(String field, String value, String suffix) throws ValidationException {
        if (!value.endsWith(suffix)) {
            throw new ValidationException(field, enquote(value), "should end with " + suffix);
        }
    }

    private static final Pattern emailWithDisplayName = Pattern.compile(".*<(.*)>");
    public static void email(String field, String value) throws ValidationException {
        EmailValidator emailValidator = EmailValidator.getInstance(true, true);

        // extract email address from between angle brackets, if present
        Matcher matcher = emailWithDisplayName.matcher(value);
        if (matcher.matches()) {
            value = matcher.group(1);
        }

        if (!emailValidator.isValid(value)) {
            throw new ValidationException(field, enquote(value), "should be a valid email");
        }
    }

    public static void address(String field, String value) throws ValidationException {
        boolean validHost = isAscii(value) && DomainValidator.getInstance(true).isValid(value);
        boolean validIp = InetAddressValidator.getInstance().isValid(value);

        if (!validHost && !validIp) {
            throw new ValidationException(field, enquote(value), "should be a valid host, or an ip address.");
        }
    }

    public static void hostName(String field, String value) throws ValidationException {
        if (!isAscii(value)) {
            throw new ValidationException(field, enquote(value), "should be a valid host containing only ascii characters");
        }

        DomainValidator domainValidator = DomainValidator.getInstance(true);
        if (!domainValidator.isValid(value)) {
            throw new ValidationException(field, enquote(value), "should be a valid host");
        }
    }

    public static void ip(String field, String value) throws ValidationException {
        InetAddressValidator ipValidator = InetAddressValidator.getInstance();
        if (!ipValidator.isValid(value)) {
            throw new ValidationException(field, enquote(value), "should be a valid ip address");
        }
    }

    public static void ipv4(String field, String value) throws ValidationException {
        InetAddressValidator ipValidator = InetAddressValidator.getInstance();
        if (!ipValidator.isValidInet4Address(value)) {
            throw new ValidationException(field, enquote(value), "should be a valid ipv4 address");
        }
    }

    public static void ipv6(String field, String value) throws ValidationException {
        InetAddressValidator ipValidator = InetAddressValidator.getInstance();
        if (!ipValidator.isValidInet6Address(value)) {
            throw new ValidationException(field, enquote(value), "should be a valid ipv6 address");
        }
    }

    public static void uri(String field, String value) throws ValidationException {
        try {
            URI uri = new URI(value);
            if (!uri.isAbsolute()) {
                throw new ValidationException(field, enquote(value), "should be a valid absolute uri");
            }
        } catch (URISyntaxException ex) {
            throw new ValidationException(field, enquote(value), "should be a valid absolute uri");
        }
    }

    public static void uriRef(String field, String value) throws ValidationException {
        try {
            new URI(value);
        } catch (URISyntaxException ex) {
            throw new ValidationException(field, enquote(value), "should be a valid absolute uri");
        }
    }

    private static final Pattern uuidPattern = Pattern.compile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$");
    public static void uuid(String field, String value) throws ValidationException {
        Matcher matcher = uuidPattern.matcher(value);
        if (!matcher.matches()) {
            throw new ValidationException(field, enquote(value), "should be a valid uuid");
        }
    }

    private static String enquote(String value) {
        return "\"" + value + "\"";
    }

    private static boolean isAscii(final String value) {
        for (char c : value.toCharArray()) {
            if (c > 127) {
                return false;
            }
        }
        return true;
    }
}
