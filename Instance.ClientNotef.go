
package errors

import (
	"fmt"
	"runtime"
)

// Like .ClientNote(), but with an fmt.Printf() like interface.
func (instance *Instance) ClientNotef(template string, args... interface{}) (Error) {
	skip, pcs := 2, make([]uintptr, MAX_STACK_SIZE)
	n := runtime.Callers(skip, pcs)

	pos := len(instance.stack) - n

	// In weird cases, the error may be generated on a concurrent stack, having a shorter stack trace.
	// Avoid panicing here, check the stack entry exists. If not, use the deep end of the stack.
	if (pos < 0) || (len(instance.stack) - 1 < pos) {
		// If the desired stack position does not exist, append info at the deep end of the stack.
		instance.stack[0].clientNote += "\n" + fmt.Sprintf(template, args...)
	} else {
		instance.stack[pos].clientNote = fmt.Sprintf(template, args...)
	}

	return instance
}

