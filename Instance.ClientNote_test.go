
package errors

import (
	"testing"
)

func returnErrorInstanceClientNote() (*Instance) {
	skip, internal, serverNote, clientNote := 0, true, "", ""
	return ErrGeneric.newInstance(skip, internal, serverNote, clientNote)
}

func TestInstanceClientNote(t *testing.T) {
	instance := returnErrorInstanceClientNote()

	instance.ClientNote("clientNote1")
	if instance.clientNotes != "clientNote1" {
		t.Fatal(instance.clientNotes)
	}

	instance.ClientNote("clientNote2")
	if instance.clientNotes != "clientNote1\nclientNote2" {
		t.Fatal(instance.clientNotes)
	}
}

