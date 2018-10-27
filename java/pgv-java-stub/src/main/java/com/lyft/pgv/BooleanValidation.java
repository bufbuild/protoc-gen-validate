package com.lyft.pgv;

public final class BooleanValidation {
    private BooleanValidation() { }

    public static void constant(String field, boolean value, boolean expected) throws ValidationException {
        if (value != expected) {
            throw new ValidationException(field, "value must equal " + expected);
        }
    }
}
