# fo - function calling utilities and controls

[![tag](https://img.shields.io/github/tag/nekomeowww/fo.svg)](https://github.com/nekomeowww/fo/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.20-%23007d9c)
[![GoDoc](https://godoc.org/github.com/nekomeowww/fo?status.svg)](https://pkg.go.dev/github.com/nekomeowww/fo)
![Build Status](https://github.com/nekomeowww/fo/actions/workflows/ci.yml/badge.svg)
[![Go report](https://goreportcard.com/badge/github.com/nekomeowww/fo)](https://goreportcard.com/report/github.com/nekomeowww/fo)
[![Contributors](https://img.shields.io/github/contributors/nekomeowww/fo)](https://github.com/nekomeowww/fo/graphs/contributors)

---

This project is inspired by [samber/lo](https://github.com/samber/lo) (A Lodash-style Go library based on Go 1.18+ Generics)

**Why this name?**

Just followed the naming convention of [samber/lo](https://github.com/samber/lo) with the **f** prefix, it stands for **function**.

## 🚀 Install

```sh
go get github.com/nekomeowww/fo@v1
```

This library is v1 and follows SemVer strictly.

No breaking changes will be made to exported APIs before v2.0.0.

## 💡 Usage

You can import `fo` using:

```go
import (
    "github.com/nekomeowww/fo"
)
```

Then use one of the helpers below:

```go
name := fo.May(func () (string, error) {
    return "John", nil
})

fmt.Println(name)
// John
```

Most of the time, the compiler will be able to infer the type so that you can call: `fo.May(...)`.

## 🤠 Spec

GoDoc: [https://godoc.org/github.com/nekomeowww/fo](https://godoc.org/github.com/nekomeowww/fo)

Global setters:

- [SetLogger](#setlogger)
- [SetHandlers](#sethandlers)

Function helpers:

- [Invoke](#invoke)
- [Invoke0 -> Invoke6](#invoke0-6)
- [InvokeWith](#invokewith)
- [InvokeWith0 -> InvokeWith6](#invokewith0-6)
- [InvokeWithTimeout](#invokewithtimeout)
- [InvokeWithTimeout0 -> InvokeWithTimeout6](#invokewithtimeout0-6)

Error handling:

- [May](#may)
- [May0 -> May6](#may0-6)
- [NewMay](#newmay)
- [NewMay0 -> NewMay6](#newmay0-6)

### SetLogger

Sets the logger for the package.

```go
fo.SetLogger(logrus.New())
```

### SetHandlers

Sets the handlers for the package.

```go
fo.SetHandlers(
    func (err error, v ...any) {
        fmt.Println(err, v...)
    },
    func (err error, v ...any) {
        fmt.Println(err, v...)
    },
)
```

### Invoke

Calls any functions with `context.Context` control supported and returns the result.

```go
ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

val, err := fo.Invoke(ctx, func() (string, error) {
    time.Sleep(2 * time.Second)
    return "John", nil
})
// val == ""
// err == context deadline exceeded
```

### Invoke{0->6}

Invoke\* has the same behavior as Invoke, but returns multiple values.

```go
func example0() (error) {}
func example1() (int, error) {}
func example2() (int, string, error) {}
func example3() (int, string, time.Date, error) {}
func example4() (int, string, time.Date, bool, error) {}
func example5() (int, string, time.Date, bool, float64, error) {}
func example6() (int, string, time.Date, bool, float64, byte, error) {}

ctx1, cancel1 := context.WithTimeout(context.Background(), 1*time.Second)
defer cancel1()

err1 := fo.Invoke0(ctx1, example0())

ctx2, cancel2 := context.WithTimeout(context.Background(), 1*time.Second)
defer cancel2()

val1, err1 := fo.Invoke1(ctx1, example1()) // alias to Invoke

ctx3, cancel3 := context.WithTimeout(context.Background(), 1*time.Second)
defer cancel3()

val1, val2, err1 := fo.Invoke2(ctx1, example2())

ctx4, cancel4 := context.WithTimeout(context.Background(), 1*time.Second)
defer cancel4()

val1, val2, val3, err1 := fo.Invoke3(ctx1, example3())

ctx5, cancel5 := context.WithTimeout(context.Background(), 1*time.Second)
defer cancel5()

val1, val2, val3, val4, err1 := fo.Invoke4(ctx1, example4())

ctx6, cancel6 := context.WithTimeout(context.Background(), 1*time.Second)
defer cancel6()

val1, val2, val3, val4, val5, err1 := fo.Invoke5(ctx1, example5())

ctx7, cancel7 := context.WithTimeout(context.Background(), 1*time.Second)
defer cancel7()

val1, val2, val3, val4, val5, val6, err1 := fo.Invoke6(ctx1, example6())
```

### InvokeWith

A short cut to `Invoke` with `context.Context` that wrapped internally.

```go
val, err := fo.InvokeWith(func() (string, error) {
    time.Sleep(2 * time.Second)
    return "John", nil
}, fo.WithContextTimeout(1*time.Second))
// val == ""
// err == context deadline exceeded
```

### InvokeWith{0->6}

InvokeWith\* has the same behavior as InvokeWith, but returns multiple values.

```go
func example0() (error) {}
func example1() (int, error) {}
func example2() (int, string, error) {}
func example3() (int, string, time.Date, error) {}
func example4() (int, string, time.Date, bool, error) {}
func example5() (int, string, time.Date, bool, float64, error) {}
func example6() (int, string, time.Date, bool, float64, byte, error) {}


err1 := fo.InvokeWith0(ctx1, example0(), fo.WithContextTimeout(1*time.Second))


val1, err1 := fo.InvokeWith1(ctx1, example1(), fo.WithContextTimeout(1*time.Second)) // alias to InvokeWith


val1, val2, err1 := fo.InvokeWith2(ctx1, example2(), fo.WithContextTimeout(1*time.Second))


val1, val2, val3, err1 := fo.InvokeWith3(ctx1, example3(), fo.WithContextTimeout(1*time.Second))


val1, val2, val3, val4, err1 := fo.InvokeWith4(ctx1, example4(), fo.WithContextTimeout(1*time.Second))


val1, val2, val3, val4, val5, err1 := fo.InvokeWith5(ctx1, example5(), fo.WithContextTimeout(1*time.Second))


val1, val2, val3, val4, val5, val6, err1 := fo.InvokeWith6(ctx1, example6(), fo.WithContextTimeout(1*time.Second))
```

### InvokeWithTimeout

A short cut to `Invoke` with `context.Context` that wrapped internally and defaults to set `fo.WithContextTimeout(...)`.

```go
val, err := fo.InvokeWithTimeout(func() (string, error) {
    time.Sleep(2 * time.Second)
    return "John", nil
}, 1*time.Second)
// val == ""
// err == context deadline exceeded
```

### InvokeWithTimeout{0->6}

InvokeWithTimeout\* has the same behavior as InvokeWithTimeout, but returns multiple values.

```go
func example0() (error) {}
func example1() (int, error) {}
func example2() (int, string, error) {}
func example3() (int, string, time.Date, error) {}
func example4() (int, string, time.Date, bool, error) {}
func example5() (int, string, time.Date, bool, float64, error) {}
func example6() (int, string, time.Date, bool, float64, byte, error) {}


err1 := fo.InvokeWithTimeout0(ctx1, example0(), 1*time.Second)


val1, err1 := fo.InvokeWithTimeout1(ctx1, example1(), 1*time.Second) // alias to InvokeWithTimeout


val1, val2, err1 := fo.InvokeWithTimeout2(ctx1, example2(), 1*time.Second)


val1, val2, val3, err1 := fo.InvokeWithTimeout3(ctx1, example3(), 1*time.Second)


val1, val2, val3, val4, err1 := fo.InvokeWithTimeout4(ctx1, example4(), 1*time.Second)


val1, val2, val3, val4, val5, err1 := fo.InvokeWithTimeout5(ctx1, example5(), 1*time.Second)


val1, val2, val3, val4, val5, val6, err1 := fo.InvokeWithTimeout6(ctx1, example6(), 1*time.Second)
```

```go

### May

Wraps a function call and filter out the error values and only returns with the result values.

```go
val := fo.May(time.Parse("2006-01-02", "2022-01-15"))
// 2022-01-15

val := fo.May(time.Parse("2006-01-02", "bad-value"))
// nil
```

### May{0->6}

May\* has the same behavior as May, but returns multiple values.

```go
func example0() (error)
func example1() (int, error)
func example2() (int, string, error)
func example3() (int, string, time.Date, error)
func example4() (int, string, time.Date, bool, error)
func example5() (int, string, time.Date, bool, float64, error)
func example6() (int, string, time.Date, bool, float64, byte, error)

fo.May0(example0())
val1 := fo.May1(example1())    // alias to May
val1, val2 := fo.May2(example2())
val1, val2, val3 := fo.May3(example3())
val1, val2, val3, val4 := fo.May4(example4())
val1, val2, val3, val4, val5 := fo.May5(example5())
val1, val2, val3, val4, val5, val6 := fo.May6(example6())
```

You can wrap functions like `func (...) (..., ok bool)` with `May*` and get the result values.

```go
// math.Signbit(float64) bool
fo.May0(math.Signbit(v))

// bytes.Cut([]byte,[]byte) ([]byte, []byte, bool)
before, after := fo.May2(bytes.Cut(s, sep))
```

You can give context to the panic message by adding some printf-like arguments.

```go
val, ok := any(someVar).(string)
fo.May1(val, ok, "someVar may be a string, got '%s'", val)

list := []int{0, 1, 2}
item := 5
fo.May0(lo.Contains[int](list, item), "'%s' may always contain '%s'", list, item)
...
```

### NewMay

Wraps a function call and filter out the error values and only returns with the result values,
behaves just like `May(...)` and `May\*(...)`, but with customizable handler and error collect
instead of using the package level internal handlers. Suitable for in-function and goroutine usage.

```go
may := fo.NewMay[time.Time]()

val := may.Invoke(time.Parse("2006-01-02", "2022-01-15"))
// 2022-01-15

val := may.Invoke(time.Parse("2006-01-02", "bad-value"))
// nil
```

You could use `may.Use(...)` to add handlers to the `May` instance for better error handling.

```go
may := fo.NewMay[time.Time]()
may.Use(func (err error, v ...any) {
    fmt.Printf("error: %s\n", err)
})

val := may.Invoke(time.Parse("2006-01-02", "bad-value"))
// error: parsing time "bad-value" as "2006-01-02": cannot parse "bad-value" as "2006"
```

`may.Use(...)` supports chained calls.

```go
may := fo.NewMay[time.Time]()
    .Use(func (err error, v ...any) {
        fmt.Printf("error from handler 1: %s\n", err)
    })
    .Use(func (err error, v ...any) {
        fmt.Printf("error from handler 2: %s\n", err)
    })
    // multiple handlers will be called in order

val := may.Invoke(time.Parse("2006-01-02", "bad-value"))
// error from handler 1: parsing time "bad-value" as "2006-01-02": cannot parse "bad-value" as "2006"
// error from handler 2: parsing time "bad-value" as "2006-01-02": cannot parse "bad-value" as "2006"
```

`fo` ships with some basic handler functions for common error handling.

```go
may := fo.NewMay[string]()
    .Use(fo.WithLoggerHandler(logger)) // log error with logger
    .Use(fo.WithLogFuncHandler(log.Printf)) // log error with log.Printf
```

### NewMay{0->6}

NewMay\* has the same behavior as NewMay, but returns multiple values.

```go
may := NewMay2[string, string]()

val1, val2 := may.Invoke(func () (string, bool, error) {
    return "John", true, nil
})
// val1 == "John"
// val2 == true
```

### CollectAsError

Get the collected errors as a single error value from the `May` instance.

```go
may := fo.NewMay[time.Time]()

val := may.Invoke(time.Parse("2006-01-02", "bad-value"))
val2 := may.Invoke(time.Parse("2006-01-02", "bad-value2"))

err := may.CollectAsError() // error
// this error has been combined with go.uber.org/multierr.Combine(...).
// You can use errors.Is(err, someErr) to check if the error contains some error.
// Or just use multierr.Errors(err) to get the slice of errors.
```

This function is often useful when you want to return the error value only when
all the invocations called / finished instead of checking the error value after
each invocation.

```go
may := NewMay[time.Time]()

val := may.Invoke(time.Parse("2006-01-02", "bad-value"))
val2 := may.Invoke(time.Parse("2006-01-02", "bad-value2"))

if err := may.CollectAsError(); err != nil {
    return err
}
```

#### CollectAsErrors

Get the collected errors as a errors slice value from the `May` instance.

```go
may := NewMay[time.Time]()

val := may.Invoke(time.Parse("2006-01-02", "bad-value"))
val2 := may.Invoke(time.Parse("2006-01-02", "bad-value2"))

errs := may.CollectAsErrors() // []error
fmt.Println(errs)
// []errors{
//     parsing time "bad-value" as "2006-01-02": cannot parse "bad-value" as "2006",
//     parsing time "bad-value2" as "2006-01-02": cannot parse "bad-value2" as "2006",
// }
```

#### HandleErrors

Sometimes you may find out you need a unique way to handle the errors from a `May` instance,
instead of using the injected handlers with `Use(...)`.

```go
may := NewMay[time.Time]()

val := may.Invoke(time.Parse("2006-01-02", "bad-value"))
val2 := may.Invoke(time.Parse("2006-01-02", "bad-value2"))

may.HandleErrors(func(errs []error) {
    fmt.Printf("error: %s\n", errs)
})
// error: [parsing time "bad-value" as "2006-01-02": cannot parse "bad-value" as "2006" parsing time "bad-value2" as "2006-01-02": cannot parse "bad-value2" as "2006"]
```

Or perhaps you could return the collected errors when `defer`ing the `may.HandleErrors(...)` call.

```go
func func1() (err error) {
    may := NewMay[time.Time]()

    defer may.HandleErrors(func(errs []error) {
        err = fmt.Errorf("error: %s\n", errs)
    })

    val := may.Invoke(time.Parse("2006-01-02", "bad-value"))
    val2 := may.Invoke(time.Parse("2006-01-02", "bad-value2"))

    return nil
}

func main() {
    err := func1()
    fmt.Println(err)
}
// error: [parsing time "bad-value" as "2006-01-02": cannot parse "bad-value" as "2006" parsing time "bad-value2" as "2006-01-02": cannot parse "bad-value2" as "2006"]
```

#### HandleErrorsWithReturn

`may.HandleErrorsWithReturn(...)` is similar to `may.HandleErrors(...)`, but it returns the handled errors.

```go
may := NewMay[time.Time]()

val := may.Invoke(time.Parse("2006-01-02", "bad-value"))
val2 := may.Invoke(time.Parse("2006-01-02", "bad-value2"))

err := may.HandleErrorsWithReturn(func(err []error) error {
    return fmt.Errorf("error: %s\n", err)
})

fmt.Println(val, val2, err)
// 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC error: [parsing time "bad-value" as "2006-01-02": cannot parse "bad-value" as "2006" parsing time "bad-value2" as "2006-01-02": cannot parse "bad-value2" as "2006"]
```

## TODOs

- [ ] implement more testable examples
- [ ] add playground link to README
- [ ] add benchmark tests
- [ ] maybe use sync.Pool for package level May instances

## 🤝 Contributing

- Fork the [project](https://github.com/nekomeowww/fo)
- Fix [open issues](https://github.com/nekomeowww/fo/issues) or request new features

Helper naming: helpers must be self explanatory and respect standards (other languages, libraries...). Feel free to suggest many names in your contributions.

## 📝 License

Copyright © 2023 [Neko Ayaka](https://github.com/nekomeowww).

This project is [MIT](./LICENSE) licensed.
