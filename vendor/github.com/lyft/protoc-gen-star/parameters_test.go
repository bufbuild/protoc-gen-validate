package pgs

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParameters_OutputPath(t *testing.T) {
	t.Parallel()

	p := Parameters{}
	assert.Equal(t, ".", p.OutputPath())
	p.SetOutputPath("foo")
	assert.Equal(t, "foo", p.OutputPath())
}

func TestParseParameters(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in  string
		out Parameters
	}{
		{
			"foo=bar",
			Parameters{"foo": "bar"},
		},
		{
			"fizz",
			Parameters{"fizz": ""},
		},
		{
			"foo=bar,fizz=buzz",
			Parameters{"foo": "bar", "fizz": "buzz"},
		},
		{
			"foo=bar,foo",
			Parameters{"foo": ""},
		},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.in, func(t *testing.T) {
			assert.Equal(t, tc.out, ParseParameters(tc.in))
		})
	}
}

func TestParameters_String(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in  Parameters
		out string
	}{
		{
			Parameters{"foo": "bar"},
			"foo=bar",
		},
		{
			Parameters{"fizz": ""},
			"fizz",
		},
		{
			Parameters{"foo": "bar", "fizz": ""},
			"fizz,foo=bar",
		},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.out, func(t *testing.T) {
			assert.Equal(t, tc.out, tc.in.String())
		})
	}
}

func TestParameters_Str(t *testing.T) {
	t.Parallel()

	p := Parameters{"foo": "bar"}

	assert.Equal(t, "bar", p.Str("foo"))
	assert.Empty(t, p.Str("baz"))
	assert.Equal(t, "fizz", p.StrDefault("baz", "fizz"))

	p.SetStr("baz", "buzz")
	assert.Equal(t, "buzz", p.Str("baz"))
}

func TestParameters_Int(t *testing.T) {
	t.Parallel()

	p := Parameters{"foo": "456", "fizz": "buzz"}

	out, err := p.Int("foo")
	assert.NoError(t, err)
	assert.Equal(t, 456, out)

	_, err = p.Int("fizz")
	assert.Error(t, err)

	out, err = p.Int("baz")
	assert.NoError(t, err)
	assert.Zero(t, out)

	out, err = p.IntDefault("baz", 123)
	assert.NoError(t, err)
	assert.Equal(t, 123, out)

	p.SetInt("baz", 789)
	out, err = p.Int("baz")
	assert.NoError(t, err)
	assert.Equal(t, 789, out)
}

func TestParameters_Uint(t *testing.T) {
	t.Parallel()

	p := Parameters{"foo": "456", "fizz": "-789"}

	out, err := p.Uint("foo")
	assert.NoError(t, err)
	assert.Equal(t, uint(456), out)

	_, err = p.Uint("fizz")
	assert.Error(t, err)

	out, err = p.Uint("buzz")
	assert.NoError(t, err)
	assert.Zero(t, out)

	out, err = p.UintDefault("baz", 123)
	assert.NoError(t, err)
	assert.Equal(t, uint(123), out)

	p.SetUint("baz", 999)
	out, err = p.Uint("baz")
	assert.NoError(t, err)
	assert.Equal(t, uint(999), out)
}

func TestParameters_Float(t *testing.T) {
	t.Parallel()

	p := Parameters{"foo": "1.23", "fizz": "buzz"}

	out, err := p.Float("foo")
	assert.NoError(t, err)
	assert.Equal(t, 1.23, out)

	_, err = p.Float("fizz")
	assert.Error(t, err)

	out, err = p.Float("baz")
	assert.NoError(t, err)
	assert.Zero(t, out)

	out, err = p.FloatDefault("baz", 4.56)
	assert.NoError(t, err)
	assert.Equal(t, 4.56, out)

	p.SetFloat("baz", -7.89)
	out, err = p.Float("baz")
	assert.NoError(t, err)
	assert.Equal(t, -7.89, out)
}

func TestParameters_Bool(t *testing.T) {
	t.Parallel()

	p := Parameters{"foo": "true", "bar": "", "fizz": "buzz"}

	out, err := p.Bool("foo")
	assert.NoError(t, err)
	assert.True(t, out)

	out, err = p.Bool("bar")
	assert.NoError(t, err)
	assert.True(t, out)

	_, err = p.Bool("fizz")
	assert.Error(t, err)

	out, err = p.Bool("baz")
	assert.NoError(t, err)
	assert.False(t, out)

	out, err = p.BoolDefault("baz", true)
	assert.NoError(t, err)
	assert.True(t, out)

	p.SetBool("baz", true)
	out, err = p.Bool("baz")
	assert.NoError(t, err)
	assert.True(t, out)
}

func TestParameters_Duration(t *testing.T) {
	t.Parallel()

	p := Parameters{"foo": "123s", "fizz": "buzz"}

	out, err := p.Duration("foo")
	assert.NoError(t, err)
	assert.Equal(t, 123*time.Second, out)

	_, err = p.Duration("fizz")
	assert.Error(t, err)

	out, err = p.Duration("baz")
	assert.NoError(t, err)
	assert.Zero(t, out)

	out, err = p.DurationDefault("baz", 456*time.Second)
	assert.NoError(t, err)
	assert.Equal(t, 456*time.Second, out)

	p.SetDuration("baz", 789*time.Second)
	out, err = p.Duration("baz")
	assert.NoError(t, err)
	assert.Equal(t, 789*time.Second, out)
}

func TestParameters_Clone(t *testing.T) {
	t.Parallel()

	orig := Parameters{"foo": "bar", "fizz": "buzz"}

	clone := orig.Clone()
	assert.Equal(t, orig, clone)

	clone.SetStr("foo", "baz")
	assert.NotEqual(t, orig, clone)
}
