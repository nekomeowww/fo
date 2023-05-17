package fo

import "sync"

var (
	mutex               = sync.Mutex{}
	internalMayHandlers = newMayHandlers()
)

// SetLoggers sets the global loggers for May and May* functions.
//
// NOTICE: This function will replace all the global existing handlers on
// package fo level.
func SetLoggers(logger ...Logger) {
	handlers := make([]MayHandler, 0, len(logger))
	for _, l := range logger {
		handlers = append(handlers, WithLoggerHandler(l))
	}

	SetHandlers(handlers...)
}

// SetHandlers sets the global handlers for May and May* functions.
//
// NOTICE: This function will replace all the global existing handlers on
// package fo level.
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
