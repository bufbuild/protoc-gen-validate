package com.lyft.pgv;

import com.google.protobuf.Duration;
import com.google.protobuf.util.Durations;

public final class DurationValidation {
    private DurationValidation() { }

    public static void constant(String field, Duration value, Duration expected) throws ValidationException {
        if (Durations.compare(value, expected) != 0) {
            throw new ValidationException(field, "value must equal " + expected);
        }
    }

    public static void lessThan(String field, Duration value, Duration lessThan) throws ValidationException {
        if (Durations.compare(value, lessThan) >= 0 ) {
            throw new ValidationException(field, "value must be less than " + lessThan);
        }
    }

    public static void lessThanOrEqual(String field, Duration value, Duration lessThan) throws ValidationException {
        if (Durations.compare(value, lessThan) > 0) {
            throw new ValidationException(field, "value must be less than or equal  " + lessThan);
        }
    }

    public static void greaterThan(String field, Duration value, Duration greaterThan) throws ValidationException {
        if (Durations.compare(value, greaterThan) <= 0) {
            throw new ValidationException(field, "value must be greater than  " + greaterThan);
        }
    }

    public static void greaterThanOrEqual(String field, Duration value, Duration greaterThan) throws ValidationException {
        if (Durations.compare(value, greaterThan) < 0) {
            throw new ValidationException(field, "value must be greater than or equal  " + greaterThan);
        }
    }

    public static Duration toDuration(long seconds, long nanos) {
      return Duration.newBuilder()
      .setSeconds(seconds)
      .setNanos((int)nanos)
      .build();
    }
}
