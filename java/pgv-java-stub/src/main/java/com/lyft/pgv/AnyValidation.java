package com.lyft.pgv;

import com.google.protobuf.Any;

import java.util.Arrays;

public final class AnyValidation {
    private AnyValidation() {
    }

    public static void in(String field, Any value, String[] set) throws ValidationException {
      for (String typeUrl : set) {
          if (value.getTypeUrl().equals(typeUrl)) {
              return;
          }
      }

      throw new ValidationException(field, "value must be of type " + Arrays.toString(set));
    }

    public static void notIn(String field, Any value, String[] set) throws ValidationException {
        for (String typeUrl : set) {
            if (value.getTypeUrl().equals(typeUrl)) {
                throw new ValidationException(field, "value must not be of type " + Arrays.toString(set));
            }
        }
    }
}
