package com.lyft.pgv;

import java.util.Comparator;

public final class ComparativeValidation {
    private ComparativeValidation() {
    }

    public static <T> void lessThan(String field, T value, T limit, Comparator<T> comparator) throws ValidationException {
        if (!lt(comparator.compare(value, limit))) {
            throw new ValidationException(field, value.toString() + " must be less than " + limit.toString());
        }
    }

    public static <T> void lessThanOrEqual(String field, T value, T limit, Comparator<T> comparator) throws ValidationException {
        if (!lte(comparator.compare(value, limit))) {
            throw new ValidationException(field, value.toString() + " must be less than or equal to " + limit.toString());
        }
    }

    public static <T> void greaterThan(String field, T value, T limit, Comparator<T> comparator) throws ValidationException {
        if (!gt(comparator.compare(value, limit))) {
            throw new ValidationException(field, value.toString() + " must be greater than " + limit.toString());
        }
    }

    public static <T> void greaterThanOrEqual(String field, T value, T limit, Comparator<T> comparator) throws ValidationException {
        if (!gte(comparator.compare(value, limit))) {
            throw new ValidationException(field, value.toString() + " must be greater than or equal to " + limit.toString());
        }
    }

    public static <T> void between(String field, T value, T lower, boolean lowerInclusive, T upper, boolean upperInclusive, Comparator<T> comparator) throws ValidationException {
        if (!between(value, lower, lowerInclusive, upper, upperInclusive, comparator)) {
            throw new ValidationException(field, value.toString() + " must be in the range " + range(lower, lowerInclusive, upper, upperInclusive));
        }
    }

    public static <T> void outside(String field, T value, T lower, boolean lowerInclusive, T upper, boolean upperInclusive, Comparator<T> comparator) throws ValidationException {
        if (between(value, lower, lowerInclusive, upper, upperInclusive, comparator)) {
            throw new ValidationException(field, value.toString() + " must be outside the range " + range(lower, lowerInclusive, upper, upperInclusive));
        }
    }

    private static <T> boolean between(T value, T lower, boolean lowerInclusive, T upper, boolean upperInclusive, Comparator<T> comparator) {
        return (lowerInclusive ? gte(comparator.compare(value, lower)) : gt(comparator.compare(value, lower))) &&
               (upperInclusive ? lte(comparator.compare(value, upper)) : lt(comparator.compare(value, upper)));
    }

    private static <T> String range(T lower, boolean lowerInclusive, T upper, boolean upperInclusive) {
        return (lowerInclusive ? "[" : "(") + lower.toString() + ", " + upper.toString() + (upperInclusive ? "]" : ")");
    }

    private static boolean lt(int comparatorResult) {
        return comparatorResult < 0;
    }

    private static boolean lte(int comparatorResult) {
        return comparatorResult <= 0;
    }

    private static boolean gt(int comparatorResult) {
        return comparatorResult > 0;
    }

    private static boolean gte(int comparatorResult) {
        return comparatorResult >= 0;
    }
}
