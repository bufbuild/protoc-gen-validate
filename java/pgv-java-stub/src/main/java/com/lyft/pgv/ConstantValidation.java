package com.lyft.pgv;

public final class ConstantValidation {
    private ConstantValidation() {
    }

    public static <T> void constant(String field, T value, T expected) throws ValidationException {
        if (!value.equals(expected)) {
            throw new ValidationException(field, "value must equal " + expected);
        }
    }
}
