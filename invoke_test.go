package fo

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInvoke(t *testing.T) {
	t.Parallel()

	t.Run("TODO", func(t *testing.T) {
		t.Parallel()

		res, err := invoke(context.TODO(), func() (any, error) {
			return nil, assert.AnError
		})

		require.Error(t, err)
		require.Nil(t, res)
		assert.ErrorIs(t, err, assert.AnError)
	})

	t.Run("Background", func(t *testing.T) {
		t.Parallel()

		res, err := invoke(context.Background(), func() (any, error) {
			return nil, assert.AnError
		})

		require.Error(t, err)
		require.Nil(t, res)
		assert.ErrorIs(t, err, assert.AnError)
	})

	t.Run("WithCancel", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithCancel(context.Background())

		time.AfterFunc(time.Millisecond*500, cancel)

		res, err := invoke(ctx, func() (any, error) {
			time.Sleep(time.Second)
			return nil, nil
		})

		require.Error(t, err)
		require.Nil(t, res)
		assert.ErrorIs(t, err, context.Canceled)
	})

	t.Run("WithCancelCause", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithCancelCause(context.Background())

		time.AfterFunc(time.Millisecond*500, func() {
			cancel(assert.AnError)
		})

		res, err := invoke(ctx, func() (any, error) {
			time.Sleep(time.Second)
			return nil, nil
		})

		require.Error(t, err)
		require.Nil(t, res)
		assert.ErrorIs(t, err, context.Canceled)

		err = context.Cause(ctx)
		require.Error(t, err)
		assert.ErrorIs(t, err, assert.AnError)
	})

	t.Run("WithTimeout", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
		defer cancel()

		res, err := invoke(ctx, func() (any, error) {
			time.Sleep(time.Second)
			return nil, nil
		})

		require.Error(t, err)
		require.Nil(t, res)
		assert.ErrorIs(t, err, context.DeadlineExceeded)
	})

	t.Run("WithDeadline", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Millisecond*500))
		defer cancel()

		res, err := invoke(ctx, func() (any, error) {
			time.Sleep(time.Second)
			return nil, nil
		})

		require.Error(t, err)
		require.Nil(t, res)
		assert.ErrorIs(t, err, context.DeadlineExceeded)
	})
}

