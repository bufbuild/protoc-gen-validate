package com.lyft.pgv;

import com.google.protobuf.ProtocolMessageEnum;

import java.util.Arrays;

public final class EnumValidation {
    private EnumValidation() {
    }

    public static void definedOnly(String field, ProtocolMessageEnum value) throws ValidationException {
        if (value.toString().equals("UNRECOGNIZED")) {
            throw new ValidationException(field, "value is not a defined Enum value " + value);
        }
    }

    public static void in(String field, ProtocolMessageEnum value, ProtocolMessageEnum[] enumValues) throws ValidationException {
        if (enumValues != null) {
            for (ProtocolMessageEnum enumValue : enumValues) {
                if (value.equals(enumValue)) {
                    return;
                }
            }
        }

        throw new ValidationException(field, "value must be in " + Arrays.toString(enumValues));
    }

    public static void notIn(String field, ProtocolMessageEnum value, ProtocolMessageEnum[] enumValues) throws ValidationException {
        if (enumValues != null) {
            for (ProtocolMessageEnum enumValue : enumValues) {
                if (value.equals(enumValue)) {
                    throw new ValidationException(field, "value must not be in " + Arrays.toString(enumValues));
                }
            }
        }
    }
}
