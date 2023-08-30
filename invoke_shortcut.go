package fo

import (
	"context"
	"time"
)

type invokeWithOptions struct {
	contextTimeout      time.Duration
	contextTimeoutIsSet bool
}

type callInvokeWithOptionType int

const (
	callInvokeWithOptionTypeContextDefault callInvokeWithOptionType = iota
	callInvokeWithOptionTypeContextTimeout
)

type CallInvokeWithOption struct {
	optionType callInvokeWithOptionType
	options    func() *invokeWithOptions
}

func WithContextTimeout(timeout time.Duration) CallInvokeWithOption {
	return CallInvokeWithOption{
		optionType: callInvokeWithOptionTypeContextTimeout,
		options: func() *invokeWithOptions {
			return &invokeWithOptions{
				contextTimeout:      timeout,
				contextTimeoutIsSet: true,
			}
		},
	}
}

func invokeWithCallOptions[R any](fn func() (R, error), callOpts ...CallInvokeWithOption) (R, error) {
	ctx := context.Background()
	cancelFuncs := make([]context.CancelFunc, 0, len(callOpts))

	for _, callOpt := range callOpts {
		options := callOpt.options()
		if options.contextTimeoutIsSet {
			timeout := options.contextTimeout
			if timeout <= 0 {
				continue
			}

			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(ctx, timeout)
			cancelFuncs = append(cancelFuncs, cancel)
		}
	}

	defer func() {
		for _, cancel := range cancelFuncs {
			cancel()
		}
	}()

	return invoke(ctx, fn)
}

// InvokeWith0 has the same behavior as InvokeWith but without return value.
func InvokeWith0(fn func() error, opts ...CallInvokeWithOption) error {
	_, err := invokeWithCallOptions(func() (any, error) {
		return nil, fn()
	}, opts...)

	return err
}

// InvokeWithTimeout0 has the same behavior as InvokeWithTimeout but without return value.
func InvokeWithTimeout0(fn func() error, timeout time.Duration) error {
	return InvokeWith0(fn, WithContextTimeout(timeout))
}

// InvokeWith invokes the callback function with the CallInvokeWithOption passed in and set for
// context.Background() as parent context and enables to control the context of the callback
// function with 1 return value and an error.
func InvokeWith[R1 any](fn func() (R1, error), opts ...CallInvokeWithOption) (R1, error) {
	return InvokeWith1(fn, opts...)
}

// InvokeWithTimeout invokes the callback function with the timeout passed in and set for
// context.Background() as parent context with context.WithTimeout(...) and enables to
// control the timeout context of the callback function with 1 return value and an error.
func InvokeWithTimeout[R1 any](fn func() (R1, error), timeout time.Duration) (R1, error) {
	return InvokeWithTimeout1(fn, timeout)
}

// InvokeWith1 is an alias of InvokeWith.
func InvokeWith1[R1 any](fn func() (R1, error), opts ...CallInvokeWithOption) (R1, error) {
	type result struct {
		r1 R1
	}

	res, err := invokeWithCallOptions(func() (result, error) {
		r1, err := fn()
		return result{r1: r1}, err
	}, opts...)

	return res.r1, err
}

// InvokeWithTimeout1 is an alias of InvokeWithTimeout.
func InvokeWithTimeout1[R1 any](fn func() (R1, error), timeout time.Duration) (R1, error) {
	return InvokeWith1(fn, WithContextTimeout(timeout))
}

// InvokeWith2 has the same behavior as InvokeWith but with 2 return values.
func InvokeWith2[R1 any, R2 any](fn func() (R1, R2, error), opts ...CallInvokeWithOption) (R1, R2, error) {
	type result struct {
		r1 R1
		r2 R2
	}

	res, err := invokeWithCallOptions(func() (result, error) {
		r1, r2, err := fn()
		return result{r1: r1, r2: r2}, err
	}, opts...)

	return res.r1, res.r2, err
}

// InvokeWithTimeout2 has the same behavior as InvokeWithTimeout but with 2 return values.
func InvokeWithTimeout2[R1 any, R2 any](fn func() (R1, R2, error), timeout time.Duration) (R1, R2, error) {
	return InvokeWith2(fn, WithContextTimeout(timeout))
}

