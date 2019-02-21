
package errors

import(
	"strings"
)

// Returns the stack trace decorated with error messages.
// MUST not be passed back to an untrusted client.
func (instance *Instance) Debug() (string) {
	str := ""
	if instance.internal {
		str += "internal server error:\n"+
			"\n"
	} else {
		str += "client error:\n"+
			"\n"
	}

	for _, se := range instance.stack {
		str += se.String()
	}

	if instance.clientNotes != "" {
		str += instance.clientNotes + "\n"
	}
	if instance.serverNotes != "" {
		str += instance.serverNotes + "\n"
	}

	return strings.TrimRight(str, "\n")
}

