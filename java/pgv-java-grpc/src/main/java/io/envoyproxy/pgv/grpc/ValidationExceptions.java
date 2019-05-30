package io.envoyproxy.pgv.grpc;

import io.envoyproxy.pgv.ValidationException;
import io.grpc.Status;

/**
 * {@code ValidationExceptions} provides utilities for converting {@link ValidationException} objects into gRPC
 * {@code Status} objects.
 */
public final class ValidationExceptions {
    private ValidationExceptions() { }

    /**
     * Convert a {@link ValidationException} into a gRPC {@code Status.INVALID_ARGUMENT} with appropriate message
     * and cause.
     * @param ex the {@code ValidationException} to convert.
     * @return a gRPC {@code Status.INVALID_ARGUMENT}
     */
    public static Status asStatus(ValidationException ex) {
        return Status.INVALID_ARGUMENT.withDescription(ex.getMessage()).withCause(ex);
    }
}
