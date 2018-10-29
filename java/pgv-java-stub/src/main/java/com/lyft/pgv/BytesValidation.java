package com.lyft.pgv;

import com.google.protobuf.ByteString;

import org.apache.commons.validator.routines.InetAddressValidator;
import org.apache.commons.validator.routines.UrlValidator;

public final class BytesValidation {
    private BytesValidation() { }

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
      if (!value.toString().contains(contains)) {
               throw new ValidationException(field, "should contain " + contains);
      }
    }

    public static void suffix(String field, ByteString value, String suffix) throws ValidationException {
      if (!value.endsWith(ByteString.copyFrom(suffix.getBytes()))) {
               throw new ValidationException(field, "should end with " + suffix);
      }
    }


    public static void ip(String field, ByteString value) throws ValidationException {
      InetAddressValidator ipValidator = InetAddressValidator.getInstance();
      if (!ipValidator.isValid(value.toString())) {
               throw new ValidationException(field, "should be valid ip address " + value);
      }
    }

    public static void ipv4(String field, ByteString value) throws ValidationException {
      InetAddressValidator ipValidator = InetAddressValidator.getInstance();
      if (!ipValidator.isValidInet4Address(value.toString())) {
               throw new ValidationException(field, "should be valid ipv4 address " + value);
      }
    }

    public static void ipv6(String field, ByteString value) throws ValidationException {
      InetAddressValidator ipValidator = InetAddressValidator.getInstance();
      if (!ipValidator.isValidInet6Address(value.toString())) {
               throw new ValidationException(field, "should be valid ipv6 address " + value);
      }
    }
}
