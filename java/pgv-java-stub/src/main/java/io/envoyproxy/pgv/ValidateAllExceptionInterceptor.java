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

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

/**
 * {@code ValidateAllExceptionInterceptor} provides for the validator to perform
 * all validations and collect all exceptions for later reporting. Exceptions are 
 * collected into a list which can be retrieved after validation is complete.
 * @author rlamont
 */
public class ValidateAllExceptionInterceptor implements ValidatorInterceptor {
    private final List<ValidationException> exceptionList = new ArrayList<>();
    @Override
    @SuppressWarnings("unchecked")
    public void validate(Validator validator, Object proto) throws ValidationException {
        try{
            validator.assertValid(proto);
        }
        catch(ValidationException ex){
            exceptionList.add(ex);
        }
    }
    
    /**
     * Retrieve the results of validation. 
     * @return the {@code List} of {@link ValidationException} objects after 
     * validation has been performed.  An empty list implies that the {@code Message}
     * was valid
     */
    public List<ValidationException> getAllValidationExceptions(){
        return Collections.unmodifiableList(exceptionList);
    }
    
    /**
     * Convenience method to quickly deteremine whether validation was successful
     * or not.
     * @return true if no {@link ValidationException} objects were created
     * during validation.
     */
    public boolean isValid(){
        return exceptionList.isEmpty();
    }
}
