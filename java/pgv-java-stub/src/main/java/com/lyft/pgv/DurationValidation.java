package com.lyft.pgv;

import com.google.protobuf.Duration;
import com.google.protobuf.util.Durations;

import java.util.Arrays;

public final class DurationValidation {
    private DurationValidation() {
    }

    public static void in(String field, Duration value, Duration[] set) throws ValidationException {
        for (Duration str : set) {
            if (value.equals(str)) {
                return;
            }
        }

        throw new ValidationException(field, "value must be in " + Arrays.toString(set));
    }

    public static void notIn(String field, Duration value, Duration[] set) throws ValidationException {
        for (Duration str : set) {
            if (value.equals(str)) {
                throw new ValidationException(field, "value must not be in " + Arrays.toString(set));
            }
        }
    }

    public static Duration toDuration(long seconds, long nanos) {
        return Duration.newBuilder()
                .setSeconds(seconds)
                .setNanos((int) nanos)
                .build();
    }
}
