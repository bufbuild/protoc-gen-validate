package com.lyft.pgv;

import com.google.protobuf.Any;
import org.junit.Test;

import static org.assertj.core.api.Assertions.assertThatThrownBy;

public class AnyValidationTest {

    @Test
    public void inWorks() throws ValidationException {
        String[] set = new String[]{"type.googleapis.com/google.protobuf.Duration"};

        // In
        AnyValidation.in("x", Any.newBuilder().setTypeUrl("type.googleapis.com/google.protobuf.Duration").build(), set);

        // Not In
        assertThatThrownBy(() -> AnyValidation.in("x", Any.newBuilder().setTypeUrl("junk").build(), set)).isInstanceOf(ValidationException.class);
    }

    @Test
    public void notInWorks() throws ValidationException {
        String[] set = new String[]{"type.googleapis.com/google.protobuf.Duration"};

        // In
        assertThatThrownBy(() -> AnyValidation.notIn("x", Any.newBuilder().setTypeUrl("type.googleapis.com/google.protobuf.Duration").build(), set)).isInstanceOf(ValidationException.class);

        // Not In
        AnyValidation.notIn("x", Any.newBuilder().setTypeUrl("junk").build(), set);
    }
}
