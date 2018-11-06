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

    @Test
    public void betweenInclusiveWorks() throws ValidationException {
        // Lower outside
        assertThatThrownBy(() -> ComparativeValidation.between("x", 5, 10, true, 20, true, Comparator.naturalOrder())).isInstanceOf(ValidationException.class);
        // Lower bound
        ComparativeValidation.between("x", 10, 10, true, 20, true, Comparator.naturalOrder());
        // Inside
        ComparativeValidation.between("x", 15, 10, true, 20, true, Comparator.naturalOrder());
        // Upper bound
        ComparativeValidation.between("x", 20, 10, true, 20, true, Comparator.naturalOrder());
        // Upper outside
        assertThatThrownBy(() -> ComparativeValidation.between("x", 25, 10, true, 20, true, Comparator.naturalOrder())).isInstanceOf(ValidationException.class);
    }

    @Test
    public void betweenExclusiveWorks() throws ValidationException {
        // Lower outside
        assertThatThrownBy(() -> ComparativeValidation.between("x", 5, 10, false, 20, false, Comparator.naturalOrder())).isInstanceOf(ValidationException.class);
        // Lower bound
        assertThatThrownBy(() -> ComparativeValidation.between("x", 10, 10, false, 20, false, Comparator.naturalOrder())).isInstanceOf(ValidationException.class);
        // Inside
        ComparativeValidation.between("x", 15, 10, false, 20, false, Comparator.naturalOrder());
        // Upper bound
        assertThatThrownBy(() -> ComparativeValidation.between("x", 20, 10, false, 20, false, Comparator.naturalOrder())).isInstanceOf(ValidationException.class);
        // Upper outside
        assertThatThrownBy(() -> ComparativeValidation.between("x", 25, 10, false, 20, false, Comparator.naturalOrder())).isInstanceOf(ValidationException.class);
    }

    @Test
    public void outsideInclusiveWorks() throws ValidationException {
        // Lower outside
        ComparativeValidation.outside("x", 5, 10, true, 20, true, Comparator.naturalOrder());
        // Lower bound
        assertThatThrownBy(() -> ComparativeValidation.outside("x", 10, 10, true, 20, true, Comparator.naturalOrder())).isInstanceOf(ValidationException.class);
        // Inside
        assertThatThrownBy(() -> ComparativeValidation.outside("x", 15, 10, true, 20, true, Comparator.naturalOrder())).isInstanceOf(ValidationException.class);
        // Upper bound
        assertThatThrownBy(() -> ComparativeValidation.outside("x", 20, 10, true, 20, true, Comparator.naturalOrder())).isInstanceOf(ValidationException.class);
        // Upper outside
        ComparativeValidation.outside("x", 25, 10, true, 20, true, Comparator.naturalOrder());
    }

    @Test
    public void outsideExclusiveWorks() throws ValidationException {
        // Lower outside
        ComparativeValidation.outside("x", 5, 10, false, 20, false, Comparator.naturalOrder());
        // Lower bound
        ComparativeValidation.outside("x", 10, 10, false, 20, false, Comparator.naturalOrder());
        // Inside
        assertThatThrownBy(() -> ComparativeValidation.outside("x", 15, 10, false, 20, false, Comparator.naturalOrder())).isInstanceOf(ValidationException.class);
        // Upper bound
        ComparativeValidation.outside("x", 20, 10, false, 20, false, Comparator.naturalOrder());
        // Upper outside
        ComparativeValidation.outside("x", 25, 10, false, 20, false, Comparator.naturalOrder());
    }
}
