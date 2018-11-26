package com.lyft.pgv;

import java.util.concurrent.ConcurrentHashMap;

public final class ReflectiveValidatorIndex implements ValidatorIndex {
    private ReflectiveValidatorIndex() {
    }

    private static final ReflectiveValidatorIndex INSTANCE = new ReflectiveValidatorIndex();

    private final ConcurrentHashMap<Class, Validator> VALIDATOR_INDEX = new ConcurrentHashMap<>();

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
    private static Validator reflectiveValidatorFor(Class clazz) throws ReflectiveOperationException {
        Class enclosingClass = clazz;
        while (enclosingClass.getEnclosingClass() != null) {
            enclosingClass = enclosingClass.getEnclosingClass();
        }

        String validatorClassName = enclosingClass.getName() + "Validator";
        Class validatorClass = clazz.getClassLoader().loadClass(validatorClassName);
        return (Validator) validatorClass.getDeclaredMethod("validatorFor", Class.class).invoke(null, clazz);
    }

    @SuppressWarnings("unchecked")
    public static <T> Validator<T> validatorFor(Object instance) {
        return INSTANCE.validatorFor(instance == null ? null : instance.getClass());
    }
}
