package pgs

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"errors"

	"github.com/stretchr/testify/assert"
)

type mockLogger struct {
	buf *bytes.Buffer
}

func newMockLogger() *mockLogger               { return &mockLogger{&bytes.Buffer{}} }
func (l *mockLogger) Println(v ...interface{}) { fmt.Fprintln(l.buf, v...) }
func (l *mockLogger) Printf(format string, v ...interface{}) {
	fmt.Fprintln(l.buf, fmt.Sprintf(format, v...))
}

func TestRootDebugger_Log(t *testing.T) {
	t.Parallel()

	l := newMockLogger()
	rd := rootDebugger{l: l}
	rd.Log("foo", "bar")
	assert.Equal(t, "foo bar\n", l.buf.String())
}

func TestRootDebugger_Logf(t *testing.T) {
	t.Parallel()

	l := newMockLogger()
	rd := rootDebugger{l: l}
	rd.Logf("foo%s", "bar")
	assert.Equal(t, "foobar\n", l.buf.String())
}

func TestRootDebugger_Fail(t *testing.T) {
	t.Parallel()

	var failed bool

	fail := func(msgs ...interface{}) {
		assert.Equal(t, "foobar", msgs[0])
		failed = true
	}

	rd := rootDebugger{l: newMockLogger(), fail: fail}
	rd.Fail("foo", "bar")

	assert.True(t, failed)
}

func TestRootDebugger_Failf(t *testing.T) {
	t.Parallel()

	var failed bool

	fail := func(msgs ...interface{}) {
		assert.Equal(t, "fizz buzz", msgs[0])
		failed = true
	}

	rd := rootDebugger{l: newMockLogger(), fail: fail}
	rd.Failf("fizz %s", "buzz")

	assert.True(t, failed)
}

func TestRootDebugger_Debug(t *testing.T) {
	t.Parallel()

	l := newMockLogger()
	rd := rootDebugger{l: l}

	rd.Debug("foo")
	assert.Empty(t, l.buf.String())

	rd.logDebugs = true

	rd.Debug("bar")
	assert.Contains(t, l.buf.String(), "bar")
}

func TestRootDebugger_Debugf(t *testing.T) {
	t.Parallel()

	l := newMockLogger()
	rd := rootDebugger{l: l}

	rd.Debug("foo")
	assert.Empty(t, l.buf.String())

	rd.logDebugs = true

	rd.Debug("bar")
	assert.Contains(t, l.buf.String(), "bar")
}

func TestRootDebugger_CheckErr(t *testing.T) {
	t.Parallel()

	e := errors.New("bad error")
	errd := false

	errfn := func(err error, msg ...interface{}) {
		assert.Equal(t, e, err)
		assert.Equal(t, "foo", msg[0])
		errd = true
	}

	rd := rootDebugger{err: errfn}
	rd.CheckErr(nil, "fizz")
	assert.False(t, errd)

	rd.CheckErr(e, "foo")
	assert.True(t, errd)
}

func TestRootDebugger_Assert(t *testing.T) {
	t.Parallel()

	failed := false

	fail := func(msgs ...interface{}) {
		assert.Equal(t, "foo", msgs[0])
		failed = true
	}

	rd := rootDebugger{fail: fail}
	rd.Assert(true, "fizz")
	assert.False(t, failed)

	rd.Assert(false, "foo")
	assert.True(t, failed)
}

func TestRootDebugger_Exit(t *testing.T) {
	t.Parallel()

	var code int

	rd := rootDebugger{exit: func(c int) { code = c }}
	rd.Exit(123)
	assert.Equal(t, 123, code)
}

func TestRootDebugger_Push(t *testing.T) {
	t.Parallel()

	rd := rootDebugger{}

	d := rd.Push("foo")
	assert.NotNil(t, d)
	assert.NotEqual(t, rd, d)
}

func TestRootDebugger_Pop(t *testing.T) {
	t.Parallel()

	rd := rootDebugger{}
	assert.Panics(t, func() { rd.Pop() })
}

func TestRootDebugger_DefaultErr(t *testing.T) {
	t.Parallel()

	exited := false
	code := 0
	l := newMockLogger()
	rd := rootDebugger{
		l: l,
		exit: func(c int) {
			code = c
			exited = true
		},
	}

	rd.defaultErr(nil, "nothing")

	assert.False(t, exited)
	assert.Empty(t, l.buf.String())

	rd.defaultErr(errors.New("some error"), "something")
	assert.True(t, exited)
	assert.Equal(t, 1, code)
	assert.Contains(t, l.buf.String(), "something")
}

func TestRootDebugger_DefaultFail(t *testing.T) {
	t.Parallel()

	exited := false
	code := 0
	l := newMockLogger()
	rd := rootDebugger{
		l: l,
		exit: func(c int) {
			code = c
			exited = true
		},
	}

	rd.defaultFail("something")
	assert.True(t, exited)
	assert.Equal(t, 1, code)
	assert.Contains(t, l.buf.String(), "something")
}

