
package errors

import (
	"fmt"
)

// .ServerError() with an fmt.Sprintf() like interface.
func (typ *Type) ServerErrorf(template string, args ...interface{}) (Error) {
	skip, internal, clientNote := 1, true, ""
	serverNote := fmt.Sprintf(template, args...)
	return typ.newInstance(skip, internal, serverNote, clientNote)
}

