package com.lyft.pgv;

import org.junit.Test;

import java.util.Comparator;

import static org.assertj.core.api.Assertions.assertThatThrownBy;

public class ComparativeValidationTest {
    @Test
    public void lessThanWorks() throws ValidationException {
        // Less than
        ComparativeValidation.lessThan("x", 10, 20, Comparator.naturalOrder());
        // Equal to
        assertThatThrownBy(() -> ComparativeValidation.lessThan("x", 10, 10, Comparator.naturalOrder()));
        // Greater than
        assertThatThrownBy(() -> ComparativeValidation.lessThan("x", 20, 10, Comparator.naturalOrder()));
    }

    @Test
    public void lessThanOrEqualWorks() throws ValidationException {
        // Less than
        ComparativeValidation.lessThanOrEqual("x", 10, 20, Comparator.naturalOrder());
        // Equal to
        ComparativeValidation.lessThanOrEqual("x", 10, 10, Comparator.naturalOrder());
        // Greater than
        assertThatThrownBy(() -> ComparativeValidation.lessThanOrEqual("x", 20, 10, Comparator.naturalOrder()));
    }

    @Test
    public void greaterThanWorks() throws ValidationException {
        // Less than
        assertThatThrownBy(() -> ComparativeValidation.greaterThan("x", 10, 20, Comparator.naturalOrder()));
        // Equal to
        assertThatThrownBy(() -> ComparativeValidation.greaterThan("x", 10, 10, Comparator.naturalOrder()));
        // Greater than
        ComparativeValidation.greaterThan("x", 20, 10, Comparator.naturalOrder());
    }

    @Test
    public void greaterThanOrEqualWorks() throws ValidationException {
        // Less than
        assertThatThrownBy(() -> ComparativeValidation.greaterThanOrEqual("x", 10, 20, Comparator.naturalOrder()));
        // Equal to
        ComparativeValidation.greaterThanOrEqual("x", 10, 10, Comparator.naturalOrder());
        // Greater than
        ComparativeValidation.greaterThanOrEqual("x", 20, 10, Comparator.naturalOrder());
    }
}
