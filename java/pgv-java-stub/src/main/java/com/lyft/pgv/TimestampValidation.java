package com.lyft.pgv;

import com.google.protobuf.Duration;
import com.google.protobuf.Timestamp;
import com.google.protobuf.util.Durations;
import com.google.protobuf.util.Timestamps;

public final class TimestampValidation {
    private TimestampValidation() { }

    public static void lessThan(String field, Timestamp value, Timestamp lessThan) throws ValidationException {
        if (Timestamps.compare(value, lessThan) >= 0 ) {
            throw new ValidationException(field, "value must be less than " + lessThan);
        }
    }

    public static void lessThanOrEqual(String field, Timestamp value, Timestamp lessThan) throws ValidationException {
        if (Timestamps.compare(value, lessThan) > 0) {
            throw new ValidationException(field, "value must be less than or equal  " + lessThan);
        }
    }

    public static void greaterThan(String field, Timestamp value, Timestamp greaterThan) throws ValidationException {
        if (Timestamps.compare(value, greaterThan) <= 0) {
            throw new ValidationException(field, "value must be greater than  " + greaterThan);
        }
    }

    public static void greaterThanOrEqual(String field, Timestamp value, Timestamp greaterThan) throws ValidationException {
        if (Timestamps.compare(value, greaterThan) < 0) {
            throw new ValidationException(field, "value must be greater than or equal  " + greaterThan);
        }
    }

    public static void within(String field, Timestamp value, Duration duration, Timestamp when) throws ValidationException {
        Duration between = Timestamps.between(when, value);
        if (Long.compare(Math.abs(Durations.toNanos(between)), Math.abs(Durations.toNanos(duration))) == 1) {
            throw new ValidationException(field, "value must be within " + Durations.toString(duration) + " of " + Timestamps.toString(when));
        }
    }

    public static Timestamp toTimestamp(long seconds, int nanos) {
        return Timestamp.newBuilder()
                .setSeconds(seconds)
                .setNanos(nanos)
                .build();
    }

    public static Timestamp currentTimestamp() {
        return Timestamps.fromMillis(System.currentTimeMillis());
    }
}
