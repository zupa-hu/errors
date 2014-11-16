
package main

import (
	"fmt"
	errors ".."
)

var ErrServer = errors.New()

func main() {
	Err := ErrServer.ServerError("inline server error")
	fmt.Println("-- client error message --")
	fmt.Println(Err)

	fmt.Println("-- server error message --")
	fmt.Println(Err.Debug())
}

