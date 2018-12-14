
package errors

import (
	"testing"
)

func TestBytes4(t *testing.T) {
	test := func(orig []byte, v []byte, exp []byte) {
		origClone := make([]byte, len(orig))
		copy(origClone, orig)
		act, Err := Serializer{}.AddBytes4(origClone, v)
		if Err != nil {
			t.Fatalf("orig%v v%v exp%v - had error: %v", orig, v, exp, Err)
		}
		if string(act) != string(exp) {
			t.Fatalf("\norig%v \nv%v \nexp%v \nact%v", orig, v, exp, act)
		}

		written := act[len(orig):]
		eatBytes := append(written, []byte("blah")...)
		remainder, actV, Err := Serializer{}.EatBytes4(eatBytes)
		if Err != nil {
			t.Fatalf("orig%v v%v exp%v - had error: %v", orig, v, exp, Err)
		}
		if string(actV) != string(v) {
			t.Fatalf("v%v actV%v", v, actV)
		}
		if string(remainder) != "blah" {
			t.Fatalf("orig%v v%v exp%v remainder%v", orig, v, exp, remainder)
		}
	}

	test(
		[]byte{},
		[]byte{},
		[]byte{0, 0, 0, 0},
	)
	test(
		[]byte{},
		[]byte{1},
		[]byte{0, 0, 0, 1, 1},
	)
	test(
		[]byte{1},
		[]byte{},
		[]byte{1, 0, 0, 0, 0},
	)
	test(
		[]byte{1, 2},
		[]byte{3, 4},
		[]byte{1, 2, 0, 0, 0, 2, 3, 4},
	)
	test(
		[]byte{1, 2},
		[]byte("hello"),
		[]byte{1, 2, 0, 0, 0, 5, 'h', 'e', 'l', 'l', 'o'},
	)
}

func TestAddBytes4(t *testing.T) {
	// Check length
	m := make([]byte, 66051)
	b, Err := Serializer{}.AddBytes4([]byte{}, m)
	if Err != nil { t.Fatal(Err) }
	if b[0] != 0 { t.Error(b[0]) }
	if b[1] != 1 { t.Error(b[1]) }
	if b[2] != 2 { t.Error(b[2]) }
	if b[3] != 3 { t.Error(b[3]) }
	
	// Max length OK
	KB := 1024
	MB := 1024*KB
	input := make([]byte, 16*MB)
	_, Err = Serializer{}.AddBytes4([]byte{}, input)
	if Err != nil { t.Fatal(Err.Debug()) }

	// Too long
	input = make([]byte, 16*MB + 1)
	_, Err = Serializer{}.AddBytes4([]byte{}, input)
	if Err == nil { t.Error("Expected an error") }
	expErr := "internal server error"
	actErr := Err.Error()
	if expErr != actErr {
		t.Fatal(expErr, actErr)
	}
}
func TestEatBytes4(t *testing.T) {
	// Length bytes not present
	_, _, Err := Serializer{}.EatBytes4([]byte{})
	if Err == nil { t.Fatal("Expected an error") }
	expErr := "internal server error"
	actErr := Err.Error()
	if expErr != actErr {
		t.Fatal(expErr, actErr)
	}

	// Length bytes not present
	_, _, Err = Serializer{}.EatBytes4([]byte{0, 0, 0})
	if Err == nil { t.Fatal("Expected an error") }
	expErr = "internal server error"
	actErr = Err.Error()
	if expErr != actErr {
		t.Fatal(expErr, actErr)
	}

	// Data bytes not present
	_, _, Err = Serializer{}.EatBytes4([]byte{0, 0, 0, 1})
	if Err == nil { t.Fatal("Expected an error") }
	expErr = "internal server error"
	actErr = Err.Error()
	if expErr != actErr {
		t.Fatal(expErr, actErr)
	}
}

