
package errors

import (
	"testing"
)


func TestString2(t *testing.T) {
	test := func(orig []byte, v string, exp []byte) {
		origClone := make([]byte, len(orig))
		copy(origClone, orig)
		act, Err := Serializer{}.AddString2(origClone, v)
		if Err != nil {
			t.Fatalf("orig%v v[%v] exp%v - had error: %v", orig, v, exp, Err)
		}
		if string(act) != string(exp) {
			t.Fatalf("\norig%v \nv[%v] \nexp%v \nact%v", orig, v, exp, act)
		}

		written := act[len(orig):]
		eatBytes := append(written, []byte("blah")...)
		remainder, actV, Err := Serializer{}.EatString2(eatBytes)
		if Err != nil {
			t.Fatalf("orig%v v[%v] exp%v - had error: %v", orig, v, exp, Err)
		}
		if string(actV) != string(v) {
			t.Fatalf("v[%v] actV%v", v, actV)
		}
		if string(remainder) != "blah" {
			t.Fatalf("orig%v v[%v] exp%v remainder%v", orig, v, exp, remainder)
		}
	}

	test(
		[]byte{},
		"",
		[]byte{0, 0},
	)
	test(
		[]byte{},
		"a",
		[]byte{0, 1, 'a'},
	)
	test(
		[]byte{1},
		"",
		[]byte{1, 0, 0},
	)
	test(
		[]byte{1, 2},
		"ab",
		[]byte{1, 2, 0, 2, 'a', 'b'},
	)
	test(
		[]byte{1, 2},
		"hello",
		[]byte{1, 2, 0, 5, 'h', 'e', 'l', 'l', 'o'},
	)
}

func TestAddString2(t *testing.T) {
	tooLongString := string(make([]byte, 65535+1, 65535+1))
	_, Err := Serializer{}.AddString2([]byte{}, tooLongString)
	if Err == nil { t.Error("Expected an error") }
	expErr := "internal server error"
	actErr := Err.Error()
	if expErr != actErr {
		t.Fatal(expErr, actErr)
	}
}
func TestEatString2(t *testing.T) {
	// Length bytes not present
	_, _, Err := Serializer{}.EatString2([]byte{1})
	if Err == nil { t.Error("Expected an error") }
	expErr := "internal server error"
	actErr := Err.Error()
	if expErr != actErr {
		t.Fatal(expErr, actErr)
	}

	// Length bytes not present
	_, _, Err = Serializer{}.EatString2([]byte{0, 1})
	if Err == nil { t.Error("Expected an error") }
	expErr = "internal server error"
	actErr = Err.Error()
	if expErr != actErr {
		t.Fatal(expErr, actErr)
	}

	// Data bytes not present
	_, _, Err = Serializer{}.EatString2([]byte{0, 2, 1})
	if Err == nil { t.Error("Expected an error") }
	expErr = "internal server error"
	actErr = Err.Error()
	if expErr != actErr {
		t.Fatal(expErr, actErr)
	}
}

