
package errors

import (
	"testing"
)

var testErrInstanceServerNotef = Type("testErrInstanceServerNotef")
func returnErrorInstanceServerNotef() (*Instance) {
	skip, internal, serverNote, clientNote := 0, true, "server error", ""
	return testErrInstanceServerNotef.newInstance(skip, internal, serverNote, clientNote)
}

func TestInstanceServerNotef(t *testing.T) {
	instance := returnErrorInstanceServerNotef()

	se := instance.stack[1]
	if se.clientNote != "" { t.Fatal(se.clientNote) }
	if se.serverNote != "" { t.Fatal(se.serverNote) }

	instance.ServerNotef("format [%s]", "serverNote1")

	se = instance.stack[1]
	if se.clientNote != "" { t.Fatal(se.clientNote) }
	if se.serverNote != "format [serverNote1]" { t.Fatal(se.serverNote) }
}

func TestInstanceServerNotef_addNoteToShorterStackTrace(t *testing.T) {
	instance := returnErrorInstanceServerNotef()

	func() {
		func() {
			func() {
				instance.ServerNotef("format [%s]", "serverNote1")

				se := instance.stack[0]
				if se.clientNote != "" { t.Fatal(se.clientNote) }
				if se.serverNote != "server error\nformat [serverNote1]" { t.Fatal(se.serverNote) }
			}()
		}()
	}()
}