func TestInvokeX(t *testing.T) {
	t.Parallel()

	t.Run("Invoke0", func(t *testing.T) {
		t.Parallel()

		t.Run("Error", func(t *testing.T) {
			t.Parallel()

			err := Invoke0(context.Background(), func() error {
				return assert.AnError
			})

			require.Error(t, err)
			assert.ErrorIs(t, err, assert.AnError)
		})

		t.Run("Return", func(t *testing.T) {
			t.Parallel()

			err := Invoke0(context.Background(), func() error {
				return nil
			})

			require.NoError(t, err)
		})
	})

	t.Run("Invoke", func(t *testing.T) {
		t.Parallel()

		t.Run("Error", func(t *testing.T) {
			t.Parallel()

			res, err := Invoke(context.Background(), func() (string, error) {
				return "", assert.AnError
			})

			require.Error(t, err)
			require.Empty(t, res)
			assert.ErrorIs(t, err, assert.AnError)
		})

		t.Run("Return", func(t *testing.T) {
			t.Parallel()

			res, err := Invoke(context.Background(), func() (string, error) {
				return "foo", nil
			})

			require.NoError(t, err)
			assert.Equal(t, "foo", res)
		})
	})

	t.Run("Invoke1", func(t *testing.T) {
		t.Parallel()

		t.Run("Error", func(t *testing.T) {
			t.Parallel()
			res, err := Invoke1(context.Background(), func() (string, error) {
				return "", assert.AnError
			})

			require.Error(t, err)
			require.Empty(t, res)
			assert.ErrorIs(t, err, assert.AnError)
		})

		t.Run("Return", func(t *testing.T) {
			t.Parallel()

			res, err := Invoke1(context.Background(), func() (string, error) {
				return "foo", nil
			})

			require.NoError(t, err)
			assert.Equal(t, "foo", res)
		})
	})

	t.Run("Invoke2", func(t *testing.T) {
		t.Parallel()

		t.Run("Error", func(t *testing.T) {
			t.Parallel()
			res1, res2, err := Invoke2(context.Background(), func() (string, int, error) {
				return "", 0, assert.AnError
			})

			require.Error(t, err)
			require.Empty(t, res1)
			require.Zero(t, res2)
			assert.ErrorIs(t, err, assert.AnError)
		})

		t.Run("Return", func(t *testing.T) {
			t.Parallel()

			res1, res2, err := Invoke2(context.Background(), func() (string, int, error) {
				return "foo", 42, nil
			})

			require.NoError(t, err)
			assert.Equal(t, "foo", res1)
			assert.Equal(t, 42, res2)
		})
	})

	t.Run("Invoke3", func(t *testing.T) {
		t.Parallel()

		t.Run("Error", func(t *testing.T) {
			t.Parallel()
			res1, res2, res3, err := Invoke3(context.Background(), func() (string, int, bool, error) {
				return "", 0, false, assert.AnError
			})

			require.Error(t, err)
			require.Empty(t, res1)
			require.Zero(t, res2)
			require.False(t, res3)
			assert.ErrorIs(t, err, assert.AnError)
		})

		t.Run("Return", func(t *testing.T) {
			t.Parallel()

			res1, res2, res3, err := Invoke3(context.Background(), func() (string, int, bool, error) {
				return "foo", 42, true, nil
			})

			require.NoError(t, err)
			assert.Equal(t, "foo", res1)
			assert.Equal(t, 42, res2)
			assert.True(t, res3)
		})
	})

	t.Run("Invoke4", func(t *testing.T) {
		t.Parallel()

		t.Run("Error", func(t *testing.T) {
			t.Parallel()
			res1, res2, res3, res4, err := Invoke4(context.Background(), func() (string, int, bool, float64, error) {
				return "", 0, false, 0, assert.AnError
			})

			require.Error(t, err)
			require.Empty(t, res1)
			require.Zero(t, res2)
			require.False(t, res3)
			require.Zero(t, res4)
			assert.ErrorIs(t, err, assert.AnError)
		})

		t.Run("Return", func(t *testing.T) {
			t.Parallel()

			res1, res2, res3, res4, err := Invoke4(context.Background(), func() (string, int, bool, float64, error) {
				return "foo", 42, true, 3.14, nil
			})

			require.NoError(t, err)
			assert.Equal(t, "foo", res1)
			assert.Equal(t, 42, res2)
			assert.True(t, res3)
			assert.Equal(t, 3.14, res4)
		})
	})

	t.Run("Invoke5", func(t *testing.T) {
		t.Parallel()

		type testStruct struct {
			name string
		}

		t.Run("Error", func(t *testing.T) {
			t.Parallel()
			res1, res2, res3, res4, res5, err := Invoke5(context.Background(), func() (string, int, bool, float64, *testStruct, error) {
				return "", 0, false, 0, nil, assert.AnError
			})

			require.Error(t, err)
			require.Empty(t, res1)
			require.Zero(t, res2)
			require.False(t, res3)
			require.Zero(t, res4)
			require.Nil(t, res5)
			assert.ErrorIs(t, err, assert.AnError)
		})

		t.Run("Return", func(t *testing.T) {
			t.Parallel()

			res1, res2, res3, res4, res5, err := Invoke5(context.Background(), func() (string, int, bool, float64, *testStruct, error) {
				return "foo", 42, true, 3.14, &testStruct{name: "bar"}, nil
			})

			require.NoError(t, err)
			assert.Equal(t, "foo", res1)
			assert.Equal(t, 42, res2)
			assert.True(t, res3)
			assert.Equal(t, 3.14, res4)
			assert.Equal(t, "bar", res5.name)
		})
	})

	t.Run("Invoke6", func(t *testing.T) {
		t.Parallel()

		type testStruct struct {
			name string
		}

		t.Run("Error", func(t *testing.T) {
			t.Parallel()
			res1, res2, res3, res4, res5, res6, err := Invoke6(context.Background(), func() (string, int, bool, float64, *testStruct, map[string]string, error) {
				return "", 0, false, 0, nil, nil, assert.AnError
			})

			require.Error(t, err)
			require.Empty(t, res1)
			require.Zero(t, res2)
			require.False(t, res3)
			require.Zero(t, res4)
			require.Nil(t, res5)
			require.Nil(t, res6)
			assert.ErrorIs(t, err, assert.AnError)
		})

		t.Run("Return", func(t *testing.T) {
			t.Parallel()

			res1, res2, res3, res4, res5, res6, err := Invoke6(context.Background(), func() (string, int, bool, float64, *testStruct, map[string]string, error) {
				return "foo", 42, true, 3.14, &testStruct{name: "bar"}, map[string]string{"foo": "bar"}, nil
			})

			require.NoError(t, err)
			assert.Equal(t, "foo", res1)
			assert.Equal(t, 42, res2)
			assert.True(t, res3)
			assert.Equal(t, 3.14, res4)
			assert.Equal(t, "bar", res5.name)
			assert.Equal(t, "bar", res6["foo"])
		})
	})
}

func ExampleInvoke() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()

	str, err := Invoke(ctx, func() (int, error) {
		time.Sleep(time.Second)
		return 0, nil
	})

	fmt.Println(str, err)
	// Output: 0 context deadline exceeded
}

func BenchmarkInvoke(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = invoke(context.Background(), func() (any, error) {
			return "string", errors.New("error")
		})
	}
}
