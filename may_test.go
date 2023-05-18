package fo

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var _ Logger = (*mayTestLogger)(nil)

type mayTestLogger struct {
	sb *strings.Builder
}

func (l *mayTestLogger) Error(v ...any) {
	inputs := fmt.Sprintln(v...)
	l.sb.WriteString(inputs)
	log.Printf("[test logger] %s", inputs)
}

func (l *mayTestLogger) Flush() {
	l.sb.Reset()
}

func newMayTestLogger() *mayTestLogger {
	return &mayTestLogger{
		sb: &strings.Builder{},
	}
}

func TestSetLogger(t *testing.T) {
	SetLoggers(newMayTestLogger())

	assert.Len(t, internalMayHandlers.handlers, 1)
}

func TestSetHandlers(t *testing.T) {
	SetHandlers(WithLoggerHandler(newMayTestLogger()))

	assert.Len(t, internalMayHandlers.handlers, 1)
}

func TestMay(t *testing.T) {
	logger := newMayTestLogger()
	internalMayHandlers.Use(WithLoggerHandler(logger))

	defer func() {
		internalMayHandlers.handlers = make([]MayHandler, 0)
	}()

	assert.New(t)

	assert.Equal(t, "foo", May("foo", nil))
	assert.Empty(t, logger.sb.String())
	logger.Flush()

	May("", errors.New("something went wrong"))
	assert.Equal(t, "something went wrong\n", logger.sb.String())
	logger.Flush()

	May("", errors.New("something went wrong"), "operation shouldn't fail")
	assert.Equal(t, "operation shouldn't fail: something went wrong\n", logger.sb.String())
	logger.Flush()

	May("", errors.New("something went wrong"), "operation shouldn't fail with %s", "foo")
	assert.Equal(t, "operation shouldn't fail with foo: something went wrong\n", logger.sb.String())
	logger.Flush()

	assert.Equal(t, 1, May(1, true))
	assert.Empty(t, logger.sb.String())
	logger.Flush()

	May(1, false)
	assert.Equal(t, "not ok\n", logger.sb.String())
	logger.Flush()

	May(1, false, "operation shouldn't fail")
	assert.Equal(t, "operation shouldn't fail\n", logger.sb.String())
	logger.Flush()

	May(1, false, "operation shouldn't fail with %s", "foo")
	assert.Equal(t, "operation shouldn't fail with foo\n", logger.sb.String())
	logger.Flush()

	cb := func() error {
		return assert.AnError
	}

	May0(cb(), "operation should fail")
	assert.Equal(t, "operation should fail: assert.AnError general error for testing\n", logger.sb.String())
	logger.Flush()

	assert.PanicsWithValue(t, "may: invalid err type 'int', should either be a bool or an error", func() {
		May0(0)
	})
	assert.PanicsWithValue(t, "may: invalid err type 'string', should either be a bool or an error", func() {
		May0("error")
	})
}

func ExampleMay() {
	funcWithError := func() (int, error) {
		return 0, errors.New("something went wrong")
	}

	funcWithNilErr := func() (int, error) {
		return 42, nil
	}

	str := May(funcWithError())
	str2 := May(funcWithNilErr())

	fmt.Println(str)
	fmt.Println(str2)
	// Output: 0
	// 42
}

