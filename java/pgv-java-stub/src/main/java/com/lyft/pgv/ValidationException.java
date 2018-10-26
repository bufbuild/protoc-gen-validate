package com.lyft.pgv;

/**
 * Base class for failed field validations.
 */
public class ValidationException extends Exception {
    public ValidationException(String field, String reason) {
        super(field + ": " + reason);
    }
}
