
// Golang errors implementation with stack traces and client/server errors.
// 
// For the package overview, please go to https://github.com/zupa-hu/errors
package errors

type Error interface {
	Is(*Type) (bool)
	Type() (*Type)
	IsClientError() (bool)
	IsServerError() (bool)
	Debug() (string)
	Error() (string)
	ClientNote(msg string) (Error)
	ClientNotef(template string, args ...interface{}) (Error)
	ServerNote(msg string) (Error)
	ServerNotef(template string, args ...interface{}) (Error)
}

// Document/enforce that Error interface implements the built-in error interface
var _ = error(Error(&Instance{}))