func TestPrefixedDebugger_Log(t *testing.T) {
	t.Parallel()

	l := newMockLogger()
	d := rootDebugger{l: l}.Push("FIZZ")
	d.Log("foo", "bar")
	assert.Contains(t, l.buf.String(), "FIZZ")
	assert.Contains(t, l.buf.String(), "foo bar")
}

func TestPrefixedDebugger_Logf(t *testing.T) {
	t.Parallel()

	l := newMockLogger()
	d := rootDebugger{l: l}.Push("FIZZ")
	d.Logf("foo%s", "bar")
	assert.Contains(t, l.buf.String(), "FIZZ")
	assert.Contains(t, l.buf.String(), "foobar")
}

func TestPrefixedDebugger_Fail(t *testing.T) {
	t.Parallel()

	var failed bool

	fail := func(msgs ...interface{}) {
		assert.Contains(t, msgs[0], "FIZZ")
		assert.Contains(t, msgs[0], "foobar")
		failed = true
	}

	d := rootDebugger{l: newMockLogger(), fail: fail}.Push("FIZZ")
	d.Fail("foo", "bar")

	assert.True(t, failed)
}

func TestPrefixedDebugger_Failf(t *testing.T) {
	t.Parallel()

	var failed bool

	fail := func(msgs ...interface{}) {
		assert.Contains(t, msgs[0], "FIZZ")
		assert.Contains(t, msgs[0], "foo bar")
		failed = true
	}

	d := rootDebugger{l: newMockLogger(), fail: fail}.Push("FIZZ")
	d.Failf("foo %s", "bar")

	assert.True(t, failed)
}

func TestPrefixedDebugger_Debug(t *testing.T) {
	t.Parallel()

	l := newMockLogger()
	rd := rootDebugger{l: l}
	d := rd.Push("FIZZ")

	d.Debug("foo")
	assert.Empty(t, l.buf.String())

	rd.logDebugs = true
	d = rd.Push("FIZZ")

	d.Debug("bar")
	assert.Contains(t, l.buf.String(), "bar")
	assert.Contains(t, l.buf.String(), "FIZZ")
}

func TestPrefixedDebugger_Debugf(t *testing.T) {
	t.Parallel()

	l := newMockLogger()
	rd := rootDebugger{l: l}
	d := rd.Push("FIZZ")

	d.Debugf("foo%s", "bar")
	assert.Empty(t, l.buf.String())

	rd.logDebugs = true
	d = rd.Push("FIZZ")

	d.Debugf("foo%s", "bar")
	assert.Contains(t, l.buf.String(), "foobar")
	assert.Contains(t, l.buf.String(), "FIZZ")
}

func TestPrefixedDebugger_CheckErr(t *testing.T) {
	t.Parallel()

	e := errors.New("bad error")
	errd := false

	errfn := func(err error, msg ...interface{}) {
		assert.Equal(t, e, err)
		assert.Contains(t, msg[0], "foo")
		assert.Contains(t, msg[0], "FIZZ")
		errd = true
	}

	d := rootDebugger{err: errfn}.Push("FIZZ")
	d.CheckErr(nil, "fizz")
	assert.False(t, errd)

	d.CheckErr(e, "foo")
	assert.True(t, errd)
}

func TestPrefixedDebugger_Assert(t *testing.T) {
	t.Parallel()

	failed := false

	fail := func(msgs ...interface{}) {
		assert.Contains(t, msgs[0], "FIZZ")
		assert.Contains(t, msgs[0], "foo")
		failed = true
	}

	d := rootDebugger{fail: fail}.Push("FIZZ")
	d.Assert(1 == 1, "fizz")
	assert.False(t, failed)

	d.Assert(1 == 0, "foo")
	assert.True(t, failed)
}

func TestPrefixedDebugger_Pop(t *testing.T) {
	t.Parallel()

	rd := rootDebugger{}
	d := rd.Push("FOO")
	assert.Equal(t, rd, d.Pop())
}

func TestPrefixedDebugger_Push(t *testing.T) {
	t.Parallel()

	l := newMockLogger()
	rd := rootDebugger{l: l}
	d := rd.Push("FOO").Push("BAR")
	d.Log("fizz")
	assert.Contains(t, l.buf.String(), "FOO")
	assert.Contains(t, l.buf.String(), "BAR")
}

func TestPrefixedDebugger_Push_Format(t *testing.T) {
	t.Parallel()

	l := newMockLogger()
	d := rootDebugger{l: l}.Push("foo").Push("bar")
	d.Logf("%s", "baz")

	assert.Equal(t, "[foo][bar] baz\n", l.buf.String())
}

func TestPrefixedDebugger_Exit(t *testing.T) {
	t.Parallel()

	md := InitMockDebugger()
	d := &prefixedDebugger{parent: md}
	d.Exit(123)

	assert.True(t, md.Exited())
	assert.Equal(t, 123, md.ExitCode())
}

func TestInitDebugger(t *testing.T) {
	t.Parallel()
	d := initDebugger(true, nil)
	assert.NotNil(t, d)
}

func TestMockDebugger_Output(t *testing.T) {
	t.Parallel()

	md := InitMockDebugger()
	md.Log("foobar")
	b, _ := ioutil.ReadAll(md.Output())
	assert.Equal(t, "foobar\n", string(b))
}
