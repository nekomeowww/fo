package fo

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleNewMay() {
	errFn := func() (string, error) {
		return "", errors.New("an error")
	}

	may := NewMay[string]()
	res1 := may.Invoke(errFn())

	errNilFn := func() (string, error) {
		return "success", nil
	}

	res2 := may.Invoke(errNilFn())

	fmt.Println(strings.Join([]string{res1, res2}, ","))
	// Output: ,success
}

func ExampleMayInvoker_Use() {
	var outErr error

	handler := func(err error, v ...any) {
		outErr = err
	}

	may := NewMay[string]().Use(handler)
	res := may.Invoke(func() (string, error) {
		return "", errors.New("an error")
	}())

	fmt.Println(strings.Join([]string{res, outErr.Error()}, ","))
	// Output:  ,an error
}

func TestNewMayX(t *testing.T) {
	t.Parallel()

	t.Run("May0", func(t *testing.T) {
		t.Parallel()

		handler := func(err error, messageArgs ...any) {
			assert.Equal(t, assert.AnError, err)
			assert.Equal(t, 0, len(messageArgs))
		}

		may := NewMay0().Use(handler)

		errFn := func() error {
			return assert.AnError
		}

		may.Invoke(errFn())

		nilErrFn := func() error {
			return nil
		}

		may.Invoke(nilErrFn())
	})

	t.Run("May", func(t *testing.T) {
		t.Parallel()

		handler := func(err error, messageArgs ...any) {
			assert.Equal(t, assert.AnError, err)
			assert.Equal(t, 0, len(messageArgs))
		}

		may := NewMay[string]().Use(handler)

		errFn := func() (string, error) {
			return "", assert.AnError
		}

		res := may.Invoke(errFn())
		assert.Empty(t, res)

		nilErrFn := func() (string, error) {
			return "success", nil
		}

		res = may.Invoke(nilErrFn())
		assert.Equal(t, "success", res)
	})

	t.Run("May1", func(t *testing.T) {
		t.Parallel()

		handler := func(err error, messageArgs ...any) {
			assert.Equal(t, assert.AnError, err)
			assert.Equal(t, 0, len(messageArgs))
		}

		may := NewMay1[string]().Use(handler)

		errFn := func() (string, error) {
			return "", assert.AnError
		}

		res := may.Invoke(errFn())
		assert.Empty(t, res)

		nilErrFn := func() (string, error) {
			return "success", nil
		}

		res = may.Invoke(nilErrFn())
		assert.Equal(t, "success", res)
	})

	t.Run("May2", func(t *testing.T) {
		t.Parallel()

		handler := func(err error, messageArgs ...any) {
			assert.Equal(t, assert.AnError, err)
			assert.Equal(t, 0, len(messageArgs))
		}

		may := NewMay2[string, int]().Use(handler)

		errFn := func() (string, int, error) {
			return "", 0, assert.AnError
		}

		res1, res2 := may.Invoke(errFn())
		assert.Empty(t, res1)
		assert.Zero(t, res2)

		nilErrFn := func() (string, int, error) {
			return "success", 1, nil
		}

		res1, res2 = may.Invoke(nilErrFn())
		assert.Equal(t, "success", res1)
		assert.Equal(t, 1, res2)
	})

	t.Run("May3", func(t *testing.T) {
		t.Parallel()

		handler := func(err error, messageArgs ...any) {
			assert.Equal(t, assert.AnError, err)
			assert.Equal(t, 0, len(messageArgs))
		}

		may := NewMay3[string, int, bool]().Use(handler)

		errFn := func() (string, int, bool, error) {
			return "", 0, false, assert.AnError
		}

		res1, res2, res3 := may.Invoke(errFn())
		assert.Empty(t, res1)
		assert.Zero(t, res2)
		assert.False(t, res3)

		nilErrFn := func() (string, int, bool, error) {
			return "success", 1, true, nil
		}

		res1, res2, res3 = may.Invoke(nilErrFn())
		assert.Equal(t, "success", res1)
		assert.Equal(t, 1, res2)
		assert.True(t, res3)
	})

	t.Run("May4", func(t *testing.T) {
		t.Parallel()

		handler := func(err error, messageArgs ...any) {
			assert.Equal(t, assert.AnError, err)
			assert.Equal(t, 0, len(messageArgs))
		}

		may := NewMay4[string, int, bool, float64]().Use(handler)

		errFn := func() (string, int, bool, float64, error) {
			return "", 0, false, 0.0, assert.AnError
		}

		res1, res2, res3, res4 := may.Invoke(errFn())
		assert.Empty(t, res1)
		assert.Zero(t, res2)
		assert.False(t, res3)
		assert.Zero(t, res4)

		nilErrFn := func() (string, int, bool, float64, error) {
			return "success", 1, true, 1.1, nil
		}

		res1, res2, res3, res4 = may.Invoke(nilErrFn())
		assert.Equal(t, "success", res1)
		assert.Equal(t, 1, res2)
		assert.True(t, res3)
		assert.Equal(t, 1.1, res4)
	})

	t.Run("May5", func(t *testing.T) {
		t.Parallel()

		handler := func(err error, messageArgs ...any) {
			assert.Equal(t, assert.AnError, err)
			assert.Equal(t, 0, len(messageArgs))
		}

		type testStruct struct {
			name string
		}

		may := NewMay5[string, int, bool, float64, *testStruct]().Use(handler)

		errFn := func() (string, int, bool, float64, *testStruct, error) {
			return "", 0, false, 0.0, nil, assert.AnError
		}

		res1, res2, res3, res4, res5 := may.Invoke(errFn())
		assert.Empty(t, res1)
		assert.Zero(t, res2)
		assert.False(t, res3)
		assert.Zero(t, res4)
		assert.Nil(t, res5)

		nilErrFn := func() (string, int, bool, float64, *testStruct, error) { //nolint:unparam
			return "success", 1, true, 1.1, &testStruct{name: "bar"}, nil
		}

		res1, res2, res3, res4, res5 = may.Invoke(nilErrFn())
		assert.Equal(t, "success", res1)
		assert.Equal(t, 1, res2)
		assert.True(t, res3)
		assert.Equal(t, 1.1, res4)
		assert.Equal(t, "bar", res5.name)
	})

	t.Run("May6", func(t *testing.T) {
		t.Parallel()

		handler := func(err error, messageArgs ...any) {
			assert.Equal(t, assert.AnError, err)
			assert.Equal(t, 0, len(messageArgs))
		}

		type testStruct struct {
			name string
		}

		may := NewMay6[string, int, bool, float64, *testStruct, map[string]string]().Use(handler)

		errFn := func() (string, int, bool, float64, *testStruct, map[string]string, error) {
			return "", 0, false, 0.0, nil, nil, assert.AnError
		}

		res1, res2, res3, res4, res5, res6 := may.Invoke(errFn())
		assert.Empty(t, res1)
		assert.Zero(t, res2)
		assert.False(t, res3)
		assert.Zero(t, res4)
		assert.Nil(t, res5)
		assert.Nil(t, res6)

		nilErrFn := func() (string, int, bool, float64, *testStruct, map[string]string, error) { //nolint:unparam
			return "success", 1, true, 1.1, &testStruct{name: "bar"}, map[string]string{"foo": "bar"}, nil
		}

		res1, res2, res3, res4, res5, res6 = may.Invoke(nilErrFn())
		assert.Equal(t, "success", res1)
		assert.Equal(t, 1, res2)
		assert.True(t, res3)
		assert.Equal(t, 1.1, res4)
		assert.Equal(t, "bar", res5.name)
		assert.Equal(t, "bar", res6["foo"])
	})
}

func BenchmarkNewMay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewMay0()
	}
}

func BenchmarkNewMayInvoke(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewMay[string]().Invoke(func() (string, error) {
			return "string", errors.New("error")
		}())
	}
}

func BenchmarkNewMayUse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewMay0().Use(func(err error, messageArgs ...any) {
			// noop
		})
	}
}
