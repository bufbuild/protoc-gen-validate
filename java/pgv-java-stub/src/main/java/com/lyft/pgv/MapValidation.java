package com.lyft.pgv;

import java.util.Iterator;
import java.util.Map;

public final class MapValidation {
    private MapValidation() {
    }

    public static void min(String field, Map value, int expected) throws ValidationException {
        if (Math.min(value.size(), expected) != expected) {
            throw new ValidationException(field, "value size must be atleast " + expected );
        }
    }

    public static void max(String field, Map value, int expected) throws ValidationException {
        if (Math.max(value.size(), expected) != expected) {
            throw new ValidationException(field, "value size must not exceed " + expected );
        }
    }

    public static void noSparse(String field, Map value) throws ValidationException {
      Iterator iterator = value.keySet().iterator();
       while (iterator.hasNext()) {
         if (value.get(iterator.next()) == null) {
           throw new ValidationException(field, "value cannot be sparse, all pairs must be non null ");
         }
       }
    }
}
