package com.lyft.pgv;

import com.google.common.primitives.Bytes;
import com.google.protobuf.ByteString;
import com.google.re2j.Pattern;

import java.util.Arrays;

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

    public static void prefix(String field, ByteString value, byte[] prefix) throws ValidationException {
        if (!value.startsWith(ByteString.copyFrom(prefix))) {
            throw new ValidationException(field, "should start with " + Arrays.toString(prefix));
        }
    }

    public static void contains(String field, ByteString value, byte[] contains) throws ValidationException {
        if (Bytes.indexOf(value.toByteArray(), contains) == -1) {
            throw new ValidationException(field, "should contain " + Arrays.toString(contains));
        }
    }

    public static void suffix(String field, ByteString value, byte[] suffix) throws ValidationException {
        if (!value.endsWith(ByteString.copyFrom(suffix))) {
            throw new ValidationException(field, "should end with " + Arrays.toString(suffix));
        }
    }

    public static void pattern(String field, ByteString value, Pattern p) throws ValidationException {
        if (!p.matches(value.toStringUtf8())) {
            throw new ValidationException(field, "must match pattern " + p.pattern());
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

    public static void in(String field, ByteString value, ByteString[] set) throws ValidationException {
        for (ByteString bs : set) {
            if (value.equals(bs)) {
                return;
            }
        }

        throw new ValidationException(field, "value must be in " + Arrays.toString(set));
    }

    public static void notIn(String field, ByteString value, ByteString[] set) throws ValidationException {
        for (ByteString bs : set) {
            if (value.equals(bs)) {
                throw new ValidationException(field, "value must not be in " + Arrays.toString(set));
            }
        }
    }
}
