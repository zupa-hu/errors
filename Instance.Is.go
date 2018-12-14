
package errors

// Check an error instance is of given type.
// Eg.: Err.Is(ErrCustom)
func (instance *Instance) Is(typ Type) (bool) {
	return instance.typ == typ
}

