package com.lyft.pgv;

import java.util.Map;

public final class MapValidation {
    private MapValidation() {
    }

    public static void min(String field, Map value, int expected) throws ValidationException {
        if (Math.min(value.size(), expected) != expected) {
            throw new ValidationException(field, "value size must be at least " + expected);
        }
    }

    public static void max(String field, Map value, int expected) throws ValidationException {
        if (Math.max(value.size(), expected) != expected) {
            throw new ValidationException(field, "value size must not exceed " + expected);
        }
    }

    public static void noSparse(String field, Map value) throws ValidationException {
        for (Object key : value.keySet()) {
            if (value.get(key) == null) {
                throw new ValidationException(field, "value cannot be sparse, all pairs must be non null ");
            }
        }
    }

    // TODO: Key and Value validation
}
