package com.lyft.pgv;

import java.util.concurrent.ConcurrentHashMap;

/**
 * {@code ExplicitValidatorIndex} is an explicit registry of {@link Validator} instances. If no validator is registered
 * for {@code type}, a default {@code ALWAYS_VALID} validator will be used.
 */
public class ExplicitValidatorIndex implements ValidatorIndex {
    private final ConcurrentHashMap<Class, Validator> VALIDATOR_INDEX = new ConcurrentHashMap<>();

    /**
     * Adds a {@link Validator} to the set of known validators.
     * @param forType the type to validate
     * @param validator the validator to apply
     * @return this
     */
    public <T> ExplicitValidatorIndex add(Class<T> forType, Validator<T> validator) {
        VALIDATOR_INDEX.put(forType, validator);
        return this;
    }

    @SuppressWarnings("unchecked")
    public <T> Validator<T> validatorFor(Class clazz) {
        return VALIDATOR_INDEX.getOrDefault(clazz, Validator.ALWAYS_VALID);
    }
}
