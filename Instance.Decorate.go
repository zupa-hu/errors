
package errors

// Extends the stack trace of the error object with a message at
// the current position in the call stack.
// This method MUST only be called from methods already on the call stack,
// as only the call stack position is examined to identify the stack entry
// which needs to be extended. Let's rephrase. When you get an Error,
// call Decorate directly, DO NOT delegate calling Error.Decorate().
// Typically, this means sg. like:
//   return Err.Decorate("some debug message")
func (instance *Instance) Decorate(msg string) (Error) {
	return instance
}

