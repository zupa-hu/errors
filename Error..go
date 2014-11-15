
package errors

type Error interface {
	Is(*Type) (bool)
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

