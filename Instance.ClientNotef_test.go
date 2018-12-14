
package errors

import (
	"testing"
)

var testErrInstanceClientNotef = Type("testErrInstanceClientNotef")
func returnErrorInstanceClientNotef() (*Instance) {
	skip, internal, serverNote, clientNote := 0, true, "", ""
	return testErrInstanceClientNotef.newInstance(skip, internal, serverNote, clientNote)
}

func TestInstanceClientNotef(t *testing.T) {
	instance := returnErrorInstanceClientNotef()

	se := instance.stack[1]
	if se.clientNote != "" { t.Fatal(se.clientNote) }
	if se.serverNote != "" { t.Fatal(se.serverNote) }

	instance.ClientNotef("format [%s]", "clientNote1")

	se = instance.stack[1]
	if se.clientNote != "format [clientNote1]" { t.Fatal(se.clientNote) }
	if se.serverNote != "" { t.Fatal(se.serverNote) }
}

