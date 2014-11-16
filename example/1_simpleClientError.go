
package main

import (
	"fmt"
	errors ".."
)

var ErrClient = errors.New()

func main() {
	Err := ErrClient.ClientError("this is a simple inline client error")
	fmt.Println("-- client error message --")
	fmt.Println(Err)

	fmt.Println("-- server error message --")
	fmt.Println(Err.Debug())
}

