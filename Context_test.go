
package errors

import (
	"testing"
)

func TestContext_ok(t *testing.T) {
	Err := InContext(func(e Context) {
		// NOP
	})
	if Err != nil { t.Fatal(Err) }
}
func TestContext_error(t *testing.T) {
	Err := InContext(func(e Context) {
		e.Client("woof")
	})
	AssertClientError(t, Err, "woof")
}
func TestContext_handle(t *testing.T) {
	Err := InContext(func(e Context) {
		defer e.Handle(func() {
			t.Fatal("should not be called if there is no error")
		})
	})
	if Err != nil { t.Fatal(Err) }
}
func TestContext_errorWithClientNote(t *testing.T) {
	Err := InContext(func(e Context) {
		defer e.Handle(func() {
			e.ClientNote("note")
		})

		e.Client("woof")
	})
	AssertClientError(t, Err, "woof\nnote")
}
func TestContext_errorWithClientNotef(t *testing.T) {
	Err := InContext(func(e Context) {
		defer e.Handle(func() {
			e.ClientNotef("note [%v]", "foo")
		})

		e.Client("woof")
	})
	AssertClientError(t, Err, "woof\nnote [foo]")
}
func TestContext_clientNoteWithoutError(t *testing.T) {
	Err := InContext(func(e Context) {
		e.ClientNote("note")
	})
	AssertServerError(t, Err, "unexpected client note on non-existent error")
	AssertServerError(t, Err, "note")
}
func TestContext_clientNotefWithoutError(t *testing.T) {
	Err := InContext(func(e Context) {
		e.ClientNotef("note [%v]", "foo")
	})
	AssertServerError(t, Err, "unexpected client note on non-existent error")
	AssertServerError(t, Err, "note [foo]")
}
func TestContext_serverNoteWithoutError(t *testing.T) {
	Err := InContext(func(e Context) {
		e.ServerNote("note")
	})
	AssertServerError(t, Err, "unexpected server note on non-existent error")
	AssertServerError(t, Err, "note")
}
func TestContext_serverNotefWithoutError(t *testing.T) {
	Err := InContext(func(e Context) {
		e.ServerNotef("note [%v]", "foo")
	})
	AssertServerError(t, Err, "unexpected server note on non-existent error")
	AssertServerError(t, Err, "note [foo]")
}
func TestContext_panic(t *testing.T) {
	Err := InContext(func(e Context) {
		panic(6)
	})
	if Err == nil { t.Fatal("expected an error") }
	AssertServerError(t, Err, "Uncaught error: 6")
}
func TestContext_nested(t *testing.T) {
	Err1 := InContext(func(e1 Context) {
		Err2 := InContext(func(e2 Context) {
			e2.Client("woof")
		})
		e1.Throw(Err2)
	})
	AssertClientError(t, Err1, "woof")
}

