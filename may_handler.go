package fo

import (
	"errors"
	"fmt"
	"reflect"
	"sync"

	"go.uber.org/multierr"
)

var (
	// errNotOK is the error returned by MayInvokers when the invoked.
	errNotOK = errors.New("not ok")
)

type MayHandler func(err error, messageArgs ...any)

type Logger interface {
	Error(v ...any)
}

// WithLoggerHandler returns a MayHandler that logs the error with the
// given logger.
func WithLoggerHandler(l Logger) MayHandler {
	return func(err error, messageArgs ...any) {
		l.Error(formatErrorWithMessageArgs(err, messageArgs...))
	}
}

// WithLogFuncHandler returns a MayHandler that logs the error with the
// given logFunc which accepts variadic arguments.
func WithLogFuncHandler(logFunc func(...any)) MayHandler {
	return func(err error, messageArgs ...any) {
		logFunc(formatErrorWithMessageArgs(err, messageArgs...))
	}
}

var (
	// Implementation check for mayHandlers to implement ErrorsCollectable.
	_ ErrorsCollectable = (*mayHandlers)(nil)
)

// mayHandlers is a collection of MayHandler.
type mayHandlers struct {
	handlers []MayHandler
	errs     error

	mutex sync.Mutex
}

// newMayHandlers creates a new mayHandlers.
func newMayHandlers() *mayHandlers {
	return &mayHandlers{
		handlers: make([]MayHandler, 0),
	}
}

// Use registers the handlers.
func (h *mayHandlers) Use(handler ...MayHandler) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.handlers = append(h.handlers, handler...)
}

// messageFromMsgAndArgs constructs the message from the given msgAndArgs.
// If the first argument is a string, it will be used as the message.
// If the first argument is not a string, it will be formatted with
// fmt.Sprintf("%+v", ...) along with the rest of the arguments.
func messageFromMsgAndArgs(msgAndArgs ...any) string {
	if len(msgAndArgs) == 1 {
		if msgAsStr, ok := msgAndArgs[0].(string); ok {
			return msgAsStr
		}

		return fmt.Sprintf("%+v", msgAndArgs[0])
	}

	if len(msgAndArgs) > 1 {
		msgAsStr, ok := msgAndArgs[0].(string)
		if ok {
			return fmt.Sprintf(msgAsStr, msgAndArgs[1:]...)
		}

		return fmt.Sprintf("%+v", msgAndArgs)
	}

	return ""
}

// formatErrorWithMessageArgs formats the error with the given messageArgs.
// If the error is not nil, it will be wrapped with the message.
func formatErrorWithMessageArgs(err error, messageArgs ...any) error {
	if err == nil {
		return nil
	}

	message := messageFromMsgAndArgs(messageArgs...)
	if errors.Is(err, errNotOK) && message != "" {
		return errors.New(message)
	}
	if message != "" {
		return fmt.Errorf("%s: %w", message, err)
	}

	return err
}

// formatError formats the error.
// It supports two types of error:
// 1. bool: if the bool is false, it will be converted to errNotOK.
// 2. error: if the error is not nil, it will be returned as is.
// Otherwise, it will panic due to invalid error type.
func formatError(err any) error {
	if err == nil {
		return nil
	}

	switch e := err.(type) {
	case bool:
		if !e {
			return errNotOK
		}

		return nil
	case error:
		return e
	default:
		panic("may: invalid err type '" + fmt.Sprintf("%v", reflect.TypeOf(err)) + "', should either be a bool or an error")
	}
}

// handleError handles the error with the registered handlers.
func (h *mayHandlers) handleError(anyErr any, messageArgs ...any) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	err := formatError(anyErr)
	if err == nil {
		return
	}

	for _, handler := range h.handlers {
		if handler != nil {
			handler(err, messageArgs...)
		}
	}

	err = formatErrorWithMessageArgs(err, messageArgs...)
	h.errs = multierr.Append(h.errs, err)
}

// implement for ErrorsCollectable. Returns errors as a single error,
// can be extracted with multierr.Errors().
func (h *mayHandlers) getCollectedError() error {
	return h.errs
}

// implement for ErrorsHandleable. Returns errors as a slice of errors,
// can be combined with multierr.Combine().
func (h *mayHandlers) getCollectedErrors() []error {
	if h.errs == nil {
		return nil
	}

	return multierr.Errors(h.errs)
}
