
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
	typ *Type
	// Server errors are internal errors, client errors are not
	internal bool
	// Recorded stack
	// The first entry is the oldest, last entry is deepest in the stack.
	stack []stackEntry
}

