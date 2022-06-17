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
 * {@code ValidatorExcecutionContext} is used to conduct and coordinate a single 
 * validation run.  It becomes finalized and inert at the end validation.
 * This class is made available to generated code only.
 * @author rlamont
 */
public class ValidatorExecutionContext {
    private final ValidatorContext context;
    private final ValidatorInterceptor interceptor;

    /**
     * package private constructor.
     * @param context
     * @param interceptor 
     */
    ValidatorExecutionContext(ValidatorContext context, ValidatorInterceptor interceptor) {
        this.context = context;
        this.interceptor = interceptor;
    }
    
    public ValidatorInterceptor getValidatorInterceptor() {
        return interceptor;
    }

    
    @SuppressWarnings("unchecked")
    public <T> Validator<T> validatorFor(Class<T> clazz) {
        return context.getValidatorIndex().validatorFor(clazz, this);
    }
    
    @SuppressWarnings("unchecked")
    public <T> Validator<T> validatorFor(Object obj) {
        return context.getValidatorIndex().validatorFor(obj==null?null:obj.getClass(), this);
    }
    
    @SuppressWarnings("unchecked")
    <T> void validate(T obj) throws ValidationException, IllegalStateException{
        if (!interceptor.isValid()){
            throw new IllegalStateException("This validation context has expired: interceptor is stale");
        }
        validatorFor(obj).validate(obj);
    }
}
