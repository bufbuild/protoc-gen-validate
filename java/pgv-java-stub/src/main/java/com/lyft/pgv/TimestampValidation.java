package com.lyft.pgv;

import com.google.protobuf.Timestamp;
import com.google.protobuf.util.Timestamps;

public final class TimestampValidation {
    private TimestampValidation() { }

    public static void constant(String field, Timestamp value, Timestamp expected) throws ValidationException {
        if (Timestamps.compare(value, expected) != 0) {
            throw new ValidationException(field, "value must equal " + expected);
        }
    }

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

    public static void lessThanNow(String field, Timestamp value) throws ValidationException {
        if (Timestamps.compare(value, currentTimestamp()) < 0) {
            throw new ValidationException(field, "value must be less than current time  " + value);
        }
    }

    public static void greaterThanNow(String field, Timestamp value) throws ValidationException {
        if (Timestamps.compare(value, currentTimestamp()) < 0) {
            throw new ValidationException(field, "value must be greater than current time  " + value);
        }
    }

    public static void withIn(String field, Timestamp value, long within) throws ValidationException {
        if (Math.abs(currentTimestamp().getSeconds() - value.getSeconds()) > within) {
            throw new ValidationException(field, "value must be within  " + within);
        }
    }

    public static Timestamp toTimestamp(long seconds, long nanos) {
      return Timestamp.newBuilder()
      .setSeconds(seconds)
      .setNanos((int)nanos)
      .build();
    }

    public static Timestamp currentTimestamp() {
      return Timestamp.newBuilder()
      .setSeconds(System.currentTimeMillis()/1000)
      .build();
    }
}
