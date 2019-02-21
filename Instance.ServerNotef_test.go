
package errors

import (
	"testing"
)

func returnErrorInstanceServerNotef() (*Instance) {
	skip, internal, serverNote, clientNote := 0, true, "", ""
	return ErrGeneric.newInstance(skip, internal, serverNote, clientNote)
}

func TestInstanceServerNotef(t *testing.T) {
	instance := returnErrorInstanceServerNotef()

	instance.ServerNotef("format [%s]", "serverNote1")
	if instance.serverNotes != "format [serverNote1]" {
		t.Fatal(instance.serverNotes)
	}

	instance.ServerNotef("number [%v]", 2)
	if instance.serverNotes != "format [serverNote1]\nnumber [2]" {
		t.Fatal(instance.serverNotes)
	}
}

