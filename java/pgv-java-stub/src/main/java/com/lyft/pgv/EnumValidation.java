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
}
