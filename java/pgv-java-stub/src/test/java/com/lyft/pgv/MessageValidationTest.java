package com.lyft.pgv;

import com.lyft.pvg.cases.Enum;
import org.junit.Test;

import static org.assertj.core.api.Assertions.assertThatThrownBy;

public class MessageValidationTest {
    @Test
    public void requiredWorks() throws ValidationException {
        // Present
        MessageValidation.required("x", Enum.Outer.getDefaultInstance());
        // Absent
        assertThatThrownBy(() -> MessageValidation.required("x", null)).isInstanceOf(ValidationException.class);
    }
}
