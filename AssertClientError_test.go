
package errors

import (
	"testing"
	"strings"
	"fmt"
)

type CatchFatal struct {
	caught bool
	s string
}
func (c *CatchFatal) Fatal(s ...interface{}) {
	c.caught = true
	c.s = fmt.Sprint(s...)
}

func TestAssertClientError_nil(t *testing.T) {
	catcher := &CatchFatal{}
	var Err Error

	AssertClientError(catcher, Err, "expErr")
	if ! catcher.caught {
		t.Fatal("expected an error to have been caught")
	}
	if ! strings.Contains(catcher.s, "expected an error") {
		t.Fatalf("caught error [%v]", catcher.s)
	}
}
func TestAssertClientError_ServerError(t *testing.T) {
	catcher := &CatchFatal{}
	Err := Server("server error")
	
	AssertClientError(catcher, Err, "expErr")
	if ! catcher.caught {
		t.Fatal("expected an error to have been caught")
	}
	if ! strings.Contains(catcher.s, "server error") {
		t.Fatalf("caught error [%v]", catcher.s)
	}
}
func TestAssertClientError_ClientError(t *testing.T) {
	catcher := &CatchFatal{}
	Err := Client("client error")
	
	AssertClientError(catcher, Err, "client error")
	if catcher.caught {
		t.Fatal("unexpectedly caught an error")
	}
}

