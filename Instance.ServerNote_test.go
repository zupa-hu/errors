
package errors

import (
	"testing"
)

func returnErrorInstanceServerNote() (*Instance) {
	skip, internal, serverNote, clientNote := 0, true, "", ""
	return ErrGeneric.newInstance(skip, internal, serverNote, clientNote)
}

func TestInstanceServerNote(t *testing.T) {
	instance := returnErrorInstanceServerNote()

	instance.ServerNote("serverNote1")
	if instance.serverNotes != "serverNote1" {
		t.Fatal(instance.serverNotes)
	}

	instance.ServerNote("serverNote2")
	if instance.serverNotes != "serverNote1\nserverNote2" {
		t.Fatal(instance.serverNotes)
	}
}

