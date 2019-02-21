
package errors

// An error instance records the stack trace upon creation and
// optionally adds further debug information.
// Error instances are flagged as client or server errors.
// A client error is caused by a client, and so any debug info needs
// to be sent back to the client, without a stack trace of course.
// A server error is a bug, and thus debug information MUST NOT be
// sent back to the client, as it may leak implementation details.
type Instance struct {
	// Error type
	typ Type
	// Server errors are internal errors, client errors are not
	internal bool
	// Recorded stack
	// The first entry is the oldest, last entry is deepest in the stack.
	stack []stackEntry
	// clientNotes holds debug info for clients which can be added at any level along the chain
	clientNotes string
	// serverNotes holds internal debug info which can be added at any level along the chain
	serverNotes string
}

// Shorthand to New().ClientError()
func Client(clientNote string) (Error) {
	return ErrGeneric.ClientError(clientNote)
}
// Shorthand to New().ClientErrorf()
func Clientf(template string, args ...interface{}) (Error) {
	return ErrGeneric.ClientErrorf(template, args...)
}
// Shorthand to New().ServerError()
func Server(serverNote string) (Error) {
	return ErrGeneric.ServerError(serverNote)
}
// Shorthand to New().ServerErrorf()
func Serverf(template string, args ...interface{}) (Error) {
	return ErrGeneric.ServerErrorf(template, args...)
}

