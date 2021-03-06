# errors
Golang errors implementation with stack traces and client/server errors.

[![Build Status](https://travis-ci.org/zupa-hu/errors.svg?branch=master)](https://travis-ci.org/zupa-hu/errors)
[![Coverage Status](https://img.shields.io/coveralls/zupa-hu/errors.svg)](https://coveralls.io/r/zupa-hu/errors)
[![GoDoc](https://godoc.org/github.com/zupa-hu/errors?status.svg)](https://godoc.org/github.com/zupa-hu/errors)

### Goal, design principles
* implement the standard errors interface
* type safe, no magic, no side effects
* offer stack traces
* allow decorating errors when passing back along the call chain (add debug info where it is available)
* do not leak implementation details to untrusted clients
* allow type-checking of errors (package level error types in dependant packages)

### Experimental
Checking for errors all the time is pain in the ass. Instead of returning them all the time as the
last argument, call `panic(Err)` instead and catch them with `Catch()`. There is a small performance penalty
(~0.2us) of using Catch, but when using at the top of a call-chain, it is offset by the reduced overhead of not
needing to check `if Err != nil { .. }` all the time. Run the benchmarks on your own machine, on my system
the speed of the 2 ways equaled around 100 function calls. So, if there are more than 100 func calls inside
the call of Catch, it starts to improve speed instead of degrading it. In the order of above 1000 calls,
function calling overhead dropped by 30%.

Note that using `panic()` is not idiomatic go. Yet if I can write 50% less code, that is easier to read, and
doing it is overall way more enjoyable, I'm willing to explore new grounds. The rules of what is idiomatic
tend to change over time anyway.


# Examples
### Example - simple client error
```go
var ErrInvalidPassword = errors.Type("ERR_INVALID_PASSWORD")

func ReturnError(pw string) (errors.Error) {
	return ErrInvalidPassword.ClientError("invalid password")
}
```

### Example - returning error message to client
```go
Err := ReturnError(password)
if Err != nil {
	fmt.Println(Err)
	// Prints "invalid password"
}
```

### Example - checking error agains type
```go
Err := ReturnError(password)
if Err.Is(ErrInvalidPassword) {
	// ...
}
```

### Example - client error with debug message
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

### Example - server error with debug message
```go
var ErrFooNotBar = errors.Type("ErrFooNotBar")

func ExpectBar(foo string) (errors.Error) {
	if foo != "bar" {
		return ErrFooNotBar.ServerErrorf("unexpected foo value [%v]", foo)
	}
	return nil
}
```

### Example - print debug info for trusted client
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

### Example - decorate error object with custom client/server messages
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

### Example - log internal errors
```go
Err := doSomething()
if Err != nil && Err.IsServerError() {
	// Log internal error, notify sys admin
}
```

### Example - temporarily turn client errors into debug messages
```go
errors.Debug = true
// From here on, Err.Error() redirects to Err.Debug()
// Warning: use for local debugging only
```

### Example - panic instead of returning errors, catch them
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