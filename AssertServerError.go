
package errors

import (
	"strings"
)

// AssertServerError is a test helper method to assert a given error Instance
func AssertServerError(t Fataler, Err Error, expErr string) {
	if Err == nil {
		t.Fatal(Clientf("expected an error").Debug())
		return
	}

	if ! Err.IsServerError() {
		t.Fatal(Err.Debug())
		return
	}

	actErr := Err.Debug()
	if ! strings.Contains(actErr, expErr) {
		t.Fatal(Serverf("\nexpErr[%v]\nactErr[%v]", expErr, actErr).Debug())
	}
}

