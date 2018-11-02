package com.lyft.pgv;

import java.util.Arrays;

public final class NumericValidation {
    private NumericValidation() {
    }

    /* INT Validation */

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
