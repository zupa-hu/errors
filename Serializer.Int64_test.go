
package errors

import (
	"testing"
)

func TestInt64(t *testing.T) {
	test := func(orig []byte, v int64, exp []byte) {
		origClone := make([]byte, len(orig))
		copy(origClone, orig)
		act := Serializer{}.AddInt64(origClone, v)
		if string(act) != string(exp) {
			t.Fatalf("orig%v v[%v] exp%v act%v", orig, v, exp, act)
		}

		written := act[len(orig):]
		eatBytes := append(written, []byte("blah")...)
		remainder, actV, Err := Serializer{}.EatInt64(eatBytes)
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
		0,
		[]byte{0, 0, 0, 0, 0, 0, 0, 0},
	)
	test(
		[]byte{},
		255,
		[]byte{0, 0, 0, 0, 0, 0, 0, 255},
	)
	test(
		[]byte{},
		-1,
		[]byte{255, 255, 255, 255, 255, 255, 255, 255},
	)
	test(
		[]byte{},
		0x0102030405060708,
		[]byte{1, 2, 3, 4, 5, 6, 7, 8},
	)
	test(
		[]byte{5, 6},
		0,
		[]byte{5, 6, 0, 0, 0, 0, 0, 0, 0, 0},
	)
	test(
		[]byte{5, 6},
		255,
		[]byte{5, 6, 0, 0, 0, 0, 0, 0, 0, 255},
	)
	test(
		[]byte{5, 6},
		-1,
		[]byte{5, 6, 255, 255, 255, 255, 255, 255, 255, 255},
	)
	test(
		[]byte{5, 6},
		0x0102030405060708,
		[]byte{5, 6, 1, 2, 3, 4, 5, 6, 7, 8},
	)
}

func TestEatInt64(t *testing.T) {
	_, _, Err := Serializer{}.EatInt64([]byte{})
	if Err == nil { t.Error("Expected an error") }
	expErr := "internal server error"
	actErr := Err.Error()
	if expErr != actErr {
		t.Fatal(expErr, actErr)
	}

	_, _, Err = Serializer{}.EatInt64([]byte{0, 0, 0, 0, 0, 0, 0})
	if Err == nil { t.Error("Expected an error") }
	expErr = "internal server error"
	actErr = Err.Error()
	if expErr != actErr {
		t.Fatal(expErr, actErr)
	}
}

