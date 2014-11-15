
package errors

// Returns true for an error caused by client input
func (instance *Instance) IsClientError() (bool) {
	return ! instance.internal
}

