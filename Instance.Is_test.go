
package errors

import (
	"testing"
)

var testErrTestInstanceIs = Type("ERROR_TYPE_FOO")
var testErrTestInstanceIsNot = Type("ERROR_TYPE_BAR")

func TestInstanceIs(t *testing.T) {
	instance := testErrTestInstanceIs.ClientError("")

	if ! instance.Is(testErrTestInstanceIs) { t.Fatal() }
	if instance.Is(testErrTestInstanceIsNot) { t.Fatal() }
}

