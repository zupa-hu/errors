
package errors

import (
	"fmt"
	"runtime"
)

// Like .ServerNote(), but with an fmt.Printf() like interface.
func (instance *Instance) ServerNotef(template string, args... interface{}) (Error) {
	skip, pcs := 2, make([]uintptr, 256)
	n := runtime.Callers(skip, pcs)

	pos := len(instance.stack) - n
	instance.stack[pos].serverNote = fmt.Sprintf(template, args...)

	return instance
}

