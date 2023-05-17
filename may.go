package fo

import "sync"

var (
	mutex               = sync.Mutex{}
	internalMayHandlers = newMayHandlers()
)

// MayHandler is a function that handles the error from May and May* functions.
func SetLoggers(logger ...Logger) {
	mutex.Lock()
	defer mutex.Unlock()
	internalMayHandlers.handlers = make([]MayHandler, 0, len(logger))

	for _, l := range logger {
		internalMayHandlers.handlers = append(internalMayHandlers.handlers, WithLoggerHandler(l))
	}
}

// SetHandlers sets the global handlers for May and May* functions.
func SetHandlers(handler ...MayHandler) {
	mutex.Lock()
	defer mutex.Unlock()
	internalMayHandlers.handlers = make([]MayHandler, 0, len(handler))
	internalMayHandlers.handlers = append(internalMayHandlers.handlers, handler...)
}

// May is a helper that wraps a call to a callback function and
// then filters out the error from the result and only returns
// the value. If the error is not nil, it will be handled by
// handlers registered by SetMayHandlers(...), or be logged by
// default logging handler.
func May[T any](val T, err any, messageArgs ...any) T {
	may := NewMay1[T]()
	may.mayHandlers = internalMayHandlers

	return may.Invoke(val, err, messageArgs...)
}

// May0 has the same behavior as May, but callback returns no variable.
func May0(err any, messageArgs ...any) {
	may := NewMay0()
	may.mayHandlers = internalMayHandlers

	may.Invoke(err, messageArgs...)
}

// May1 is an alias of May.
func May1[T any](t1 T, err any, messageArgs ...any) T {
	return May(t1, err, messageArgs...)
}

// May2 has the same behavior as May, but callback returns 2 variables.
func May2[T1 any, T2 any](t1 T1, t2 T2, err any, messageArgs ...any) (T1, T2) {
	may := NewMay2[T1, T2]()
	may.mayHandlers = internalMayHandlers

	return may.Invoke(t1, t2, err, messageArgs...)
}

// May3 has the same behavior as May, but callback returns 3 variables.
func May3[T1 any, T2 any, T3 any](t1 T1, t2 T2, t3 T3, err any, messageArgs ...any) (T1, T2, T3) {
	may := NewMay3[T1, T2, T3]()
	may.mayHandlers = internalMayHandlers

	return may.Invoke(t1, t2, t3, err, messageArgs...)
}

// May4 has the same behavior as May, but callback returns 4 variables.
func May4[T1 any, T2 any, T3 any, T4 any](t1 T1, t2 T2, t3 T3, t4 T4, err any, messageArgs ...any) (T1, T2, T3, T4) {
	may := NewMay4[T1, T2, T3, T4]()
	may.mayHandlers = internalMayHandlers

	return may.Invoke(t1, t2, t3, t4, err, messageArgs...)
}

// May5 has the same behavior as May, but callback returns 5 variables.
func May5[T1 any, T2 any, T3 any, T4 any, T5 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, err any, messageArgs ...any) (T1, T2, T3, T4, T5) {
	may := NewMay5[T1, T2, T3, T4, T5]()
	may.mayHandlers = internalMayHandlers

	return may.Invoke(t1, t2, t3, t4, t5, err, messageArgs...)
}

// May6 has the same behavior as May, but callback returns 6 variables.
func May6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, err any, messageArgs ...any) (T1, T2, T3, T4, T5, T6) {
	may := NewMay6[T1, T2, T3, T4, T5, T6]()
	may.mayHandlers = internalMayHandlers

	return may.Invoke(t1, t2, t3, t4, t5, t6, err, messageArgs...)
}

// ErrorCollectable is an interface for collecting error,
// the error could be the error combined by multierr.Combine().
type ErrorCollectable interface {
	getCollectedError() error
}

// ErrorsCollectable is an interface for collecting slice of errors.
type ErrorsCollectable interface {
	getCollectedErrors() []error
}

// CollectAsError collects error from the invoked result from MayInvoker
// for post error handling.
//
// The error can be extracted with
//
//	multierr.Errors().
func CollectAsError(collectable ErrorCollectable) error {
	return collectable.getCollectedError()
}

// CollectAsErrors collects errors from the invoked result from
// MayInvoker for post error handling.
//
// The errors can be combined with
//
//	multierr.Combine().
func CollectAsErrors(collectable ErrorsCollectable) []error {
	return collectable.getCollectedErrors()
}

// HandleErrors executes the handler with the collected error from
// MayInvoker.
func HandleErrors(collectable ErrorsCollectable, handler func(errs []error)) {
	handler(collectable.getCollectedErrors())
}

// HandleErrorsWithReturn executes the handler with the collected error from
// MayInvoker, and returns the error that handled by the handler.
func HandleErrorsWithReturn(collectable ErrorsCollectable, handler func(errs []error) error) error {
	return handler(collectable.getCollectedErrors())
}
