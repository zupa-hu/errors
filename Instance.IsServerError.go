
package errors

// Returns true for an internal error (bug)
func (instance *Instance) IsServerError() (bool) {
	return instance.internal
}