func TestMayX(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		logger := newMayTestLogger()
		internalMayHandlers.Use(WithLoggerHandler(logger))

		defer func() {
			internalMayHandlers.handlers = make([]MayHandler, 0)
		}()

		{
			May0(errors.New("something went wrong"))
			assert.Equal(t, "something went wrong\n", logger.sb.String())
			logger.Flush()

			May0(errors.New("something went wrong"), "operation shouldn't fail with %s", "foo")
			assert.Equal(t, "operation shouldn't fail with foo: something went wrong\n", logger.sb.String())
			logger.Flush()

			May0(nil)
			assert.Empty(t, logger.sb.String())
		}

		{
			val1 := May1(1, nil)
			assert.Equal(t, 1, val1)

			May1(1, errors.New("something went wrong"))
			assert.Equal(t, "something went wrong\n", logger.sb.String())
			logger.Flush()

			May1(1, errors.New("something went wrong"), "operation shouldn't fail with %s", "foo")
			assert.Equal(t, "operation shouldn't fail with foo: something went wrong\n", logger.sb.String())
			logger.Flush()
		}

		{
			val1, val2 := May2(1, 2, nil)
			assert.Equal(t, 1, val1)
			assert.Equal(t, 2, val2)

			May2(1, 2, errors.New("something went wrong"))
			assert.Equal(t, "something went wrong\n", logger.sb.String())
			logger.Flush()

			May2(1, 2, errors.New("something went wrong"), "operation shouldn't fail with %s", "foo")
			assert.Equal(t, "operation shouldn't fail with foo: something went wrong\n", logger.sb.String())
			logger.Flush()
		}

		{
			val1, val2, val3 := May3(1, 2, 3, nil)
			assert.Equal(t, 1, val1)
			assert.Equal(t, 2, val2)
			assert.Equal(t, 3, val3)

			May3(1, 2, 3, errors.New("something went wrong"))
			assert.Equal(t, "something went wrong\n", logger.sb.String())
			logger.Flush()

			May3(1, 2, 3, errors.New("something went wrong"), "operation shouldn't fail with %s", "foo")
			assert.Equal(t, "operation shouldn't fail with foo: something went wrong\n", logger.sb.String())
			logger.Flush()
		}

		{
			val1, val2, val3, val4 := May4(1, 2, 3, 4, nil)
			assert.Equal(t, 1, val1)
			assert.Equal(t, 2, val2)
			assert.Equal(t, 3, val3)
			assert.Equal(t, 4, val4)

			May4(1, 2, 3, 4, errors.New("something went wrong"))
			assert.Equal(t, "something went wrong\n", logger.sb.String())
			logger.Flush()

			May4(1, 2, 3, 4, errors.New("something went wrong"), "operation shouldn't fail with %s", "foo")
			assert.Equal(t, "operation shouldn't fail with foo: something went wrong\n", logger.sb.String())
			logger.Flush()
		}

		{
			val1, val2, val3, val4, val5 := May5(1, 2, 3, 4, 5, nil)
			assert.Equal(t, 1, val1)
			assert.Equal(t, 2, val2)
			assert.Equal(t, 3, val3)
			assert.Equal(t, 4, val4)
			assert.Equal(t, 5, val5)

			May5(1, 2, 3, 4, 5, errors.New("something went wrong"))
			assert.Equal(t, "something went wrong\n", logger.sb.String())
			logger.Flush()

			May5(1, 2, 3, 4, 5, errors.New("something went wrong"), "operation shouldn't fail with %s", "foo")
			assert.Equal(t, "operation shouldn't fail with foo: something went wrong\n", logger.sb.String())
			logger.Flush()
		}

		{
			val1, val2, val3, val4, val5, val6 := May6(1, 2, 3, 4, 5, 6, nil)
			assert.Equal(t, 1, val1)
			assert.Equal(t, 2, val2)
			assert.Equal(t, 3, val3)
			assert.Equal(t, 4, val4)
			assert.Equal(t, 5, val5)
			assert.Equal(t, 6, val6)

			May6(1, 2, 3, 4, 5, 6, errors.New("something went wrong"))
			assert.Equal(t, "something went wrong\n", logger.sb.String())
			logger.Flush()

			May6(1, 2, 3, 4, 5, 6, errors.New("something went wrong"), "operation shouldn't fail with %s", "foo")
			assert.Equal(t, "operation shouldn't fail with foo: something went wrong\n", logger.sb.String())
			logger.Flush()
		}
	})

	t.Run("Bool", func(t *testing.T) {
		logger := newMayTestLogger()
		internalMayHandlers.Use(WithLoggerHandler(logger))

		defer func() {
			internalMayHandlers.handlers = make([]MayHandler, 0)
		}()

		{
			May0(false)
			assert.Equal(t, "not ok\n", logger.sb.String())
			logger.Flush()

			May0(false, "operation shouldn't fail with %s", "foo")
			assert.Equal(t, "operation shouldn't fail with foo\n", logger.sb.String())
			logger.Flush()

			May0(true)
			assert.Empty(t, logger.sb.String())
			logger.Flush()
		}

		{
			val1 := May1(1, true)
			assert.Equal(t, 1, val1)

			May1(1, false)
			assert.Equal(t, "not ok\n", logger.sb.String())
			logger.Flush()

			May1(1, false, "operation shouldn't fail with %s", "foo")
			assert.Equal(t, "operation shouldn't fail with foo\n", logger.sb.String())
			logger.Flush()
		}

		{
			val1, val2 := May2(1, 2, true)
			assert.Equal(t, 1, val1)
			assert.Equal(t, 2, val2)

			May2(1, 2, false)
			assert.Equal(t, "not ok\n", logger.sb.String())
			logger.Flush()

			May2(1, 2, false, "operation shouldn't fail with %s", "foo")
			assert.Equal(t, "operation shouldn't fail with foo\n", logger.sb.String())
			logger.Flush()
		}

		{
			val1, val2, val3 := May3(1, 2, 3, true)
			assert.Equal(t, 1, val1)
			assert.Equal(t, 2, val2)
			assert.Equal(t, 3, val3)

			May3(1, 2, 3, false)
			assert.Equal(t, "not ok\n", logger.sb.String())
			logger.Flush()

			May3(1, 2, 3, false, "operation shouldn't fail with %s", "foo")
			assert.Equal(t, "operation shouldn't fail with foo\n", logger.sb.String())
			logger.Flush()
		}

		{
			val1, val2, val3, val4 := May4(1, 2, 3, 4, true)
			assert.Equal(t, 1, val1)
			assert.Equal(t, 2, val2)
			assert.Equal(t, 3, val3)
			assert.Equal(t, 4, val4)

			May4(1, 2, 3, 4, false)
			assert.Equal(t, "not ok\n", logger.sb.String())
			logger.Flush()

			May4(1, 2, 3, 4, false, "operation shouldn't fail with %s", "foo")
			assert.Equal(t, "operation shouldn't fail with foo\n", logger.sb.String())
			logger.Flush()
		}

		{
			val1, val2, val3, val4, val5 := May5(1, 2, 3, 4, 5, true)
			assert.Equal(t, 1, val1)
			assert.Equal(t, 2, val2)
			assert.Equal(t, 3, val3)
			assert.Equal(t, 4, val4)
			assert.Equal(t, 5, val5)

			May5(1, 2, 3, 4, 5, false)
			assert.Equal(t, "not ok\n", logger.sb.String())
			logger.Flush()

			May5(1, 2, 3, 4, 5, false, "operation shouldn't fail with %s", "foo")
			assert.Equal(t, "operation shouldn't fail with foo\n", logger.sb.String())
			logger.Flush()
		}

		{
			val1, val2, val3, val4, val5, val6 := May6(1, 2, 3, 4, 5, 6, true)
			assert.Equal(t, 1, val1)
			assert.Equal(t, 2, val2)
			assert.Equal(t, 3, val3)
			assert.Equal(t, 4, val4)
			assert.Equal(t, 5, val5)
			assert.Equal(t, 6, val6)

			May6(1, 2, 3, 4, 5, 6, false)
			assert.Equal(t, "not ok\n", logger.sb.String())
			logger.Flush()

			May6(1, 2, 3, 4, 5, 6, false, "operation shouldn't fail with %s", "foo")
			assert.Equal(t, "operation shouldn't fail with foo\n", logger.sb.String())
			logger.Flush()
		}
	})
}

func BenchmarkMay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		May("string", errors.New("error"))
	}
}
