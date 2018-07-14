
package errors

import (
	"testing"
	"strings"
)

func TestAssertServerError_nil(t *testing.T) {
	catcher := &CatchFatal{}
	var Err Error

	AssertServerError(catcher, Err, "expErr")
	if ! catcher.caught {
		t.Fatal("expected an error to have been caught")
	}
	if ! strings.Contains(catcher.s, "expected an error") {
		t.Fatalf("caught error [%v]", catcher.s)
	}
}
func TestAssertServerError_ServerError(t *testing.T) {
	catcher := &CatchFatal{}
	Err := Server("server error")
	
	AssertServerError(catcher, Err, "server error")
	if catcher.caught {
		t.Fatal("unexpectedly caught an error")
	}
}
func TestAssertServerError_ClientError(t *testing.T) {
	catcher := &CatchFatal{}
	Err := Client("client error")
	
	AssertServerError(catcher, Err, "expErr")
	if ! catcher.caught {
		t.Fatal("expected an error to have been caught")
	}
	if ! strings.Contains(catcher.s, "client error") {
		t.Fatalf("caught error [%v]", catcher.s)
	}
}

