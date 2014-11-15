
package errors

import (
	"fmt"
	"runtime"
)

// Like .ClientNote(), but with an fmt.Printf() like interface.
func (instance *Instance) ClientNotef(template string, args... interface{}) (Error) {
	skip, pcs := 2, make([]uintptr, 256)
	n := runtime.Callers(skip, pcs)

	pos := len(instance.stack) - n
	instance.stack[pos].clientNote = fmt.Sprintf(template, args...)

	return instance
}

