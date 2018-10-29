package com.lyft.pgv;

import com.google.protobuf.ByteString;

import com.google.re2j.Pattern;
import org.apache.commons.validator.routines.InetAddressValidator;
import org.apache.commons.validator.routines.UrlValidator;

public final class BytesValidation {
    private BytesValidation() {
    }

    public static void constant(String field, ByteString value, ByteString expected) throws ValidationException {
        if (!value.equals(expected)) {
            throw new ValidationException(field, "value must equal " + expected);
        }
    }

    public static void length(String field, ByteString value, int expected) throws ValidationException {
        if (value.size() != expected) {
            throw new ValidationException(field, "length must be " + expected);
        }
    }

    public static void minLength(String field, ByteString value, int expected) throws ValidationException {
        if (value.size() < expected) {
            throw new ValidationException(field, "length must be at least " + expected);
        }
    }

    public static void maxLength(String field, ByteString value, int expected) throws ValidationException {
        if (value.size() > expected) {
            throw new ValidationException(field, "length must be at maximum " + expected);
        }
    }

    public static void prefix(String field, ByteString value, String prefix) throws ValidationException {
        if (!value.startsWith(ByteString.copyFrom(prefix.getBytes()))) {
            throw new ValidationException(field, "should start with " + prefix);
        }
    }

    public static void contains(String field, ByteString value, String contains) throws ValidationException {
        if (!value.toStringUtf8().contains(contains)) {
            throw new ValidationException(field, "should contain " + contains);
        }
    }

    public static void suffix(String field, ByteString value, String suffix) throws ValidationException {
        if (!value.endsWith(ByteString.copyFrom(suffix.getBytes()))) {
            throw new ValidationException(field, "should end with " + suffix);
        }
    }

    public static void pattern(String field, ByteString value, String pattern) throws ValidationException {
        Pattern p = Pattern.compile(pattern);
        if (!p.matches(value.toStringUtf8())) {
            throw new ValidationException(field, "must match pattern " + pattern);
        }
    }

    public static void ip(String field, ByteString value) throws ValidationException {
        if (value.toByteArray().length != 4 && value.toByteArray().length != 16) {
            throw new ValidationException(field, "should be valid ip address " + value);
        }
    }

    public static void ipv4(String field, ByteString value) throws ValidationException {
        if (value.toByteArray().length != 4) {
            throw new ValidationException(field, "should be valid ipv4 address " + value);
        }
    }

    public static void ipv6(String field, ByteString value) throws ValidationException {
        if (value.toByteArray().length != 16) {
            throw new ValidationException(field, "should be valid ipv6 address " + value);
        }
    }
}
