
package errors

import (
	"fmt"
)

// .ClientError() with an fmt.Sprintf() like interface.
func (t *Type) ClientErrorf(template string, args ...interface{}) (Error) {
	skip, internal, serverNote := 1, false, ""
	clientNote := fmt.Sprintf(template, args...)
	return t.newInstance(skip, internal, serverNote, clientNote)
}

