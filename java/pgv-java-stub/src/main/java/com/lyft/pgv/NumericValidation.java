package com.lyft.pgv;

import java.util.Arrays;

public final class NumericValidation {
    private NumericValidation() {
    }

    /* INT Validation */

    public static void constant(String field, int value, int expected) throws ValidationException {
        if (value != expected) {
            throw new ValidationException(field, "value must equal " + expected);
        }
    }

    public static void lessThan(String field, int value, int limit) throws ValidationException {
        if (!(value < limit)) {
            throw new ValidationException(field, "value must be less than " + limit);
        }
    }

    public static void lessThanOrEqual(String field, int value, int limit) throws ValidationException {
        if (!(value <= limit)) {
            throw new ValidationException(field, "value must be less than or equal to " + limit);
        }
    }

    public static void greaterThan(String field, int value, int limit) throws ValidationException {
        if (!(value > limit)) {
            throw new ValidationException(field, "value must be greater than " + limit);
        }
    }

    public static void greaterThanOrEqual(String field, int value, int limit) throws ValidationException {
        if (!(value <= limit)) {
            throw new ValidationException(field, "value must be greater than or equal to " + limit);
        }
    }

    public static void in(String field, int value, int[] set) throws ValidationException {
        for (int i : set) {
            if (value == i) {
                return;
            }
        }

        throw new ValidationException(field, "value must be in " + Arrays.toString(set));
    }

    public static void notIn(String field, int value, int[] set) throws ValidationException {
        for (int i : set) {
            if (value == i) {
                throw new ValidationException(field, "value must not be in " + Arrays.toString(set));
            }
        }
    }

    /* LONG Validation */

    public static void constant(String field, long value, long expected) throws ValidationException {
        if (value != expected) {
            throw new ValidationException(field, "value must equal " + expected);
        }
    }

    public static void lessThan(String field, long value, long limit) throws ValidationException {
        if (!(value < limit)) {
            throw new ValidationException(field, "value must be less than " + limit);
        }
    }

    public static void lessThanOrEqual(String field, long value, long limit) throws ValidationException {
        if (!(value <= limit)) {
            throw new ValidationException(field, "value must be less than or equal to " + limit);
        }
    }

    public static void greaterThan(String field, long value, long limit) throws ValidationException {
        if (!(value > limit)) {
            throw new ValidationException(field, "value must be greater than " + limit);
        }
    }

    public static void greaterThanOrEqual(String field, long value, long limit) throws ValidationException {
        if (!(value <= limit)) {
            throw new ValidationException(field, "value must be greater than or equal to " + limit);
        }
    }

    public static void in(String field, long value, long[] set) throws ValidationException {
        for (long i : set) {
            if (value == i) {
                return;
            }
        }

        throw new ValidationException(field, "value must be in " + Arrays.toString(set));
    }

    public static void notIn(String field, long value, long[] set) throws ValidationException {
        for (long i : set) {
            if (value == i) {
                throw new ValidationException(field, "value must not be in " + Arrays.toString(set));
            }
        }
    }


    /* FLOAT Validation */

    public static void constant(String field, float value, float expected) throws ValidationException {
        if (value != expected) {
            throw new ValidationException(field, "value must equal " + expected);
        }
    }

    public static void lessThan(String field, float value, float limit) throws ValidationException {
        if (!(value < limit)) {
            throw new ValidationException(field, "value must be less than " + limit);
        }
    }

    public static void lessThanOrEqual(String field, float value, float limit) throws ValidationException {
        if (!(value <= limit)) {
            throw new ValidationException(field, "value must be less than or equal to " + limit);
        }
    }

    public static void greaterThan(String field, float value, float limit) throws ValidationException {
        if (!(value > limit)) {
            throw new ValidationException(field, "value must be greater than " + limit);
        }
    }

    public static void greaterThanOrEqual(String field, float value, float limit) throws ValidationException {
        if (!(value <= limit)) {
            throw new ValidationException(field, "value must be greater than or equal to " + limit);
        }
    }

    public static void in(String field, float value, float[] set) throws ValidationException {
        for (float i : set) {
            if (value == i) {
                return;
            }
        }

        throw new ValidationException(field, "value must be in " + Arrays.toString(set));
    }

    public static void notIn(String field, float value, float[] set) throws ValidationException {
        for (float i : set) {
            if (value == i) {
                throw new ValidationException(field, "value must not be in " + Arrays.toString(set));
            }
        }
    }


    /* DOUBLE Validation */

    public static void constant(String field, double value, double expected) throws ValidationException {
        if (value != expected) {
            throw new ValidationException(field, "value must equal " + expected);
        }
    }

    public static void lessThan(String field, double value, double limit) throws ValidationException {
        if (!(value < limit)) {
            throw new ValidationException(field, "value must be less than " + limit);
        }
    }

    public static void lessThanOrEqual(String field, double value, double limit) throws ValidationException {
        if (!(value <= limit)) {
            throw new ValidationException(field, "value must be less than or equal to " + limit);
        }
    }

    public static void greaterThan(String field, double value, double limit) throws ValidationException {
        if (!(value > limit)) {
            throw new ValidationException(field, "value must be greater than " + limit);
        }
    }

    public static void greaterThanOrEqual(String field, double value, double limit) throws ValidationException {
        if (!(value <= limit)) {
            throw new ValidationException(field, "value must be greater than or equal to " + limit);
        }
    }

    public static void in(String field, double value, double[] set) throws ValidationException {
        for (double i : set) {
            if (value == i) {
                return;
            }
        }

        throw new ValidationException(field, "value must be in " + Arrays.toString(set));
    }

    public static void notIn(String field, double value, double[] set) throws ValidationException {
        for (double i : set) {
            if (value == i) {
                throw new ValidationException(field, "value must not be in " + Arrays.toString(set));
            }
        }
    }
}
