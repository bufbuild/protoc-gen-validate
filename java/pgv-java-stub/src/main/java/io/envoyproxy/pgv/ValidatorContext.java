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

    public ValidatorContext(ValidatorIndex validationIndex) {
        this.validationIndex = validationIndex;
    }

    /**
     * Package private method to make the {@code ValidatorIndex} available to the
     * ValidatorContextImpl during validation.
     * @return the {@link ValidatorIndex} used for validations conducted with
     * this context.
     */
    ValidatorIndex getValidatorIndex() {
        return validationIndex;
    }
    
    public void validate(Object obj) throws ValidationException{
        validate(obj,new ValidatorInterceptor.PassThroughInterceptor());
    }

    public void validate(Object obj, ValidatorInterceptor interceptor) throws ValidationException {
        new ValidatorExecutionContext(this, interceptor).validate(obj);
    }
    
    public boolean isValid(Object obj){
        return isValid(obj,new ValidatorInterceptor.PassThroughInterceptor());
    }

    public boolean isValid(Object obj, ValidatorInterceptor interceptor) {
        try{
            validate(obj,interceptor);
        }
        catch(ValidationException ex){
            return false;
        }
        return interceptor.isValid();
    }
    
    public void assertValid(Object obj) throws ValidationException{
        assertValid(obj,new ValidatorInterceptor.PassThroughInterceptor());
    }

    public void assertValid(Object obj, ValidatorInterceptor interceptor)  throws ValidationException {
        validate(obj,interceptor);
        if (!interceptor.isValid()){
            throw new ValidationException("UNKNOWN",obj,"Validations failed, check interceptor for details");
        }
    }
}
