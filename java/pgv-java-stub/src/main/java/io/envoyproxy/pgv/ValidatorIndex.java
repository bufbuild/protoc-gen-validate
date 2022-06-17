package io.envoyproxy.pgv;

/**
 * {@code ValidatorIndex} defines the point for finding {@link Validator} instances for a given type.
 */
@FunctionalInterface
public interface ValidatorIndex {
    /**
     * Returns the validator for {@code clazz}, or {@code ALWAYS_VALID} if not found.
     */
    <T> Validator<T> validatorFor(Class clazz,ValidatorExecutionContext context);

    /**
     * Returns the validator for {@code <T>}, or {@code ALWAYS_VALID} if not found.
     */
    @SuppressWarnings("unchecked")
    default <T> Validator<T> validatorFor(Object instance,ValidatorExecutionContext context) {
        return validatorFor(instance == null ? null : instance.getClass(),context);
    }

    @SuppressWarnings("unchecked")
    default <T> Validator<T> validatorFor(Class clazz) {
        ValidatorContext context = new ValidatorContext(this);
        return validatorFor(clazz == null ? null : clazz,new ValidatorExecutionContext(context, ValidatorInterceptor.BACKWARD_COMPATIBILITY_INTERCEPTOR));
    }

    @SuppressWarnings("unchecked")
    default <T> Validator<T> validatorFor(Object instance) {
        ValidatorContext context = new ValidatorContext(this);
        return validatorFor(instance == null ? null : instance.getClass(),new ValidatorExecutionContext(context, ValidatorInterceptor.BACKWARD_COMPATIBILITY_INTERCEPTOR));
    }

    ValidatorIndex ALWAYS_VALID = new ValidatorIndex() {
        @Override
        @SuppressWarnings("unchecked")
        public <T> Validator<T> validatorFor(Class clazz,ValidatorExecutionContext context) {
            return Validator.ALWAYS_VALID;
        }
    };

    ValidatorIndex ALWAYS_INVALID = new ValidatorIndex() {
        @Override
        @SuppressWarnings("unchecked")
        public <T> Validator<T> validatorFor(Class clazz,ValidatorExecutionContext context) {
            return Validator.ALWAYS_INVALID;
        }
    };
}
