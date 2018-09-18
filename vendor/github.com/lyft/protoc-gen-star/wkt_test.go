package pgs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLookupWKT(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     Name
		expected WellKnownType
	}{
		{"Any", AnyWKT},
		{"Duration", DurationWKT},
		{"Empty", EmptyWKT},
		{"Foobar", UnknownWKT},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name.String(), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expected, LookupWKT(tc.name))
		})
	}
}

func TestWellKnownType_Name(t *testing.T) {
	t.Parallel()

	wkt := WellKnownType("Foobar")
	assert.Equal(t, Name("Foobar"), wkt.Name())
}

func TestWellKnownType_Valid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		wkt      WellKnownType
		expected bool
	}{
		{AnyWKT, true},
		{Int64ValueWKT, true},
		{UnknownWKT, false},
		{WellKnownType("Foobar"), false},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.wkt.Name().String(), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expected, tc.wkt.Valid())
		})
	}
}
