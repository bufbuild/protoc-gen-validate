package com.lyft.pgv;

import com.google.protobuf.GeneratedMessageV3;

public class MessageValidation {
    public static void required(String field, GeneratedMessageV3 value) throws ValidationException {
        if (value == null) {
            throw new ValidationException(field, "is required");
        }
    }
}
