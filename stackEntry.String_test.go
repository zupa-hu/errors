
package errors

import (
	"testing"
)

func TestStackEntryString(t *testing.T) {
	se := &stackEntry{
		file: "file",
		line: 123,
		funcName: "funcName",
		serverNote: "serverNote",
		clientNote: "clientNote",
	}
	expected := ""+
		"fn:  funcName\n"+
		"src: file:123\n"+
		"srv: serverNote\n"+
		"cli: clientNote\n"+
		"\n";
	actual := se.String()
	if actual != expected { t.Fatalf("\n[%v]\n[%v]", expected, actual) }


	// Omit clientError and serverError lines when empty
	se = &stackEntry{
		file: "file",
		line: 123,
		funcName: "funcName",
		serverNote: "",
		clientNote: "",
	}
	expected = ""+
		"fn:  funcName\n"+
		"src: file:123\n"+
		"\n";
	actual = se.String()
	if actual != expected { t.Fatalf("\n[%v]\n[%v]", expected, actual) }
}

