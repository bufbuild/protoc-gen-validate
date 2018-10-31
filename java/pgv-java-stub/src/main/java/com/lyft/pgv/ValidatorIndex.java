package com.lyft.pgv;

public interface ValidatorIndex {
    <T> Validator<T> validatorFor(Class clazz);

    ValidatorIndex ALWAYS_VALID = new ValidatorIndex() {
        @Override
        public <T> Validator<T> validatorFor(Class clazz) {
            return new Validator<T>() {
                @Override
                public void assertValid(T proto) throws ValidationException {
                    // it's valid
                }
            };
        }
    };
}
