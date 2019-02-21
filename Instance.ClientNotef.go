
package errors

import (
	"fmt"
)

// Like .ClientNote(), but with an fmt.Printf() like interface.
func (instance *Instance) ClientNotef(template string, args... interface{}) (Error) {
	if instance.clientNotes != "" {
		instance.clientNotes += "\n"
	}
	instance.clientNotes += fmt.Sprintf(template, args...)

	return instance
}

