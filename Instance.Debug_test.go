
package errors

import (
	"testing"
)

func TestInstanceDebug(t *testing.T) {
	instance := &Instance{
		internal: true,
		stack: []stackEntry{
			newStackEntry("file1", 1, "funcName1", "serverNote1", "clientNote1"),
			newStackEntry("file2", 2, "funcName2", "serverNote2", "clientNote2"),
		},
		clientNotes: "client-notes",
		serverNotes: "server-notes",
	}
	expected := ""+
		"internal server error:\n"+
		"\n"+
		"fn:  funcName1\n"+
		"src: file1:1\n"+
		"srv: serverNote1\n"+
		"cli: clientNote1\n"+
		"\n"+
		"fn:  funcName2\n"+
		"src: file2:2\n"+
		"srv: serverNote2\n"+
		"cli: clientNote2\n"+
		"\n"+
		"client-notes\n"+
		"server-notes"+
		"";
	actual := instance.Debug()
	if actual != expected { t.Fatalf("\n[%v]\n[%v]", expected, actual) }
}

