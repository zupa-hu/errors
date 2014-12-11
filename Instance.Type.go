
package errors

// Return the error instance's type.
// Useful in switch{} statements.
func (instance *Instance) Type() (*Type) {
	return instance.typ
}

