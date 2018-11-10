package com.lyft.pgv;

public class UnimplementedException extends ValidationException {
    public UnimplementedException(String field, String reason) {
        super(field, reason);
    }
}
