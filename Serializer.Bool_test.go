
package errors

import (
	"testing"
)

func TestBool(t *testing.T) {
	test := func(orig []byte, v bool, exp []byte) {
		origClone := make([]byte, len(orig))
		copy(origClone, orig)
		act := Serializer{}.AddBool(origClone, v)
		if string(act) != string(exp) {
			t.Fatalf("orig%v v[%v] exp%v act%v", orig, v, exp, act)
		}

		written := act[len(orig):]
		eatBytes := append(written, []byte("blah")...)
		remainder, actV, Err := Serializer{}.EatBool(eatBytes)
		if Err != nil {
			t.Fatalf("orig%v v[%v] exp%v - had error: %v", orig, v, exp, Err)
		}
		if actV != v {
			t.Fatalf("orig%v v[%v] exp%v actV[%v]", orig, v, exp, actV)
		}
		if string(remainder) != "blah" {
			t.Fatalf("orig%v v[%v] exp%v remainder%v", orig, v, exp, remainder)
		}
	}

	test(
		[]byte{},
		false,
		[]byte{0},
	)
	test(
		[]byte{},
		true,
		[]byte{1},
	)
	test(
		[]byte{1, 2},
		false,
		[]byte{1, 2, 0},
	)
	test(
		[]byte{1, 2},
		true,
		[]byte{1, 2, 1},
	)
}

func TestEatBool(t *testing.T) {
	b := []byte{}
	b, _, Err := Serializer{}.EatBool(b)
	if Err == nil { t.Error("Expected an error") }
	expErr := "internal server error"
	actErr := Err.Error()
	if expErr != actErr {
		t.Fatal(expErr, actErr)
	}
}

