package fo

import (
	"context"
)

func invoke[R any](ctx context.Context, fn func() (R, error)) (r R, e error) {
	var res R
	var err error

	resChan := make(chan struct{})

	go func() {
		res, err = fn()
		resChan <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		e = ctx.Err()
	case <-resChan:
		r = res
		e = err
	}

	return
}

// Invoke0 has the same behavior as Invoke but without return value.
func Invoke0(ctx context.Context, fn func() error) error {
	_, err := invoke(ctx, func() (any, error) {
		return nil, fn()
	})

	return err
}

// Invoke invokes the callback function and enables to control the
// context of the callback function with 1 return value.
func Invoke[R1 any](ctx context.Context, fn func() (R1, error)) (R1, error) {
	return Invoke1(ctx, fn)
}

// Invoke1 is an alias of Invoke.
func Invoke1[R1 any](ctx context.Context, fn func() (R1, error)) (R1, error) {
	type result struct {
		r1 R1
	}

	res, err := invoke(ctx, func() (result, error) {
		r1, err := fn()
		return result{r1: r1}, err
	})

	return res.r1, err
}

// Invoke2 has the same behavior as Invoke but with 2 return values.
func Invoke2[R1 any, R2 any](ctx context.Context, fn func() (R1, R2, error)) (R1, R2, error) {
	type result struct {
		r1 R1
		r2 R2
	}

	res, err := invoke(ctx, func() (result, error) {
		r1, r2, err := fn()
		return result{r1: r1, r2: r2}, err
	})

	return res.r1, res.r2, err
}

// Invoke3 has the same behavior as Invoke but with 3 return values.
func Invoke3[R1 any, R2 any, R3 any](ctx context.Context, fn func() (R1, R2, R3, error)) (R1, R2, R3, error) {
	type result struct {
		r1 R1
		r2 R2
		r3 R3
	}

	res, err := invoke(ctx, func() (result, error) {
		r1, r2, r3, err := fn()
		return result{r1: r1, r2: r2, r3: r3}, err
	})

	return res.r1, res.r2, res.r3, err
}

// Invoke4 has the same behavior as Invoke but with 4 return values.
func Invoke4[R1 any, R2 any, R3 any, R4 any](ctx context.Context, fn func() (R1, R2, R3, R4, error)) (R1, R2, R3, R4, error) {
	type result struct {
		r1 R1
		r2 R2
		r3 R3
		r4 R4
	}

	res, err := invoke(ctx, func() (result, error) {
		r1, r2, r3, r4, err := fn()
		return result{r1: r1, r2: r2, r3: r3, r4: r4}, err
	})

	return res.r1, res.r2, res.r3, res.r4, err
}

// Invoke5 has the same behavior as Invoke but with 5 return values.
func Invoke5[R1 any, R2 any, R3 any, R4 any, R5 any](ctx context.Context, fn func() (R1, R2, R3, R4, R5, error)) (R1, R2, R3, R4, R5, error) {
	type result struct {
		r1 R1
		r2 R2
		r3 R3
		r4 R4
		r5 R5
	}

	res, err := invoke(ctx, func() (result, error) {
		r1, r2, r3, r4, r5, err := fn()
		return result{r1: r1, r2: r2, r3: r3, r4: r4, r5: r5}, err
	})

	return res.r1, res.r2, res.r3, res.r4, res.r5, err
}

// Invoke6 has the same behavior as Invoke but with 6 return values.
func Invoke6[R1 any, R2 any, R3 any, R4 any, R5 any, R6 any](ctx context.Context, fn func() (R1, R2, R3, R4, R5, R6, error)) (R1, R2, R3, R4, R5, R6, error) {
	type result struct {
		r1 R1
		r2 R2
		r3 R3
		r4 R4
		r5 R5
		r6 R6
	}

	res, err := invoke(ctx, func() (result, error) {
		r1, r2, r3, r4, r5, r6, err := fn()
		return result{r1: r1, r2: r2, r3: r3, r4: r4, r5: r5, r6: r6}, err
	})

	return res.r1, res.r2, res.r3, res.r4, res.r5, res.r6, err
}
