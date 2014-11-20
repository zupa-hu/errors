# errors
Golang errors implementation with stack traces and client/server errors.

[![Build Status](https://travis-ci.org/zupa-hu/errors.svg?branch=master)](https://travis-ci.org/zupa-hu/errors)
[![Coverage Status](https://img.shields.io/coveralls/zupa-hu/errors.svg)](https://coveralls.io/r/zupa-hu/errors)

### Goal, design principles
* implement the standard errors interface
* type safe, no magic, no side effects
* offer stack traces
* allow decorating errors when passing back along the call chain (add debug info where it is available)
* do not leak implementation details to untrusted clients
* allow type-checking of errors (package level error types in dependant packages)


# Examples
### Example - simple client error
```go
var ErrInvalidPassword = errors.New()

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
var ErrPwTooShort = errors.New()
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
var ErrFooNotBar = errors.New()

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
var ErrDecorate := errors.New()
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