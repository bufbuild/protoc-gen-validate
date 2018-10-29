package com.lyft.pgv;

import com.google.protobuf.Duration;
import com.google.protobuf.util.Durations;
import org.junit.Test;

import static org.assertj.core.api.Assertions.assertThatThrownBy;

public class DurationValidationTest {
    @Test
    public void constantWorks() throws ValidationException {
        // Equal
        DurationValidation.constant("x", Durations.fromSeconds(10), Durations.fromSeconds(10));
        // Not Equal
        assertThatThrownBy(() -> DurationValidation.constant("x", Durations.fromSeconds(10), Durations.fromSeconds(20))).isInstanceOf(ValidationException.class);
    }

    @Test
    public void lessThanWorks() throws ValidationException {
        // Less
        DurationValidation.lessThan("x", Durations.fromSeconds(10), Durations.fromSeconds(20));
        // Equal
        assertThatThrownBy(() -> DurationValidation.lessThan("x", Durations.fromSeconds(10), Durations.fromSeconds(10))).isInstanceOf(ValidationException.class);
        // Greater
        assertThatThrownBy(() -> DurationValidation.lessThan("x", Durations.fromSeconds(20), Durations.fromSeconds(10))).isInstanceOf(ValidationException.class);
    }

    @Test
    public void lessThanOrEqualsWorks() throws ValidationException {
        // Less
        DurationValidation.lessThanOrEqual("x", Durations.fromSeconds(10), Durations.fromSeconds(20));
        // Equal
        DurationValidation.lessThanOrEqual("x", Durations.fromSeconds(10), Durations.fromSeconds(10));
        // Greater
        assertThatThrownBy(() -> DurationValidation.lessThanOrEqual("x", Durations.fromSeconds(20), Durations.fromSeconds(10))).isInstanceOf(ValidationException.class);
    }

    @Test
    public void greaterThanWorks() throws ValidationException {
        // Less
        assertThatThrownBy(() -> DurationValidation.greaterThan("x", Durations.fromSeconds(10), Durations.fromSeconds(20))).isInstanceOf(ValidationException.class);
        // Equal
        assertThatThrownBy(() -> DurationValidation.greaterThan("x", Durations.fromSeconds(10), Durations.fromSeconds(10))).isInstanceOf(ValidationException.class);
        // Greater
        DurationValidation.greaterThan("x", Durations.fromSeconds(20), Durations.fromSeconds(10));
    }

    @Test
    public void greaterThanOrEqualsWorks() throws ValidationException {
        // Less
        assertThatThrownBy(() -> DurationValidation.greaterThanOrEqual("x", Durations.fromSeconds(10), Durations.fromSeconds(20))).isInstanceOf(ValidationException.class);
        // Equal
        DurationValidation.greaterThanOrEqual("x", Durations.fromSeconds(10), Durations.fromSeconds(10));
        // Greater
        DurationValidation.greaterThanOrEqual("x", Durations.fromSeconds(20), Durations.fromSeconds(10));
    }

    @Test
    public void inWorks() throws ValidationException {
        Duration[] set = new Duration[]{DurationValidation.toDuration(1, 0), DurationValidation.toDuration(2, 0)};
        // In
        DurationValidation.in("x", DurationValidation.toDuration(1, 0), set);
        // Not In
        assertThatThrownBy(() -> DurationValidation.in("x", DurationValidation.toDuration(3, 0), set)).isInstanceOf(ValidationException.class);
    }

    @Test
    public void notInWorks() throws ValidationException {
        Duration[] set = new Duration[]{DurationValidation.toDuration(1, 0), DurationValidation.toDuration(2, 0)};
        // In
        assertThatThrownBy(() -> DurationValidation.notIn("x", DurationValidation.toDuration(1, 0), set)).isInstanceOf(ValidationException.class);
        // Not In
        DurationValidation.notIn("x", DurationValidation.toDuration(3, 0), set);
    }
}
