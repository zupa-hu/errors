
package errors

import (
	"runtime"
)

// Set the server note of the current stack level to the given message.
// Use it to add debug information. See .ClientNote() for more details.
// Warning: do not delegate calling .ServerNote().
func (instance *Instance) ServerNote(msg string) (Error) {
	skip, pcs := 2, make([]uintptr, 256)
	n := runtime.Callers(skip, pcs)

	pos := len(instance.stack) - n
	instance.stack[pos].serverNote = msg

	return instance
}

