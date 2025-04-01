package fo

// MayInvoker is a helper instance that enables to invoke a
// call to a callback function and then filters out the error
// from the result and only returns the value. If the error
// is not nil, it will be collected to mayHandlers and then
// handled by handlers registered by Use(...)
type MayInvoker[T any] struct {
	*mayHandlers
}

// NewMay creates a helper instance that enables to invoke a
// call to a callback function and then filters out the error
// from the result and only returns the value. If the error
// is not nil, it will be collected to mayHandlers and then
// handled by handlers registered by Use(...)
func NewMay[T any]() *MayInvoker[T] {
	return &MayInvoker[T]{
		mayHandlers: newMayHandlers(),
	}
}

// NewMay1 is an alias of NewMay.
func NewMay1[T any]() *MayInvoker[T] {
	return NewMay[T]()
}

// Use registers the handlers.
func (f *MayInvoker[T]) Use(handler ...MayHandler) *MayInvoker[T] {
	f.mayHandlers.Use(handler...)
	return f
}

// Invoke invokes the callback function and filters out the error
// from the result and only returns the value. If the error is not
// nil, it will be collected and handled by handlers registered by
// Use(...)
func (f *MayInvoker[T]) Invoke(t1 T, err any, messageArgs ...any) T {
	if err == nil {
		return t1
	}

	f.handleError(err, messageArgs...)

	return t1
}

// MayInvoker0 is a helper instance that behaves like MayInvoker
// , but it allows to invoke a callback function that returns no
// value.
type MayInvoker0 struct {
	*mayHandlers
}

// NewMay0 creates a helper instance behaves like MayInvoker, but
// it allows to invoke a callback function that returns no value.
func NewMay0() *MayInvoker0 {
	return &MayInvoker0{
		mayHandlers: newMayHandlers(),
	}
}

// Use registers the handlers.
func (f *MayInvoker0) Use(handler ...MayHandler) *MayInvoker0 {
	f.mayHandlers.Use(handler...)
	return f
}

// Invoke invokes the callback function and filters out the error
// from the result and only returns the value. If the error is not
// nil, it will be collected and handled by handlers registered by
// Use(...)
func (f *MayInvoker0) Invoke(anyErr any, messageArgs ...any) {
	if anyErr == nil {
		return
	}

	f.handleError(anyErr, messageArgs...)
}

// MayInvoker2 is a helper instance behaves like MayInvoker, but
// it allows to invoke a callback function that returns 2 values.
type MayInvoker2[T1 any, T2 any] struct {
	*mayHandlers
}

// NewMay2 creates a helper instance behaves like MayInvoker, but
// it allows to invoke a callback function that returns 2 values.
func NewMay2[T1 any, T2 any]() *MayInvoker2[T1, T2] {
	return &MayInvoker2[T1, T2]{
		mayHandlers: newMayHandlers(),
	}
}

// Use registers the handlers.
func (f *MayInvoker2[T1, T2]) Use(handler ...MayHandler) *MayInvoker2[T1, T2] {
	f.mayHandlers.Use(handler...)
	return f
}

// Invoke invokes the callback function and filters out the error
// from the result and only returns the value. If the error is not
// nil, it will be collected and handled by handlers registered by
// Use(...)
func (f *MayInvoker2[T1, T2]) Invoke(t1 T1, t2 T2, err any, messageArgs ...any) (T1, T2) {
	if err == nil {
		return t1, t2
	}

	f.handleError(err, messageArgs...)

	return t1, t2
}

// MayInvoker3 is a helper instance behaves like MayInvoker, but
// it allows to invoke a callback function that returns 3 values.
type MayInvoker3[T1 any, T2 any, T3 any] struct {
	*mayHandlers
}

// NewMay3 creates a helper instance behaves like MayInvoker, but
// it allows to invoke a callback function that returns 3 values.
func NewMay3[T1 any, T2 any, T3 any]() *MayInvoker3[T1, T2, T3] {
	return &MayInvoker3[T1, T2, T3]{
		mayHandlers: newMayHandlers(),
	}
}

// Use registers the handlers.
func (f *MayInvoker3[T1, T2, T3]) Use(handler ...MayHandler) *MayInvoker3[T1, T2, T3] {
	f.mayHandlers.Use(handler...)
	return f
}

