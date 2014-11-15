
package errors

import (
	"testing"
)

var testErrInstanceClientNote = New()
func returnErrorInstanceClientNote() (*Instance) {
	skip, internal, serverNote, clientNote := 0, true, "", ""
	return testErrInstanceClientNote.newInstance(skip, internal, serverNote, clientNote)
}

func TestInstanceClientNote(t *testing.T) {
	instance := returnErrorInstanceClientNote()

	se := instance.stack[1]
	if se.clientNote != "" { t.Fatal(se.clientNote) }
	if se.serverNote != "" { t.Fatal(se.serverNote) }

	instance.ClientNote("clientNote1")

	se = instance.stack[1]
	if se.clientNote != "clientNote1" { t.Fatal(se.clientNote) }
	if se.serverNote != "" { t.Fatal(se.serverNote) }
}

