
package errors

import (
	"testing"
)

func TestCatch_ok(t *testing.T) {
	Err := Catch(func() {
		// NOP
	})
	if Err != nil { t.Fatal(Err) }
}
func TestCatch_error(t *testing.T) {
	ErrTest := New()
	Err := Catch(func() {
		panic(ErrTest.ServerError("woof!"))
	})
	if Err == nil { t.Fatal("expected an error") }
	if ! Err.Is(ErrTest) { t.Fatal(Err) }
}
func TestCatch_panicNonError(t *testing.T) {
	Err := Catch(func() {
		panic(6)
	})
	if Err == nil { t.Fatal("expected an error") }
}

