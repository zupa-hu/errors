
package errors

import (
	"runtime"
)

// Set the server note of the current stack level to the given message.
// Use it to add debug information. See .ClientNote() for more details.
// Warning: do not delegate calling .ServerNote().
func (instance *Instance) ServerNote(msg string) (Error) {
	skip, pcs := 2, make([]uintptr, MAX_STACK_SIZE)
	n := runtime.Callers(skip, pcs)

	pos := len(instance.stack) - n

	// In weird cases, the error may be generated on a concurrent stack, having a shorter stack trace.
	// Avoid panicing here, check the stack entry exists. If not, use the deep end of the stack.
	if (pos < 0) || (len(instance.stack) - 1 < pos) {
		// If the desired stack position does not exist, append info at the deep end of the stack.
		instance.stack[0].serverNote += "\n" + msg
	} else {
		instance.stack[pos].serverNote = msg
	}

	return instance
}

