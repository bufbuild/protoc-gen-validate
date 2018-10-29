package com.lyft.pgv;

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
}
