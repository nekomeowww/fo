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

type testStruct struct {
	name string
}

type testCase struct {
	name                     string
	functionSeries           string
	targetFunc               any
	invokeElapsedLessOrEqual time.Duration
	expectedError            error
	expectedR1               string
	expectedR2               int
	expectedR3               bool
	expectedR4               float64
	expectedR5               *testStruct
	expectedR6               map[string]string
}

func assertInvokeWith(functionSeries string, t *testing.T, testCases []testCase) {
	for _, tc := range testCases {
		tc := tc

		t.Run(fmt.Sprintf("%s/%s", tc.functionSeries, tc.name), func(t *testing.T) {
			t.Parallel()

			start := time.Now()

			var err error
			var r1 string
			var r2 int
			var r3 bool
			var r4 float64
			var r5 *testStruct
			var r6 map[string]string

			switch tc.functionSeries {
			case functionSeries + "0":
				targetFuncAsserted, ok := tc.targetFunc.(func() error)
				require.True(t, ok)

				err = targetFuncAsserted()
			case functionSeries + "1":
				targetFuncAsserted, ok := tc.targetFunc.(func() (string, error))
				require.True(t, ok)

				r1, err = targetFuncAsserted()
			case functionSeries + "":
				targetFuncAsserted, ok := tc.targetFunc.(func() (string, error))
				require.True(t, ok)

				r1, err = targetFuncAsserted()
			case functionSeries + "2":
				targetFuncAsserted, ok := tc.targetFunc.(func() (string, int, error))
				require.True(t, ok)

				r1, r2, err = targetFuncAsserted()
			case functionSeries + "3":
				targetFuncAsserted, ok := tc.targetFunc.(func() (string, int, bool, error))
				require.True(t, ok)

				r1, r2, r3, err = targetFuncAsserted()
			case functionSeries + "4":
				targetFuncAsserted, ok := tc.targetFunc.(func() (string, int, bool, float64, error))
				require.True(t, ok)

				r1, r2, r3, r4, err = targetFuncAsserted()
			case functionSeries + "5":
				targetFuncAsserted, ok := tc.targetFunc.(func() (string, int, bool, float64, *testStruct, error))
				require.True(t, ok)

				r1, r2, r3, r4, r5, err = targetFuncAsserted()
			case functionSeries + "6":
				targetFuncAsserted, ok := tc.targetFunc.(func() (string, int, bool, float64, *testStruct, map[string]string, error))
				require.True(t, ok)

				r1, r2, r3, r4, r5, r6, err = targetFuncAsserted()
			}

			elapsed := time.Since(start)
			assert.LessOrEqual(t, elapsed, tc.invokeElapsedLessOrEqual)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedR1, r1)
			assert.Equal(t, tc.expectedR2, r2)
			assert.Equal(t, tc.expectedR3, r3)
			assert.Equal(t, tc.expectedR4, r4)
			assert.Equal(t, tc.expectedR5, r5)
			assert.Equal(t, tc.expectedR6, r6)
		})
	}
}

