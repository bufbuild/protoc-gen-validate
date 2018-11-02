package com.lyft.pgv;

import org.junit.Test;

import static org.assertj.core.api.Assertions.assertThatThrownBy;

public class ConstantValidationTest {
    @Test
    public void constantWorks() throws ValidationException {
        ConstantValidation.constant("x", true, true);
        assertThatThrownBy(() -> ConstantValidation.constant("x", true, false)).isInstanceOf(ValidationException.class);
    }
}
