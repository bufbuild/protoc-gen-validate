package com.lyft.pgv;

import java.util.Arrays;

public final class EnumValidation {
    private EnumValidation() {
    }

    public static void constant(String field, Object value, Object expected) throws ValidationException {
        if (!value.equals(expected)) {
            throw new ValidationException(field, "value must equal " + expected);
        }
    }

    public static void definedOnly(String field, Object value) throws ValidationException {
        Object[] enumValues = value.getClass().getEnumConstants();
        if (enumValues != null) {
            for (Object enumValue : enumValues) {
                if (((Enum) value).name().equals(enumValue)) {
                    return;
                }
            }
        }
        throw new ValidationException(field, "value is not a defined Enum value " + value);
    }

    public static void in(String field, Object value, Object[] enumValues) throws ValidationException {
        if (enumValues != null) {
            for (Object enumValue : enumValues) {
                if (value.equals(enumValue)) {
                    return;
                }
            }
        }

        throw new ValidationException(field, "value must be in " + Arrays.toString(enumValues));
    }

    public static void notIn(String field, Object value, Object[] enumValues) throws ValidationException {
        if (enumValues != null) {
            for (Object enumValue : enumValues) {
                if (!value.equals(enumValue)) {
                    throw new ValidationException(field, "value must not be in " + Arrays.toString(enumValues));
                }
            }
        }
    }
}
