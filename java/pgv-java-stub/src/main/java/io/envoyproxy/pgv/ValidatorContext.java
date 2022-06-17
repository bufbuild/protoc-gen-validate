/*
 * Copyright 2022 rlamont.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package io.envoyproxy.pgv;

/**
 * {@code ValidatorContext} defines the principal entry point for conducting 
 * validations. The context is available whilst the message is being walked and
 * provides for various behaviors including Validator lookup (via 
 * {@link ValidatorIndex}) and interceptor injection (via {@link ValidatorInterceptor}.
 * 
 * @author rlamont
 */
public class ValidatorContext {
    public final ValidatorIndex validationIndex;
    public final ValidatorInterceptor validationInterceptor;

    public ValidatorContext(ValidatorIndex validationIndex, ValidatorInterceptor validationInterceptor) {
        this.validationIndex = validationIndex;
        this.validationInterceptor = validationInterceptor;
    }

    public ValidatorContext(ValidatorIndex validationIndex) {
        this(validationIndex,ValidatorInterceptor.PASS_THROUGH);
    }

    public ValidatorIndex getValidatorIndex() {
        return validationIndex;
    }

    public ValidatorInterceptor getValidatorInterceptor() {
        return validationInterceptor;
    }

    /**
     * Returns the validator for {@code clazz} from the ValidatorIndex, or 
     * {@code ALWAYS_VALID} if not found.
     */
    public <T> Validator<T> validatorFor(Class<T> clazz) {
        return validationIndex.validatorFor(clazz, this);
    }
    
    public <T> Validator<T> validatorFor(T object) {
        return validationIndex.validatorFor(object, this);
    }
    
    
}
