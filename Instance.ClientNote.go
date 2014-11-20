
package errors

import (
	"runtime"
)

// Each stack level has its own associated client and server notes, which add
// extra debug information either for untrusted clients or to help debugging.
// This method sets the client note of the current stack level to the given message.
// Use it to inform the client about client input errors that may be relevant
// for debugging. This method comes handy when not all the information is
// available at the place where the error was thrown.
// This method MODIFIES the existing error instance.
// Warning: do not delegate calling .ClientNote(), use it like:
//   return Err.ClientNote("message")
func (instance *Instance) ClientNote(msg string) (Error) {
	skip, pcs := 2, make([]uintptr, 256)
	n := runtime.Callers(skip, pcs)

	pos := len(instance.stack) - n
	instance.stack[pos].clientNote = msg

	return instance
}