func TestInvokeWith(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{
			name:           "Timeout",
			functionSeries: "InvokeWith0",
			targetFunc: func() error {
				return InvokeWith0(func() error {
					time.Sleep(time.Second)
					return nil
				}, WithContextTimeout(time.Millisecond))
			},
			invokeElapsedLessOrEqual: 10 * time.Millisecond,
			expectedError:            context.DeadlineExceeded,
		},
		{
			name:           "Return",
			functionSeries: "InvokeWith0",
			targetFunc: func() error {
				return InvokeWith0(func() error {
					return nil
				})
			},
			invokeElapsedLessOrEqual: time.Millisecond,
		},
		{
			name:           "Timeout",
			functionSeries: "InvokeWith1",
			targetFunc: func() (string, error) {
				return InvokeWith1(func() (string, error) {
					time.Sleep(time.Second)
					return "abcd", nil
				}, WithContextTimeout(time.Millisecond))
			},
			invokeElapsedLessOrEqual: 10 * time.Millisecond,
			expectedError:            context.DeadlineExceeded,
		},
		{
			name:           "Return",
			functionSeries: "InvokeWith1",
			targetFunc: func() (string, error) {
				return InvokeWith1(func() (string, error) {
					return "abcd", nil
				})
			},
			invokeElapsedLessOrEqual: time.Millisecond,
			expectedR1:               "abcd",
		},
		{
			name:           "Timeout",
			functionSeries: "InvokeWith",
			targetFunc: func() (string, error) {
				return InvokeWith(func() (string, error) {
					time.Sleep(time.Second)
					return "abcd", nil
				}, WithContextTimeout(time.Millisecond))
			},
			invokeElapsedLessOrEqual: 10 * time.Millisecond,
			expectedError:            context.DeadlineExceeded,
		},
		{
			name:           "Return",
			functionSeries: "InvokeWith",
			targetFunc: func() (string, error) {
				return InvokeWith(func() (string, error) {
					return "abcd", nil
				})
			},
			invokeElapsedLessOrEqual: time.Millisecond,
			expectedR1:               "abcd",
		},
		{
			name:           "Timeout",
			functionSeries: "InvokeWith2",
			targetFunc: func() (string, int, error) {
				return InvokeWith2(func() (string, int, error) {
					time.Sleep(time.Second)
					return "abcd", 42, nil
				}, WithContextTimeout(time.Millisecond))
			},
			invokeElapsedLessOrEqual: 10 * time.Millisecond,
			expectedError:            context.DeadlineExceeded,
		},
		{
			name:           "Return",
			functionSeries: "InvokeWith2",
			targetFunc: func() (string, int, error) {
				return InvokeWith2(func() (string, int, error) {
					return "abcd", 42, nil
				})
			},
			invokeElapsedLessOrEqual: time.Millisecond,
			expectedR1:               "abcd",
			expectedR2:               42,
		},
		{
			name:           "Timeout",
			functionSeries: "InvokeWith3",
			targetFunc: func() (string, int, bool, error) {
				return InvokeWith3(func() (string, int, bool, error) {
					time.Sleep(time.Second)
					return "abcd", 42, true, nil
				}, WithContextTimeout(time.Millisecond))
			},
			invokeElapsedLessOrEqual: 10 * time.Millisecond,
			expectedError:            context.DeadlineExceeded,
		},
		{
			name:           "Return",
			functionSeries: "InvokeWith3",
			targetFunc: func() (string, int, bool, error) {
				return InvokeWith3(func() (string, int, bool, error) {
					return "abcd", 42, true, nil
				})
			},
			invokeElapsedLessOrEqual: time.Millisecond,
			expectedR1:               "abcd",
			expectedR2:               42,
			expectedR3:               true,
		},
		{
			name:           "Timeout",
			functionSeries: "InvokeWith4",
			targetFunc: func() (string, int, bool, float64, error) {
				return InvokeWith4(func() (string, int, bool, float64, error) {
					time.Sleep(time.Second)
					return "abcd", 42, true, 42.24, nil
				}, WithContextTimeout(time.Millisecond))
			},
			invokeElapsedLessOrEqual: 10 * time.Millisecond,
			expectedError:            context.DeadlineExceeded,
		},
		{
			name:           "Return",
			functionSeries: "InvokeWith4",
			targetFunc: func() (string, int, bool, float64, error) {
				return InvokeWith4(func() (string, int, bool, float64, error) {
					return "abcd", 42, true, 42.24, nil
				})
			},
			invokeElapsedLessOrEqual: time.Millisecond,
			expectedR1:               "abcd",
			expectedR2:               42,
			expectedR3:               true,
			expectedR4:               42.24,
		},
		{
			name:           "Timeout",
			functionSeries: "InvokeWith5",
			targetFunc: func() (string, int, bool, float64, *testStruct, error) {
				return InvokeWith5(func() (string, int, bool, float64, *testStruct, error) {
					time.Sleep(time.Second)
					return "abcd", 42, true, 42.24, &testStruct{
						name: "foo",
					}, nil
				}, WithContextTimeout(time.Millisecond))
			},
			invokeElapsedLessOrEqual: 10 * time.Millisecond,
			expectedError:            context.DeadlineExceeded,
		},
		{
			name:           "Return",
			functionSeries: "InvokeWith5",
			targetFunc: func() (string, int, bool, float64, *testStruct, error) {
				return InvokeWith5(func() (string, int, bool, float64, *testStruct, error) {
					return "abcd", 42, true, 42.24, &testStruct{
						name: "foo",
					}, nil
				})
			},
			invokeElapsedLessOrEqual: time.Millisecond,
			expectedR1:               "abcd",
			expectedR2:               42,
			expectedR3:               true,
			expectedR4:               42.24,
			expectedR5: &testStruct{
				name: "foo",
			},
		},
		{
			name:           "Timeout",
			functionSeries: "InvokeWith6",
			targetFunc: func() (string, int, bool, float64, *testStruct, map[string]string, error) {
				return InvokeWith6(func() (string, int, bool, float64, *testStruct, map[string]string, error) {
					time.Sleep(time.Second)
					return "abcd", 42, true, 42.24, &testStruct{
							name: "foo",
						}, map[string]string{
							"foo": "bar",
						}, nil
				}, WithContextTimeout(time.Millisecond))
			},
			invokeElapsedLessOrEqual: 10 * time.Millisecond,
			expectedError:            context.DeadlineExceeded,
		},
		{
			name:           "Return",
			functionSeries: "InvokeWith6",
			targetFunc: func() (string, int, bool, float64, *testStruct, map[string]string, error) {
				return InvokeWith6(func() (string, int, bool, float64, *testStruct, map[string]string, error) {
					return "abcd", 42, true, 42.24, &testStruct{
							name: "foo",
						}, map[string]string{
							"foo": "bar",
						}, nil
				})
			},
			invokeElapsedLessOrEqual: time.Millisecond,
			expectedR1:               "abcd",
			expectedR2:               42,
			expectedR3:               true,
			expectedR4:               42.24,
			expectedR5: &testStruct{
				name: "foo",
			},
			expectedR6: map[string]string{
				"foo": "bar",
			},
		},
	}

	assertInvokeWith("InvokeWith", t, testCases)
}

