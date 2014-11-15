
package errors

import (
	"testing"
)

var testErrInstanceServerNotef = New()
func returnErrorInstanceServerNotef() (*Instance) {
	skip, internal, serverNote, clientNote := 0, true, "", ""
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

