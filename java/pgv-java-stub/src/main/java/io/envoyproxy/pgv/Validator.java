package io.envoyproxy.pgv;

/**
 * {@code Validator} asserts the validity of a protobuf object.
 * @param <T> The type to validate
 */
@FunctionalInterface
public interface Validator<T> {
    /**
     * Asserts validation rules on a protobuf object.
     *
     * @deprecated assertion style validation enforces strict semantics which 
     * don't translate well to different validation error handling strategies.  
     * Use {@link #validate} method instead.
     * @param proto the protobuf object to validate.
     * @throws ValidationException with the first validation error encountered.
     */
    default void assertValid(T proto) throws ValidationException{
        validate(proto);
    }

    
    void validate(T proto) throws ValidationException;
    /**
     * Checks validation rules on a protobuf object.
     *
     * @param proto the protobuf object to validate.
     * @return {@code true} if all rules are valid, {@code false} if not.
     */
    default boolean isValid(T proto) {
        try {
            assertValid(proto);
            return true;
        } catch (io.envoyproxy.pgv.ValidationException ex) {
            return false;
        }
    }

    Validator ALWAYS_VALID = (proto) -> {
        // Do nothing. Always valid.
    };

    Validator ALWAYS_INVALID = (proto) -> {
        throw new ValidationException("UNKNOWN", new Object(), "Explicitly invalid");
    };
}
