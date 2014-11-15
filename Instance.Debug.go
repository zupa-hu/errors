
package errors

// Returns the stack trace decorated with error messages.
// MUST not be passed back to an untrusted client.
func (instance *Instance) Debug() (string) {
	str := ""
	if instance.internal {
		str += "internal error:\n"+
			"\n"
	} else {
		str += "client error:\n"+
			"\n"
	}

	for _, se := range instance.stack {
		str += se.String()
	}

	return str
}

