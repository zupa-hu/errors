
package errors

// Create a server error object from given type.
// The message is intended for internal debugging.
func (typ *Type) ServerError(serverNote string) (Error) {
	skip, internal, clientNote := 1, true, ""
	return typ.newInstance(skip, internal, serverNote, clientNote)
}

