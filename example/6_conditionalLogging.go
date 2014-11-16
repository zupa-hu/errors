
package main

import (
	"fmt"
	errors ".."
)

var ErrServer = errors.New()
func newErrServer() (errors.Error) {
	return ErrServer.ServerError("this is a bug")
}

func main() {
	Err := newErrServer()
	if Err != nil {
		if Err.IsServerError() {
			// Notify sys admin
			fmt.Println("TODO: notify sys admin")
		}
		fmt.Println(Err)
	}
}

