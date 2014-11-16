
package main

import (
	"fmt"
	errors ".."
)

var ErrFirst = errors.New()
var ErrSecond = errors.New()

func getError() (errors.Error) {
	Err := ErrFirst.ClientError("ooups")
	return Err
}

func main() {
	Err := getError()
	if Err.Is(ErrFirst) {
		fmt.Println("ErrFirst")
	}
	if Err.Is(ErrSecond) {
		fmt.Println("ErrSecond")
	}
}

