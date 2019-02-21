
package errors

import (
	"fmt"
)

// Like .ServerNote(), but with an fmt.Printf() like interface.
func (instance *Instance) ServerNotef(template string, args... interface{}) (Error) {
	if instance.serverNotes != "" {
		instance.serverNotes += "\n"
	}
	instance.serverNotes += fmt.Sprintf(template, args...)

	return instance
}

