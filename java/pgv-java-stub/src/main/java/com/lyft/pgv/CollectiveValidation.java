package com.lyft.pgv;

import java.util.Arrays;

public final class CollectiveValidation {
    private CollectiveValidation() {
    }

    public static <T> void in(String field, T value, T[] set) throws ValidationException {
        for (T i : set) {
            if (value.equals(i)) {
                return;
            }
        }

        throw new ValidationException(field, "value must be in " + Arrays.toString(set));
    }

    public static <T> void notIn(String field, T value, T[] set) throws ValidationException {
        for (T i : set) {
            if (value.equals(i)) {
                throw new ValidationException(field, "value must not be in " + Arrays.toString(set));
            }
        }
    }
}
