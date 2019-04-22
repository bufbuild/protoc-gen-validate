package io.envoyproxy.pgv;

/**
 * Base class for failed field validations.
 */
public class ValidationException extends Exception {
    public ValidationException(String field, Object value, String reason) {
        super(field + ": " + reason + " - Got " + value.toString());
    }
}
