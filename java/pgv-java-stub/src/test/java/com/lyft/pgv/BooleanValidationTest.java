package com.lyft.pgv;

import org.junit.Test;

import static org.assertj.core.api.Assertions.assertThatThrownBy;

public class BooleanValidationTest {
    @Test
    public void constantWorks() throws ValidationException {
        BooleanValidation.constant("x", true, true);
        assertThatThrownBy(() -> BooleanValidation.constant("x", true, false)).isInstanceOf(ValidationException.class);
    }
}
