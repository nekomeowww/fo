package fo

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/multierr"
)

func ExampleCollectAsError() {
	// Lets assume we have 3 functions that return a string and an error.
	func1 := func() (string, error) {
		return "", errors.New("something went wrong")
	}
	func2 := func() (string, error) {
		return "", errors.New("another thing went wrong")
	}
	// This one returns a string and nil error.
	func3 := func() (string, error) {
		return "success", nil
	}

	useCase1 := func() (s1 string, s2 string, s3 string, err error) {
		may := NewMay1[string]()

		// Instead of check and handle errors manually after each invocation,
		// we can use CollectAsError to do it for us to collect all errors and
		// return them as one error at the end.
		defer func() {
			err = CollectAsError(may)
		}()

		res1 := may.Invoke(func1())
		res2 := may.Invoke(func2())
		res3 := may.Invoke(func3())

		return res1, res2, res3, nil
	}

	useCase2 := func() (s1 string, s2 string, s3 string, err error) {
		may := NewMay1[string]()

		res1 := may.Invoke(func1())
		res2 := may.Invoke(func2())
		res3 := may.Invoke(func3())

		// Instead of handling the errors manually, we can use CollectAsError
		// to do it for us and return the collected errors.
		return res1, res2, res3, CollectAsError(may)
	}

	c11, c12, c13, err := useCase1()
	fmt.Println(strings.Join([]string{c11, c12, c13}, ","))
	fmt.Println(err)

	c21, c22, c23, err := useCase2()
	fmt.Println(strings.Join([]string{c21, c22, c23}, ","))
	fmt.Println(err)
	// Output: ,,success
	// something went wrong; another thing went wrong
	// ,,success
	// something went wrong; another thing went wrong
}

func ExampleHandleErrors() {
	// Lets assume we have 3 functions that return a string and an error.
	func1 := func() (string, error) {
		return "", errors.New("something went wrong")
	}
	func2 := func() (string, error) {
		return "", errors.New("another thing went wrong")
	}
	// This one returns a string and nil error.
	func3 := func() (string, error) {
		return "success", nil
	}

	// This is the use case we want to implement.
	useCase1 := func() (s1 string, s2 string, s3 string, err error) {
		may := NewMay1[string]()

		// Instead of check and handle errors manually after each invocation,
		// we can use HandleErrors to do it for us.
		// To do so, we need to defer HandleErrors with the MayInvoker as argument
		// while we defined the named return err error to be able to return the
		// collected errors.
		defer func() {
			HandleErrors(may, func(errs []error) {
				err = fmt.Errorf("something went wrong: %d errors occurred", len(errs))
			})
		}()

		res1 := may.Invoke(func1())
		res2 := may.Invoke(func2())
		res3 := may.Invoke(func3())

		return res1, res2, res3, nil
	}

	res1, res2, res3, err := useCase1()
	fmt.Println(strings.Join([]string{res1, res2, res3}, ","))
	fmt.Println(err)
	// Output: ,,success
	// something went wrong: 2 errors occurred
}

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

func TestCollectAsError(t *testing.T) {
	may := NewMay[string]()

	may.Invoke("", errors.New("something went wrong"))
	may.Invoke("", errors.New("something went wrong"), "operation shouldn't fail")
	may.Invoke("", errors.New("something went wrong"), "operation shouldn't fail with %s", "foo")

	err := CollectAsError(may)
	assert.EqualError(t, err, "something went wrong; operation shouldn't fail: something went wrong; operation shouldn't fail with foo: something went wrong")
}

func TestCollectAsErrors(t *testing.T) {
	may := NewMay[string]()

	may.Invoke("", errors.New("something went wrong"))
	may.Invoke("", errors.New("something went wrong"), "operation shouldn't fail")
	may.Invoke("", errors.New("something went wrong"), "operation shouldn't fail with %s", "foo")

	errs := CollectAsErrors(may)
	assert.EqualError(t, errs[0], "something went wrong")
	assert.EqualError(t, errs[1], "operation shouldn't fail: something went wrong")
	assert.EqualError(t, errs[2], "operation shouldn't fail with foo: something went wrong")
}

func handleErrorTestFunc() (err error) {
	may := NewMay[string]()

	defer HandleErrors(may, func(errs []error) {
		err = fmt.Errorf("error occurred: %w", multierr.Combine(errs...))
	})

	may.Invoke("", errors.New("something went wrong"))
	may.Invoke("", errors.New("something went wrong"), "operation shouldn't fail")
	may.Invoke("", errors.New("something went wrong"), "operation shouldn't fail with %s", "foo")

	return nil
}

func TestHandleErrors(t *testing.T) {
	err := handleErrorTestFunc()
	assert.EqualError(t, err, "error occurred: something went wrong; operation shouldn't fail: something went wrong; operation shouldn't fail with foo: something went wrong")
}

func handleErrorWithReturnTestFunc() (err error) {
	may := NewMay[string]()

	defer func() {
		err = HandleErrorsWithReturn(may, func(errs []error) error {
			return fmt.Errorf("error occurred: %w", multierr.Combine(errs...))
		})
	}()

	may.Invoke("", errors.New("something went wrong"))
	may.Invoke("", errors.New("something went wrong"), "operation shouldn't fail")
	may.Invoke("", errors.New("something went wrong"), "operation shouldn't fail with %s", "foo")

	return nil
}

func TestHandleErrorsWithReturn(t *testing.T) {
	err := handleErrorWithReturnTestFunc()
	assert.EqualError(t, err, "error occurred: something went wrong; operation shouldn't fail: something went wrong; operation shouldn't fail with foo: something went wrong")
}