// InvokeWith3 has the same behavior as InvokeWith but with 3 return values.
func InvokeWith3[R1 any, R2 any, R3 any](fn func() (R1, R2, R3, error), opts ...CallInvokeWithOption) (R1, R2, R3, error) {
	type result struct {
		r1 R1
		r2 R2
		r3 R3
	}

	res, err := invokeWithCallOptions(func() (result, error) {
		r1, r2, r3, err := fn()
		return result{r1: r1, r2: r2, r3: r3}, err
	}, opts...)

	return res.r1, res.r2, res.r3, err
}

// InvokeWithTimeout3 has the same behavior as InvokeWithTimeout but with 3 return values.
func InvokeWithTimeout3[R1 any, R2 any, R3 any](fn func() (R1, R2, R3, error), timeout time.Duration) (R1, R2, R3, error) {
	return InvokeWith3(fn, WithContextTimeout(timeout))
}

// InvokeWith4 has the same behavior as InvokeWith but with 4 return values.
func InvokeWith4[R1 any, R2 any, R3 any, R4 any](fn func() (R1, R2, R3, R4, error), opts ...CallInvokeWithOption) (R1, R2, R3, R4, error) {
	type result struct {
		r1 R1
		r2 R2
		r3 R3
		r4 R4
	}

	res, err := invokeWithCallOptions(func() (result, error) {
		r1, r2, r3, r4, err := fn()
		return result{r1: r1, r2: r2, r3: r3, r4: r4}, err
	}, opts...)

	return res.r1, res.r2, res.r3, res.r4, err
}

// InvokeWithTimeout4 has the same behavior as InvokeWithTimeout but with 4 return values.
func InvokeWithTimeout4[R1 any, R2 any, R3 any, R4 any](fn func() (R1, R2, R3, R4, error), timeout time.Duration) (R1, R2, R3, R4, error) {
	return InvokeWith4(fn, WithContextTimeout(timeout))
}

// InvokeWith5 has the same behavior as InvokeWith but with 5 return values.
func InvokeWith5[R1 any, R2 any, R3 any, R4 any, R5 any](fn func() (R1, R2, R3, R4, R5, error), opts ...CallInvokeWithOption) (R1, R2, R3, R4, R5, error) {
	type result struct {
		r1 R1
		r2 R2
		r3 R3
		r4 R4
		r5 R5
	}

	res, err := invokeWithCallOptions(func() (result, error) {
		r1, r2, r3, r4, r5, err := fn()
		return result{r1: r1, r2: r2, r3: r3, r4: r4, r5: r5}, err
	}, opts...)

	return res.r1, res.r2, res.r3, res.r4, res.r5, err
}

// InvokeWithTimeout5 has the same behavior as InvokeWithTimeout but with 5 return values.
func InvokeWithTimeout5[R1 any, R2 any, R3 any, R4 any, R5 any](fn func() (R1, R2, R3, R4, R5, error), timeout time.Duration) (R1, R2, R3, R4, R5, error) {
	return InvokeWith5(fn, WithContextTimeout(timeout))
}

// InvokeWith6 has the same behavior as InvokeWith but with 6 return values.
func InvokeWith6[R1 any, R2 any, R3 any, R4 any, R5 any, R6 any](fn func() (R1, R2, R3, R4, R5, R6, error), opts ...CallInvokeWithOption) (R1, R2, R3, R4, R5, R6, error) {
	type result struct {
		r1 R1
		r2 R2
		r3 R3
		r4 R4
		r5 R5
		r6 R6
	}

	res, err := invokeWithCallOptions(func() (result, error) {
		r1, r2, r3, r4, r5, r6, err := fn()
		return result{r1: r1, r2: r2, r3: r3, r4: r4, r5: r5, r6: r6}, err
	}, opts...)

	return res.r1, res.r2, res.r3, res.r4, res.r5, res.r6, err
}

// InvokeWithTimeout6 has the same behavior as InvokeWithTimeout but with 6 return values.
func InvokeWithTimeout6[R1 any, R2 any, R3 any, R4 any, R5 any, R6 any](fn func() (R1, R2, R3, R4, R5, R6, error), timeout time.Duration) (R1, R2, R3, R4, R5, R6, error) {
	return InvokeWith6(fn, WithContextTimeout(timeout))
}
