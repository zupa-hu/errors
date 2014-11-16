
package main

import (
	"fmt"
	errors ".."
)

var ErrDecorate = errors.New()

func causeError() (errors.Error) {
	Err := ErrDecorate.ClientError("original client errors message")
	return Err
}
func getError() (errors.Error) {
	Err := causeError()
	return Err.ClientNote("some debug info for client")
}

func main() {
	Err := getError()
	fmt.Println("-- client error message --")
	fmt.Println(Err)

	fmt.Println("-- server error message --")
	fmt.Println(Err.Debug())
}

