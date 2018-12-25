package com.lyft.pgv;

/**
 * {@code ValidatorIndex} defines the entry point for finding {@link Validator} instances for a given type.
 */
public interface ValidatorIndex {
    <T> Validator<T> validatorFor(Class clazz);

    /**
     * Retuns the validator for {@code <T>}, or {@code ALWAYS_VALID} if not found.
     */
    @SuppressWarnings("unchecked")
    default <T> Validator<T> validatorFor(Object instance) {
        return validatorFor(instance == null ? null : instance.getClass());
    }

    ValidatorIndex ALWAYS_VALID = new ValidatorIndex() {
        @Override
        @SuppressWarnings("unchecked")
        public <T> Validator<T> validatorFor(Class clazz) {
            return Validator.ALWAYS_VALID;
        }
    };
}
