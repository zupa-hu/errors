
package errors

// When debugging is turned on, .Error() redirects to .Debug()
var Debug = false
// Localize internal error message
var InternalServerError = "internal server error"

// Format an error message to be returned to the client.
// Merge all clientNotes from top to bottom and return it.
// Should that be empty, for internal errors, return InternalServerError
// instead. That should only happen for server errors. For client errors,
// it is caller's responsibility to provide a non-empty error message.
func (instance *Instance) Error() (string) {
	if Debug {
		return instance.Debug()
	} else {
		return instance.error()
	}
}
func (instance *Instance) error() (string) {
	msg := ""

	n := len(instance.stack)
	var se stackEntry
	for i:=0; i<n; i++ {
		se = instance.stack[i]
		if se.clientNote != "" {
			msg += se.clientNote + "\n"
		}
	}

	if instance.internal && (msg == "") {
		msg = InternalServerError + "\n"
	}

	return msg
}

