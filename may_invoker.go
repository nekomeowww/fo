package fo

type MayInvoker0 struct {
	*mayHandlers
}

func NewMay0() *MayInvoker0 {
	return &MayInvoker0{
		mayHandlers: newMayHandlers(),
	}
}

func (f *MayInvoker0) Use(handler ...MayHandler) *MayInvoker0 {
	f.mayHandlers.Use(handler...)
	return f
}

func (f *MayInvoker0) Invoke(anyErr any, messageArgs ...any) {
	if anyErr == nil {
		return
	}

	f.mayHandlers.handleError(anyErr, messageArgs...)
}

type MayInvoker[T any] struct {
	*mayHandlers
}

func NewMay[T any]() *MayInvoker[T] {
	return &MayInvoker[T]{
		mayHandlers: newMayHandlers(),
	}
}

func NewMay1[T any]() *MayInvoker[T] {
	return NewMay[T]()
}

func (f *MayInvoker[T]) Use(handler ...MayHandler) *MayInvoker[T] {
	f.mayHandlers.Use(handler...)
	return f
}

func (f *MayInvoker[T]) Invoke(t1 T, err any, messageArgs ...any) T {
	if err == nil {
		return t1
	}

	f.mayHandlers.handleError(err, messageArgs...)

	return t1
}

type MayInvoker2[T1 any, T2 any] struct {
	*mayHandlers
}

func NewMay2[T1 any, T2 any]() *MayInvoker2[T1, T2] {
	return &MayInvoker2[T1, T2]{
		mayHandlers: newMayHandlers(),
	}
}

func (f *MayInvoker2[T1, T2]) Use(handler ...MayHandler) *MayInvoker2[T1, T2] {
	f.mayHandlers.Use(handler...)
	return f
}

func (f *MayInvoker2[T1, T2]) Invoke(t1 T1, t2 T2, err any, messageArgs ...any) (T1, T2) {
	if err == nil {
		return t1, t2
	}

	f.mayHandlers.handleError(err, messageArgs...)

	return t1, t2
}

type MayInvoker3[T1 any, T2 any, T3 any] struct {
	*mayHandlers
}

func NewMay3[T1 any, T2 any, T3 any]() *MayInvoker3[T1, T2, T3] {
	return &MayInvoker3[T1, T2, T3]{
		mayHandlers: newMayHandlers(),
	}
}

func (f *MayInvoker3[T1, T2, T3]) Use(handler ...MayHandler) *MayInvoker3[T1, T2, T3] {
	f.mayHandlers.Use(handler...)
	return f
}

func (f *MayInvoker3[T1, T2, T3]) Invoke(t1 T1, t2 T2, t3 T3, err any, messageArgs ...any) (T1, T2, T3) {
	if err == nil {
		return t1, t2, t3
	}

	f.mayHandlers.handleError(err, messageArgs...)

	return t1, t2, t3
}

type MayInvoker4[T1 any, T2 any, T3 any, T4 any] struct {
	*mayHandlers
}

func NewMay4[T1 any, T2 any, T3 any, T4 any]() *MayInvoker4[T1, T2, T3, T4] {
	return &MayInvoker4[T1, T2, T3, T4]{
		mayHandlers: newMayHandlers(),
	}
}

func (f *MayInvoker4[T1, T2, T3, T4]) Use(handler ...MayHandler) *MayInvoker4[T1, T2, T3, T4] {
	f.mayHandlers.Use(handler...)
	return f
}

func (f *MayInvoker4[T1, T2, T3, T4]) Invoke(t1 T1, t2 T2, t3 T3, t4 T4, err any, messageArgs ...any) (T1, T2, T3, T4) {
	if err == nil {
		return t1, t2, t3, t4
	}

	f.mayHandlers.handleError(err, messageArgs...)

	return t1, t2, t3, t4
}

type MayInvoker5[T1 any, T2 any, T3 any, T4 any, T5 any] struct {
	*mayHandlers
}

func NewMay5[T1 any, T2 any, T3 any, T4 any, T5 any]() *MayInvoker5[T1, T2, T3, T4, T5] {
	return &MayInvoker5[T1, T2, T3, T4, T5]{
		mayHandlers: newMayHandlers(),
	}
}

func (f *MayInvoker5[T1, T2, T3, T4, T5]) Use(handler ...MayHandler) *MayInvoker5[T1, T2, T3, T4, T5] {
	f.mayHandlers.Use(handler...)
	return f
}

func (f *MayInvoker5[T1, T2, T3, T4, T5]) Invoke(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, err any, messageArgs ...any) (T1, T2, T3, T4, T5) {
	if err == nil {
		return t1, t2, t3, t4, t5
	}

	f.mayHandlers.handleError(err, messageArgs...)

	return t1, t2, t3, t4, t5
}

type MayInvoker6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any] struct {
	*mayHandlers
}

func NewMay6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any]() *MayInvoker6[T1, T2, T3, T4, T5, T6] {
	return &MayInvoker6[T1, T2, T3, T4, T5, T6]{
		mayHandlers: newMayHandlers(),
	}
}

func (f *MayInvoker6[T1, T2, T3, T4, T5, T6]) Use(handler ...MayHandler) *MayInvoker6[T1, T2, T3, T4, T5, T6] {
	f.mayHandlers.Use(handler...)
	return f
}

func (f *MayInvoker6[T1, T2, T3, T4, T5, T6]) Invoke(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, err any, messageArgs ...any) (T1, T2, T3, T4, T5, T6) {
	if err == nil {
		return t1, t2, t3, t4, t5, t6
	}

	f.mayHandlers.handleError(err, messageArgs...)

	return t1, t2, t3, t4, t5, t6
}
