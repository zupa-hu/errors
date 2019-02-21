
package errors

import ()

// ServerNote adds internal debug information to the current error Instance.
// Warning: do not delegate calling .ServerNote().
func (instance *Instance) ServerNote(msg string) (Error) {
	if instance.serverNotes != "" {
		instance.serverNotes += "\n"
	}
	instance.serverNotes += msg

	return instance
}

