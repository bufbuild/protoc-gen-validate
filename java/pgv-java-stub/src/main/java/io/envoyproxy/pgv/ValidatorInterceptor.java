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
 * {@code ValidatorInterceptor} provides a means for injecting behavior into the
 * validation chain.  It wraps calls to any Validator instance.
 * @author rlamont
 */
@FunctionalInterface
public interface ValidatorInterceptor {
    
    void assertValid(Validator validator,Object proto ) throws ValidationException;
    
    @SuppressWarnings("unchecked")
    ValidatorInterceptor PASS_THROUGH = (validator,proto) -> {
         
        validator.assertValid(proto);
    };
}
