# errors
Golang errors implementation with stack traces and client/server errors.

[![Build Status](https://travis-ci.org/zupa-hu/errors.svg?branch=master)](https://travis-ci.org/zupa-hu/errors)
[![Coverage Status](https://img.shields.io/coveralls/zupa-hu/errors.svg)](https://coveralls.io/r/zupa-hu/errors)
[![GoDoc](https://godoc.org/github.com/zupa-hu/errors?status.svg)](https://godoc.org/github.com/zupa-hu/errors)

## Goal, design principles
* implement the standard errors interface
* type safe, no magic, no side effects
* offer stack traces
* allow decorating errors when passing back along the call chain (add debug info where it is available)
* do not leak implementation details to untrusted clients
* allow type-checking of errors (package level error types in dependant packages)

## Experimental error handling mechanism

This package comes with an experimental error handling mechanism for Go that very interestingly seem to check
all the requirements boxes:

- Go v1.* compatible - use it today
- type safe
- enjoy the simplicity and selfdocumenting nature of returned errors
- no need to actually return the errors
- faster then returning errors
- feature parity with the Go v2 error handling draft

**Impossible. Or is it?**

The trick is that instead of returning errors, we can pass in error contexts.

### Example 1

**Before**

```go
// Function definition
func Foo(input string) (output string, Err errors.Error) {
	if input == "foo" {
		return "", errors.Client("you can't use the word foo")
	}

	return input, nil
}

// Function usage
output, Err := Foo("input")
if Err != nil {
	log.Fatal(Err)
}
...
```

**New approach**

Here is how you would use it in most of the places:

```go
// Function definition
func Foo(e E, input string) (output string) {
	if input == "foo" {
		e.Client("you can't use the word foo")
	}

	return input
}

// Function usage
output := Foo(e, "input")
```

I said most of the places because at the very top of the call chain, you have to create an error context and handle it like so:

```go
Err := errors.InContext(func(e E) {
	output := Foo(e, "input")
	...
})
if Err != nil {
	log.Fatal(Err)
}
```

Also, `E` is a type alias that you may want to set up 1x in all your packages for convenience.
It is a shorthand for `*errors.Context`. Of course, you can use that latter as well.

```go
type E = *errors.Context
```

**In the examples below, I will leave out the `errors.InContext()` wrapper because you won't see it in most code segments.**

## Example 2

Let's see an example with more function calls.

**Before**

```go
func ManyFoos(input string) (string, errors.Error)
	result := ""

	output, Err := Foo(input)
	if Err != nil {
		return "", Err
	}
	result += output

	output, Err = Foo(input)
	if Err != nil {
		return "", Err
	}
	result += output

	output, Err = Foo(input)
	if Err != nil {
		return "", Err
	}
	result += output

	output, Err = Foo(input)
	if Err != nil {
		return "", Err
	}
	result += output

	return result, nil
}

output, Err := Foo("input")
if Err != nil {
	log.Fatal(Err)
}
```

**New Approach**

```go
func ManyFoos(e E, input string) (string)
	result := ""

	result += Foo(e, input)
	result += Foo(e, input)
	result += Foo(e, input)
	result += Foo(e, input)

	return result
}

output := Foo(e, "input")
```

### Example 3 - chaining

**Before**

```go
func (s *Storage) Get(key string) (value string, Err errors.Error) {
	if key == "invalid" {
		reutrn "", errors.Clientf("invalid key [%v]", key)
	}
	return s[key], nil
}
func (s *Storage) Set(key, value string) (Err errors.Error) {
	_, exists := s[key]
	if exists {
		return errors.Clientg("key is already set [%v]", key)
	}
	s[key] = value
	return nil
}

s := NewStorage()
Err := s.Set("key", "value")
if Err != nil {
	log.Fatal(Err)
}
value, Err := s.Get("key")
if Err != nil {
	log.Fatal(Err)
}

fmt.Println(value)
```

**New Approach**

```go
func (s *Storage) Get(e E, key string) (value string) {
	if key == "invalid" {
		e.Clientf("invalid key [%v]", key)
	}
	return s[key]
}
func (s *Storage) Set(e E, key, value string) {
	_, exists := s[key]
	if exists {
		e.Clientg("key is already set [%v]", key)
	}
	s[key] = value
}

s := NewStorage()
value := s.Set(e, "key", "value").Get(e, "key")

fmt.Println(value)
```

### Enjoy the simplicity and selfdocumenting nature of returned errors

Because an error context is passed in, it clearly documents that the function call may result in an error.
That's equal to returning an error. If, as a caller, you want to handle the error at any level, just
create a new error context:

```go
var output string
Err := errors.InContext(func(e E) {
	output = Foo(e, "input")
})
// output, Err - both available here
```

Note that this approach is not uncontrolled as `panic()`. You can `panic()` anywhere, but with this approach,
you can only "throw" errors where an error context is available, which means some codepath is guaranteed
to be expecting it.

### Faster then returning errors

Checking `if Err != nil { ... }` everywhere has a cost. On my machine, it's 0.25ns per check.
Setting up an error context has a cost. On my machine, it's around 100ns per context.

This means if you use a context in a place where your app checks `if Err != nil { ... }` more then 400 times,
this is actually faster. That will most likely be the case in any real world program.

### Feature parity with the Go v2 error handling draft

The Go v2 error handling draft includes a `handle` keyword that defines a code block for error handling.

**Example Go v2 draft compatible code:**

```go
func oops() (string, error) {
	if true {
		return "", errors.Client("oops")
	}
	return "keep calm"
}

func Foo(input string) (output string, err error) {
    handle err {
        // do something
    }

    return check oops()
}

result := check Foo("input")
```

**This package**

```go
func oops(e E) (string) {
	if true {
		e.Client("oops")
	}
	return "keep calm"
}

func Foo(e E, input string) (output string) {
    defer e.Handle(func() {
        // e.Err is the error object - do something
    })

    return oops(e)
}

result := Foo(e, "input")
```

-- End of experimental error contexts. --

----

# Examples
## Example - simple client error
```go
var ErrInvalidPassword = errors.Type("ERR_INVALID_PASSWORD")

func ReturnError(pw string) (errors.Error) {
	return ErrInvalidPassword.ClientError("invalid password")
}
```

## Example - returning error message to client
```go
Err := ReturnError(password)
if Err != nil {
	fmt.Println(Err)
	// Prints "invalid password"
}
```

## Example - checking error agains type
```go
Err := ReturnError(password)
if Err.Is(ErrInvalidPassword) {
	// ...
}
```

## Example - client error with debug message
```go
var ErrPwTooShort = errors.Type("ErrPwTooShort")
const MIN_LEN_PW = 8

func PasswordTooShort(pw string) (errors.Error) {
	if len(pw) < MIN_LEN_PW {
		return ErrPwTooShort.ClientErrorf("password too short [%v], minimum [%v] chars required", len(pw), MIN_LEN_PW)
	}
	return nil
}
```

## Example - server error with debug message
```go
var ErrFooNotBar = errors.Type("ErrFooNotBar")

func ExpectBar(foo string) (errors.Error) {
	if foo != "bar" {
		return ErrFooNotBar.ServerErrorf("unexpected foo value [%v]", foo)
	}
	return nil
}
```

## Example - print debug info for trusted client
```go
Err := ExpectBar("oh")
if Err != nil {
	// Prints stack trace with error message
	fmt.Println(Err.Debug())
	// Err.Error() prints debug message for untrusted client "internal server error"
}
```
Example debug message:
```
client error:

fn:  main.main
src: /home/user/project/example/clientError.go:21
cli: errors message

fn:  runtime.main
src: /usr/local/go/src/pkg/runtime/proc.c:255

fn:  runtime.goexit
src: /usr/local/go/src/pkg/runtime/proc.c:1445
```

## Example - decorate error object with custom client/server messages
```go
var ErrDecorate := errors.Type("ErrDecorate")
func ReturnOriginalError() (errors.Error) {
	return ErrDecorate.ServerError("original error message")
}
func ReturnDecoratedError() (errors.Error) {
	Err := ReturnOriginalError()
	Err.ClientNote("custom error decoration for untrusted clients")
	Err.ServerNotef("error decoration for [%v]", "debugging")
	return Err
}
```
The above may return a debug message like:
```
server error:

fn:  main.ReturnOriginalError
src: /home/user/project/example/clientError.go:21
srv: original error message

fn:  main.ReturnDecoratedError
src: /home/user/project/example/clientError.go:24
cli: custom error decoration for untrusted clients
srv: error decoration for [debugging]

fn:  main.main
src: /home/user/project/example/clientError.go:41

fn:  runtime.main
src: /usr/local/go/src/pkg/runtime/proc.c:255

fn:  runtime.goexit
src: /usr/local/go/src/pkg/runtime/proc.c:1445
```

## Example - log internal errors
```go
Err := doSomething()
if Err != nil && Err.IsServerError() {
	// Log internal error, notify sys admin
}
```

## Example - temporarily turn client errors into debug messages
```go
errors.Debug = true
// From here on, Err.Error() redirects to Err.Debug()
// Warning: use for local debugging only
```

## Example - panic instead of returning errors, catch them
```go
var ErrPanic = errors.Type("ErrPanic")

func FailDoingSg() {
	panic(ErrPanic.ClientError("panicing!"))
}

func CatchingPanic() (errors.Error) {
	return Catch(func() {
		FailDoingSg()
	})
}
```


# Installation

To install according to your `$GOPATH`:

```console
$ go get github.com/zupa-hu/errors
```

Import the errors package

```go
import "github.com/zupa-hu/errors"
```

# Testing

run `go test`.