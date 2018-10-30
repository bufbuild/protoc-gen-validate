package com.lyft.pgv;

import com.lyft.pvg.cases.Enum;
import org.junit.Test;

import static org.assertj.core.api.Assertions.assertThatThrownBy;

public class EnumValidationTest {
    @Test
    public void constantWorks() throws ValidationException {
        // Equals
        EnumValidation.constant("x", Enum.TestEnum.ONE, Enum.TestEnum.ONE);
        // Not Equals
        assertThatThrownBy(() -> EnumValidation.constant("x", Enum.TestEnum.ONE, Enum.TestEnum.TWO)).isInstanceOf(ValidationException.class);
    }

    @Test
    public void definedOnlyWorks() throws ValidationException {
        // Defined
        EnumValidation.definedOnly("x", Enum.TestEnum.ONE);
        // Not Defined
        assertThatThrownBy(() -> EnumValidation.definedOnly("x", Enum.TestEnum.UNRECOGNIZED)).isInstanceOf(ValidationException.class);
    }

    @Test
    public void inWorks() throws ValidationException {
        Enum.TestEnum[] set = new Enum.TestEnum[]{
                Enum.TestEnum.forNumber(0),
                Enum.TestEnum.forNumber(2),
        };
        // In
        EnumValidation.in("x", Enum.TestEnum.TWO, set);
        // Not In
        assertThatThrownBy(() -> EnumValidation.in("x", Enum.TestEnum.ONE, set)).isInstanceOf(ValidationException.class);
    }

    @Test
    public void notInWorks() throws ValidationException {
        Enum.TestEnum[] set = new Enum.TestEnum[]{
                Enum.TestEnum.forNumber(0),
                Enum.TestEnum.forNumber(2),
        };
        // In
        assertThatThrownBy(() -> EnumValidation.notIn("x", Enum.TestEnum.TWO, set)).isInstanceOf(ValidationException.class);
        // Not In
        EnumValidation.notIn("x", Enum.TestEnum.ONE, set);
    }
}
