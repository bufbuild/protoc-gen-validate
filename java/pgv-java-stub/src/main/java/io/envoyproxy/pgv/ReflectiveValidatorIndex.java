package io.envoyproxy.pgv;

import java.util.concurrent.ConcurrentHashMap;

/**
 * {@code ReflectiveValidatorIndex} uses reflection to discover {@link Validator} implementations lazily the first
 * time a type is validated. If no validator can be found for {@code type}, a default {@code ALWAYS_VALID} validator
 * will be used.
 */
public final class ReflectiveValidatorIndex implements ValidatorIndex {
    private final ConcurrentHashMap<Class, Validator> VALIDATOR_INDEX = new ConcurrentHashMap<>();

    /**
     * Retuns the validator for {@code <T>}, or {@code ALWAYS_VALID} if not found.
     */
    @Override
    @SuppressWarnings("unchecked")
    public <T> Validator<T> validatorFor(Class clazz) {
        return VALIDATOR_INDEX.computeIfAbsent(clazz, c -> {
            try {
                return reflectiveValidatorFor(c);
            } catch (ReflectiveOperationException ex) {
                return Validator.ALWAYS_VALID;
            }
        });
    }

    @SuppressWarnings("unchecked")
    private Validator reflectiveValidatorFor(Class clazz) throws ReflectiveOperationException {
        Class enclosingClass = clazz;
        while (enclosingClass.getEnclosingClass() != null) {
            enclosingClass = enclosingClass.getEnclosingClass();
        }

        String validatorClassName = enclosingClass.getName() + "Validator";
        Class validatorClass = clazz.getClassLoader().loadClass(validatorClassName);
        ValidatorImpl impl = (ValidatorImpl) validatorClass.getDeclaredMethod("validatorFor", Class.class).invoke(null, clazz);

        return proto -> impl.assertValid(proto, ReflectiveValidatorIndex.this);
    }
}
