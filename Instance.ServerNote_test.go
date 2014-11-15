
package errors

import (
	"testing"
)

var testErrInstanceServerNote = New()
func returnErrorInstanceServerNote() (*Instance) {
	skip, internal, serverNote, clientNote := 0, true, "", ""
	return testErrInstanceServerNote.newInstance(skip, internal, serverNote, clientNote)
}

func TestInstanceServerNote(t *testing.T) {
	instance := returnErrorInstanceServerNote()

	se := instance.stack[1]
	if se.clientNote != "" { t.Fatal(se.clientNote) }
	if se.serverNote != "" { t.Fatal(se.serverNote) }

	instance.ServerNote("serverNote1")

	se = instance.stack[1]
	if se.clientNote != "" { t.Fatal(se.clientNote) }
	if se.serverNote != "serverNote1" { t.Fatal(se.serverNote) }
}

