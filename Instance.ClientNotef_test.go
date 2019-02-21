
package errors

import (
	"testing"
)

func returnErrorInstanceClientNotef() (*Instance) {
	skip, internal, serverNote, clientNote := 0, true, "", ""
	return ErrGeneric.newInstance(skip, internal, serverNote, clientNote)
}

func TestInstanceClientNotef(t *testing.T) {
	instance := returnErrorInstanceClientNotef()

	instance.ClientNotef("format [%s]", "clientNote1")
	if instance.clientNotes != "format [clientNote1]" {
		t.Fatal(instance.clientNotes)
	}

	instance.ClientNotef("number [%v]", 2)
	if instance.clientNotes != "format [clientNote1]\nnumber [2]" {
		t.Fatal(instance.clientNotes)
	}
}

