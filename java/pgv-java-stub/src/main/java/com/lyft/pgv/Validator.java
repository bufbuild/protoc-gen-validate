package com.lyft.pgv;

/**
 * {@code Validator} is the base class for all generated PGV validators.
 * @param <T>
 */
public abstract class Validator<T> {
    /**
     * Asserts validation rules on a protobuf object.
     *
     * @param proto the protobuf object to validate.
     * @throws ValidationException with the first validation error encountered.
     */
    public abstract void assertValid(T proto) throws ValidationException;

    /**
     * Checks validation rules on a protobuf object.
     *
     * @param proto the protobuf object to validate.
     * @return {@code true} if all rules are valid, {@code false} if not.
     */
    public boolean isValid(T proto) {
        try {
            assertValid(proto);
            return true;
        } catch (com.lyft.pgv.ValidationException ex) {
            return false;
        }
    }

    public static Validator ALWAYS_VALID = new Validator() {
        @Override
        public void assertValid(Object proto) {
            // Do nothing. Always valid.
        }
    };
}