// Invoke invokes the callback function and filters out the error
// from the result and only returns the value. If the error is not
// nil, it will be collected and handled by handlers registered by
// Use(...)
func (f *MayInvoker3[T1, T2, T3]) Invoke(t1 T1, t2 T2, t3 T3, err any, messageArgs ...any) (T1, T2, T3) {
	if err == nil {
		return t1, t2, t3
	}

	f.handleError(err, messageArgs...)

	return t1, t2, t3
}

// MayInvoker4 is a helper instance behaves like MayInvoker, but
// it allows to invoke a callback function that returns 4 values.
type MayInvoker4[T1 any, T2 any, T3 any, T4 any] struct {
	*mayHandlers
}

// NewMay4 creates a helper instance behaves like MayInvoker, but
// it allows to invoke a callback function that returns 4 values.
func NewMay4[T1 any, T2 any, T3 any, T4 any]() *MayInvoker4[T1, T2, T3, T4] {
	return &MayInvoker4[T1, T2, T3, T4]{
		mayHandlers: newMayHandlers(),
	}
}

// Use registers the handlers.
func (f *MayInvoker4[T1, T2, T3, T4]) Use(handler ...MayHandler) *MayInvoker4[T1, T2, T3, T4] {
	f.mayHandlers.Use(handler...)
	return f
}

// Invoke invokes the callback function and filters out the error
// from the result and only returns the value. If the error is not
// nil, it will be collected and handled by handlers registered by
// Use(...)
func (f *MayInvoker4[T1, T2, T3, T4]) Invoke(t1 T1, t2 T2, t3 T3, t4 T4, err any, messageArgs ...any) (T1, T2, T3, T4) {
	if err == nil {
		return t1, t2, t3, t4
	}

	f.handleError(err, messageArgs...)

	return t1, t2, t3, t4
}

// MayInvoker5 is a helper instance behaves like MayInvoker, but
// it allows to invoke a callback function that returns 5 values.
type MayInvoker5[T1 any, T2 any, T3 any, T4 any, T5 any] struct {
	*mayHandlers
}

// NewMay5 creates a helper instance behaves like MayInvoker, but
// it allows to invoke a callback function that returns 5 values.
func NewMay5[T1 any, T2 any, T3 any, T4 any, T5 any]() *MayInvoker5[T1, T2, T3, T4, T5] {
	return &MayInvoker5[T1, T2, T3, T4, T5]{
		mayHandlers: newMayHandlers(),
	}
}

// Use registers the handlers.
func (f *MayInvoker5[T1, T2, T3, T4, T5]) Use(handler ...MayHandler) *MayInvoker5[T1, T2, T3, T4, T5] {
	f.mayHandlers.Use(handler...)
	return f
}

// Invoke invokes the callback function and filters out the error
// from the result and only returns the value. If the error is not
// nil, it will be collected and handled by handlers registered by
// Use(...)
func (f *MayInvoker5[T1, T2, T3, T4, T5]) Invoke(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, err any, messageArgs ...any) (T1, T2, T3, T4, T5) {
	if err == nil {
		return t1, t2, t3, t4, t5
	}

	f.handleError(err, messageArgs...)

	return t1, t2, t3, t4, t5
}

// MayInvoker6 is a helper instance behaves like MayInvoker, but
// it allows to invoke a callback function that returns 6 values.
type MayInvoker6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any] struct {
	*mayHandlers
}

// NewMay6 creates a helper instance behaves like MayInvoker, but
// it allows to invoke a callback function that returns 6 values.
func NewMay6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any]() *MayInvoker6[T1, T2, T3, T4, T5, T6] {
	return &MayInvoker6[T1, T2, T3, T4, T5, T6]{
		mayHandlers: newMayHandlers(),
	}
}

// Use registers the handlers.
func (f *MayInvoker6[T1, T2, T3, T4, T5, T6]) Use(handler ...MayHandler) *MayInvoker6[T1, T2, T3, T4, T5, T6] {
	f.mayHandlers.Use(handler...)
	return f
}

// Invoke invokes the callback function and filters out the error
// from the result and only returns the value. If the error is not
// nil, it will be collected and handled by handlers registered by
// Use(...)
func (f *MayInvoker6[T1, T2, T3, T4, T5, T6]) Invoke(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, err any, messageArgs ...any) (T1, T2, T3, T4, T5, T6) {
	if err == nil {
		return t1, t2, t3, t4, t5, t6
	}

	f.handleError(err, messageArgs...)

	return t1, t2, t3, t4, t5, t6
}
