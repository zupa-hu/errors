
package errors

import (
	"testing"
	"strings"
)

var testErrTypeNewInstance = Type("testErrTypeNewInstance")
func returnErrorInstance() (*Instance) {
	skip, internal := 0, true
	return testErrTypeNewInstance.newInstance(skip, internal, "serverNote", "clientNote")
}

func TestTypeNewInstance(t *testing.T) {
	instance := returnErrorInstance()

	if instance.internal != true { t.Fatal(instance.Debug()) }
	if instance.typ != testErrTypeNewInstance { t.Fatal(instance.Debug()) }
	n := len(instance.stack)
	if n < 3 { t.Fatal(instance.Debug()) }
	
	stackEntry0 := instance.stack[0]
	if ! strings.HasSuffix(stackEntry0.funcName, "errors.returnErrorInstance") { t.Fatal(instance.Debug()) }
	if stackEntry0.serverNote != "serverNote" { t.Fatal(instance.Debug()) }
	if stackEntry0.clientNote != "clientNote" { t.Fatal(instance.Debug()) }
	if stackEntry0.line != 12 { t.Fatal(instance.Debug()) }
	if ! strings.HasSuffix(stackEntry0.file, "Type.newInstance_test.go") { t.Fatal(instance.Debug()) }
	
	stackEntry1 := instance.stack[1]
	if ! strings.HasSuffix(stackEntry1.funcName, "errors.TestTypeNewInstance") { t.Fatal(instance.Debug()) }
	if stackEntry1.serverNote != "" { t.Fatal(instance.Debug()) }
	if stackEntry1.clientNote != "" { t.Fatal(instance.Debug()) }
	if stackEntry1.line != 16 { t.Fatal(instance.Debug()) }
	if ! strings.HasSuffix(stackEntry1.file, "Type.newInstance_test.go") { t.Fatal(instance.Debug()) }
}

