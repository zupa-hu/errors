
package main

import (
	errors ".."
	"fmt"
	"net/http"
)

var ErrServer = errors.New()
// Note that this method returns a general error object
func newErrServer() (error) {
	Err :=  ErrServer.ServerError("this is a bug")
	return Err
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := newErrServer()
    fmt.Fprint(w, err)
}

func main() {
	// This webserver will return "internal server error"
	// unless the errors.Debug is set to true, in which case it
	// will print a stack trace with the message "this is a bug".
	// Uncomment the line below to try it out
	errors.Debug = true

    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
