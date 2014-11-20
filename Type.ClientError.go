
package errors

// Create a client error instance from the given type.
// The message is intended for being passed back to an untrusted client.
func (t *Type) ClientError(clientNote string) (Error) {
	skip, internal, serverNote := 1, false, ""
	return t.newInstance(skip, internal, serverNote, clientNote)
}

