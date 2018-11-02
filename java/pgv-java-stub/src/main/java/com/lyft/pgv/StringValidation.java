package com.lyft.pgv;

import com.google.re2j.Pattern;
import org.apache.commons.validator.routines.DomainValidator;
import org.apache.commons.validator.routines.EmailValidator;
import org.apache.commons.validator.routines.InetAddressValidator;

import java.net.URI;
import java.net.URISyntaxException;

public final class StringValidation {
    private StringValidation() {
    }

    public static void length(String field, String value, int expected) throws ValidationException {
        if (value.length() != expected) {
            throw new ValidationException(field, "length must be " + expected);
        }
    }

    public static void minLength(String field, String value, int expected) throws ValidationException {
        if (value.length() < expected) {
            throw new ValidationException(field, "length must be at least " + expected);
        }
    }

    public static void maxLength(String field, String value, int expected) throws ValidationException {
        if (value.length() > expected) {
            throw new ValidationException(field, "length must be at maximum " + expected);
        }
    }

    public static void lenBytes(String field, String value, int expected) throws ValidationException {
        if (value.getBytes().length != expected) {
            throw new ValidationException(field, "bytes length must be " + expected);
        }
    }

    public static void minBytes(String field, String value, int expected) throws ValidationException {
        if (value.getBytes().length < expected) {
            throw new ValidationException(field, "bytes length must be at least " + expected);
        }
    }

    public static void maxBytes(String field, String value, int expected) throws ValidationException {
        if (value.getBytes().length > expected) {
            throw new ValidationException(field, "bytes length must be at maximum " + expected);
        }
    }

    public static void pattern(String field, String value, Pattern p) throws ValidationException {
        if (!p.matches(value)) {
            throw new ValidationException(field, "must match pattern " + p.pattern());
        }
    }

    public static void prefix(String field, String value, String prefix) throws ValidationException {
        if (!value.startsWith(prefix)) {
            throw new ValidationException(field, "should start with " + prefix);
        }
    }

    public static void contains(String field, String value, String contains) throws ValidationException {
        if (!value.contains(contains)) {
            throw new ValidationException(field, "should contain " + contains);
        }
    }

    public static void suffix(String field, String value, String suffix) throws ValidationException {
        if (!value.endsWith(suffix)) {
            throw new ValidationException(field, "should end with " + suffix);
        }
    }

    public static void email(String field, String value) throws ValidationException {
        EmailValidator emailValidator = EmailValidator.getInstance();
        if (!emailValidator.isValid(value)) {
            throw new ValidationException(field, "should be valid email");
        }
    }

    public static void hostName(String field, String value) throws ValidationException {
        DomainValidator domainValidator = DomainValidator.getInstance();
        if (!domainValidator.isValid(value)) {
            throw new ValidationException(field, "should be valid host");
        }
    }

    public static void ip(String field, String value) throws ValidationException {
        InetAddressValidator ipValidator = InetAddressValidator.getInstance();
        if (!ipValidator.isValid(value)) {
            throw new ValidationException(field, "should be valid ip address");
        }
    }

    public static void ipv4(String field, String value) throws ValidationException {
        InetAddressValidator ipValidator = InetAddressValidator.getInstance();
        if (!ipValidator.isValidInet4Address(value)) {
            throw new ValidationException(field, "should be valid ipv4 address");
        }
    }

    public static void ipv6(String field, String value) throws ValidationException {
        InetAddressValidator ipValidator = InetAddressValidator.getInstance();
        if (!ipValidator.isValidInet6Address(value)) {
            throw new ValidationException(field, "should be valid ipv6 address");
        }
    }

    public static void uri(String field, String value) throws ValidationException {
        try {
            URI uri = new URI(value);
            if (!uri.isAbsolute()) {
                throw new ValidationException(field, "should be valid absolute uri");
            }
        } catch (URISyntaxException ex) {
            throw new ValidationException(field, "should be valid absolute uri");
        }
    }

    public static void uriRef(String field, String value) throws ValidationException {
        try {
            URI uri = new URI(value);
        } catch (URISyntaxException ex) {
            throw new ValidationException(field, "should be valid absolute uri");
        }
    }
}
