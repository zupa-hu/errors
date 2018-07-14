
package errors

// AssertClientError is a test helper method to assert a given error Instance
func AssertClientError(t Fataler, Err Error, expErr string) {
	if Err == nil {
		t.Fatal(Clientf("expected an error").Debug())
		return
	}

	if ! Err.IsClientError() {
		t.Fatal(Err.Debug())
		return
	}

	actErr := Err.Error()
	if expErr != actErr {
		t.Fatal(Serverf("\nexpErr[%v]\nactErr[%v]", expErr, actErr).Debug())
	}
}

