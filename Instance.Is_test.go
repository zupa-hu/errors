
package errors

import (
	"testing"
)

var testErrTestInstanceIs = New()
var testErrTestInstanceIsNot = New()

func TestInstanceIs(t *testing.T) {
	instance := testErrTestInstanceIs.ClientError("")

	if ! instance.Is(testErrTestInstanceIs) { t.Fatal() }
	if instance.Is(testErrTestInstanceIsNot) { t.Fatal() }
}