func TestInvokeWithTimeout(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{
			name:           "Timeout",
			functionSeries: "InvokeWithTimeout0",
			targetFunc: func() error {
				return InvokeWithTimeout0(func() error {
					time.Sleep(time.Second)
					return nil
				}, time.Millisecond)
			},
			invokeElapsedLessOrEqual: 10 * time.Millisecond,
			expectedError:            context.DeadlineExceeded,
		},
		{
			name:           "Return",
			functionSeries: "InvokeWithTimeout0",
			targetFunc: func() error {
				return InvokeWithTimeout0(func() error {
					return nil
				}, 0)
			},
			invokeElapsedLessOrEqual: time.Millisecond,
		},
		{
			name:           "Timeout",
			functionSeries: "InvokeWithTimeout1",
			targetFunc: func() (string, error) {
				return InvokeWithTimeout1(func() (string, error) {
					time.Sleep(time.Second)
					return "abcd", nil
				}, time.Millisecond)
			},
			invokeElapsedLessOrEqual: 10 * time.Millisecond,
			expectedError:            context.DeadlineExceeded,
		},
		{
			name:           "Return",
			functionSeries: "InvokeWithTimeout1",
			targetFunc: func() (string, error) {
				return InvokeWithTimeout1(func() (string, error) {
					return "abcd", nil
				}, 0)
			},
			invokeElapsedLessOrEqual: time.Millisecond,
			expectedR1:               "abcd",
		},
		{
			name:           "Timeout",
			functionSeries: "InvokeWithTimeout",
			targetFunc: func() (string, error) {
				return InvokeWithTimeout(func() (string, error) {
					time.Sleep(time.Second)
					return "abcd", nil
				}, time.Millisecond)
			},
			invokeElapsedLessOrEqual: 10 * time.Millisecond,
			expectedError:            context.DeadlineExceeded,
		},
		{
			name:           "Return",
			functionSeries: "InvokeWithTimeout",
			targetFunc: func() (string, error) {
				return InvokeWithTimeout(func() (string, error) {
					return "abcd", nil
				}, 0)
			},
			invokeElapsedLessOrEqual: time.Millisecond,
			expectedR1:               "abcd",
		},
		{
			name:           "Timeout",
			functionSeries: "InvokeWithTimeout2",
			targetFunc: func() (string, int, error) {
				return InvokeWithTimeout2(func() (string, int, error) {
					time.Sleep(time.Second)
					return "abcd", 42, nil
				}, time.Millisecond)
			},
			invokeElapsedLessOrEqual: 10 * time.Millisecond,
			expectedError:            context.DeadlineExceeded,
		},
		{
			name:           "Return",
			functionSeries: "InvokeWithTimeout2",
			targetFunc: func() (string, int, error) {
				return InvokeWithTimeout2(func() (string, int, error) {
					return "abcd", 42, nil
				}, 0)
			},
			invokeElapsedLessOrEqual: time.Millisecond,
			expectedR1:               "abcd",
			expectedR2:               42,
		},
		{
			name:           "Timeout",
			functionSeries: "InvokeWithTimeout3",
			targetFunc: func() (string, int, bool, error) {
				return InvokeWithTimeout3(func() (string, int, bool, error) {
					time.Sleep(time.Second)
					return "abcd", 42, true, nil
				}, time.Millisecond)
			},
			invokeElapsedLessOrEqual: 10 * time.Millisecond,
			expectedError:            context.DeadlineExceeded,
		},
		{
			name:           "Return",
			functionSeries: "InvokeWithTimeout3",
			targetFunc: func() (string, int, bool, error) {
				return InvokeWithTimeout3(func() (string, int, bool, error) {
					return "abcd", 42, true, nil
				}, 0)
			},
			invokeElapsedLessOrEqual: time.Millisecond,
			expectedR1:               "abcd",
			expectedR2:               42,
			expectedR3:               true,
		},
		{
			name:           "Timeout",
			functionSeries: "InvokeWithTimeout4",
			targetFunc: func() (string, int, bool, float64, error) {
				return InvokeWithTimeout4(func() (string, int, bool, float64, error) {
					time.Sleep(time.Second)
					return "abcd", 42, true, 42.24, nil
				}, time.Millisecond)
			},
			invokeElapsedLessOrEqual: 10 * time.Millisecond,
			expectedError:            context.DeadlineExceeded,
		},
		{
			name:           "Return",
			functionSeries: "InvokeWithTimeout4",
			targetFunc: func() (string, int, bool, float64, error) {
				return InvokeWithTimeout4(func() (string, int, bool, float64, error) {
					return "abcd", 42, true, 42.24, nil
				}, 0)
			},
			invokeElapsedLessOrEqual: time.Millisecond,
			expectedR1:               "abcd",
			expectedR2:               42,
			expectedR3:               true,
			expectedR4:               42.24,
		},
		{
			name:           "Timeout",
			functionSeries: "InvokeWithTimeout5",
			targetFunc: func() (string, int, bool, float64, *testStruct, error) {
				return InvokeWithTimeout5(func() (string, int, bool, float64, *testStruct, error) {
					time.Sleep(time.Second)
					return "abcd", 42, true, 42.24, &testStruct{
						name: "foo",
					}, nil
				}, time.Millisecond)
			},
			invokeElapsedLessOrEqual: 10 * time.Millisecond,
			expectedError:            context.DeadlineExceeded,
		},
		{
			name:           "Return",
			functionSeries: "InvokeWithTimeout5",
			targetFunc: func() (string, int, bool, float64, *testStruct, error) {
				return InvokeWithTimeout5(func() (string, int, bool, float64, *testStruct, error) {
					return "abcd", 42, true, 42.24, &testStruct{
						name: "foo",
					}, nil
				}, 0)
			},
			invokeElapsedLessOrEqual: time.Millisecond,
			expectedR1:               "abcd",
			expectedR2:               42,
			expectedR3:               true,
			expectedR4:               42.24,
			expectedR5: &testStruct{
				name: "foo",
			},
		},
		{
			name:           "Timeout",
			functionSeries: "InvokeWithTimeout6",
			targetFunc: func() (string, int, bool, float64, *testStruct, map[string]string, error) {
				return InvokeWithTimeout6(func() (string, int, bool, float64, *testStruct, map[string]string, error) {
					time.Sleep(time.Second)
					return "abcd", 42, true, 42.24, &testStruct{
							name: "foo",
						}, map[string]string{
							"foo": "bar",
						}, nil
				}, time.Millisecond)
			},
			invokeElapsedLessOrEqual: 10 * time.Millisecond,
			expectedError:            context.DeadlineExceeded,
		},
		{
			name:           "Return",
			functionSeries: "InvokeWithTimeout6",
			targetFunc: func() (string, int, bool, float64, *testStruct, map[string]string, error) {
				return InvokeWithTimeout6(func() (string, int, bool, float64, *testStruct, map[string]string, error) {
					return "abcd", 42, true, 42.24, &testStruct{
							name: "foo",
						}, map[string]string{
							"foo": "bar",
						}, nil
				}, 0)
			},
			invokeElapsedLessOrEqual: time.Millisecond,
			expectedR1:               "abcd",
			expectedR2:               42,
			expectedR3:               true,
			expectedR4:               42.24,
			expectedR5: &testStruct{
				name: "foo",
			},
			expectedR6: map[string]string{
				"foo": "bar",
			},
		},
	}

	assertInvokeWith("InvokeWithTimeout", t, testCases)
}

func ExampleInvokeWith() {
	str, err := InvokeWith(func() (int, error) {
		time.Sleep(time.Second)
		return 0, nil
	}, WithContextTimeout(time.Millisecond*500))

	fmt.Println(str, err)
	// Output: 0 context deadline exceeded
}

func ExampleInvokeWithTimeout() {
	str, err := InvokeWithTimeout(func() (int, error) {
		time.Sleep(time.Second)
		return 0, nil
	}, time.Millisecond*500)

	fmt.Println(str, err)
	// Output: 0 context deadline exceeded
}

func BenchmarkInvokeWith(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = invokeWithCallOptions(func() (any, error) {
			return "string", errors.New("error")
		}, WithContextTimeout(time.Millisecond))
	}
}
