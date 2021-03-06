
package errors

import (
	"testing"
)

var testErrInstanceClientNote = Type("testErrInstanceClientNote")
func returnErrorInstanceClientNote() (*Instance) {
	skip, internal, serverNote, clientNote := 0, true, "", "client error"
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

func TestInstanceClientNote_addNoteToShorterStackTrace(t *testing.T) {
	instance := returnErrorInstanceClientNote()

	func() {
		func() {
			func() {
				instance.ClientNote("clientNote1")

				se := instance.stack[0]
				if se.clientNote != "client error\nclientNote1" { t.Fatal(se.clientNote) }
				if se.serverNote != "" { t.Fatal(se.serverNote) }
			}()
		}()
	}()
}

