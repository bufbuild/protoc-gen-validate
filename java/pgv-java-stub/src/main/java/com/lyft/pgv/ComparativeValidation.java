package com.lyft.pgv;

import java.util.Comparator;

public final class ComparativeValidation {
    private ComparativeValidation() {
    }

    public static <T> void lessThan(String field, T value, T limit, Comparator<T> comparitor) throws ValidationException {
        if (!(comparitor.compare(value, limit) < 0)) {
            throw new ValidationException(field, "value must be less than " + limit);
        }
    }

    public static <T> void lessThanOrEqual(String field, T value, T limit, Comparator<T> comparitor) throws ValidationException {
        if (!(comparitor.compare(value, limit) <= 0)) {
            throw new ValidationException(field, "value must be less than or equal to " + limit);
        }
    }

    public static <T> void greaterThan(String field, T value, T limit, Comparator<T> comparitor) throws ValidationException {
        if (!(comparitor.compare(value, limit) > 0)) {
            throw new ValidationException(field, "value must be greater than " + limit);
        }
    }

    public static <T> void greaterThanOrEqual(String field, T value, T limit, Comparator<T> comparitor) throws ValidationException {
        if (!(comparitor.compare(value, limit) >= 0)) {
            throw new ValidationException(field, "value must be greater than or equal to " + limit);
        }
    }
}
