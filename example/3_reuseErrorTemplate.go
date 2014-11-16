
package main

import (
	"fmt"
	errors ".."
)

var ErrInvalidUrl = errors.New()
// When using more then once, create a central method for filling the errors template
var newErrInvalidUrl = func(url string) (errors.Error) {
	return ErrInvalidUrl.ServerErrorf("invalid url [%s]", url)
}

func causeError(url string) (errors.Error) {
	return newErrInvalidUrl(url)
}
func getError(url string) (errors.Error) {
	return newErrInvalidUrl(url)
}

func main() {
	Err := causeError("http:://example.com/")
	fmt.Println(Err.Debug())
}

