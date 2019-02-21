
package errors

import ()

// Each error Instance has its own associated client and server notes, which add
// extra debug information either for untrusted clients or to help debugging.
// This method appends to the client notes of the current error Instance.
// Use it to inform the client about client input errors that may be relevant
// for debugging. This method comes handy when not all the information is
// available at the place where the error was thrown.
// This method MODIFIES the existing error instance.
func (instance *Instance) ClientNote(msg string) (Error) {
	if instance.clientNotes != "" {
		instance.clientNotes += "\n"
	}
	instance.clientNotes += msg

	return instance
}

