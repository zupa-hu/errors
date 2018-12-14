
package errors

import (
	"testing"
	"fmt"
)

func TestBytes2(t *testing.T) {
	test := func(orig []byte, v []byte, exp []byte) {
		origClone := make([]byte, len(orig))
		copy(origClone, orig)
		act, Err := Serializer{}.AddBytes2(origClone, v)
		if Err != nil {
			t.Fatalf("orig%v v%v exp%v - had error: %v", orig, v, exp, Err)
		}
		if string(act) != string(exp) {
			t.Fatalf("\norig%v \nv%v \nexp%v \nact%v", orig, v, exp, act)
		}

		written := act[len(orig):]
		eatBytes := append(written, []byte("blah")...)
		remainder, actV, Err := Serializer{}.EatBytes2(eatBytes)
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
		[]byte{0, 0},
	)
	test(
		[]byte{},
		[]byte{1},
		[]byte{0, 1, 1},
	)
	test(
		[]byte{1},
		[]byte{},
		[]byte{1, 0, 0},
	)
	test(
		[]byte{1, 2},
		[]byte{3, 4},
		[]byte{1, 2, 0, 2, 3, 4},
	)
	test(
		[]byte{1, 2},
		[]byte("hello"),
		[]byte{1, 2, 0, 5, 'h', 'e', 'l', 'l', 'o'},
	)

	// 256 long
	test(
		[]byte{},
		[]byte{
			1, 2, 3, 4, 5, 6, 7, 8,   9, 10, 11, 12, 13, 14, 15, 16,   17, 18, 19, 20, 21, 22, 23, 24,   25, 26, 27, 28, 29, 30, 31, 32,
			1, 2, 3, 4, 5, 6, 7, 8,   9, 10, 11, 12, 13, 14, 15, 16,   17, 18, 19, 20, 21, 22, 23, 24,   25, 26, 27, 28, 29, 30, 31, 32,
			1, 2, 3, 4, 5, 6, 7, 8,   9, 10, 11, 12, 13, 14, 15, 16,   17, 18, 19, 20, 21, 22, 23, 24,   25, 26, 27, 28, 29, 30, 31, 32,
			1, 2, 3, 4, 5, 6, 7, 8,   9, 10, 11, 12, 13, 14, 15, 16,   17, 18, 19, 20, 21, 22, 23, 24,   25, 26, 27, 28, 29, 30, 31, 32,
			1, 2, 3, 4, 5, 6, 7, 8,   9, 10, 11, 12, 13, 14, 15, 16,   17, 18, 19, 20, 21, 22, 23, 24,   25, 26, 27, 28, 29, 30, 31, 32,
			1, 2, 3, 4, 5, 6, 7, 8,   9, 10, 11, 12, 13, 14, 15, 16,   17, 18, 19, 20, 21, 22, 23, 24,   25, 26, 27, 28, 29, 30, 31, 32,
			1, 2, 3, 4, 5, 6, 7, 8,   9, 10, 11, 12, 13, 14, 15, 16,   17, 18, 19, 20, 21, 22, 23, 24,   25, 26, 27, 28, 29, 30, 31, 32,
			1, 2, 3, 4, 5, 6, 7, 8,   9, 10, 11, 12, 13, 14, 15, 16,   17, 18, 19, 20, 21, 22, 23, 24,   25, 26, 27, 28, 29, 30, 31, 32,
		},
		[]byte{
			1, 0,
			1, 2, 3, 4, 5, 6, 7, 8,   9, 10, 11, 12, 13, 14, 15, 16,   17, 18, 19, 20, 21, 22, 23, 24,   25, 26, 27, 28, 29, 30, 31, 32,
			1, 2, 3, 4, 5, 6, 7, 8,   9, 10, 11, 12, 13, 14, 15, 16,   17, 18, 19, 20, 21, 22, 23, 24,   25, 26, 27, 28, 29, 30, 31, 32,
			1, 2, 3, 4, 5, 6, 7, 8,   9, 10, 11, 12, 13, 14, 15, 16,   17, 18, 19, 20, 21, 22, 23, 24,   25, 26, 27, 28, 29, 30, 31, 32,
			1, 2, 3, 4, 5, 6, 7, 8,   9, 10, 11, 12, 13, 14, 15, 16,   17, 18, 19, 20, 21, 22, 23, 24,   25, 26, 27, 28, 29, 30, 31, 32,
			1, 2, 3, 4, 5, 6, 7, 8,   9, 10, 11, 12, 13, 14, 15, 16,   17, 18, 19, 20, 21, 22, 23, 24,   25, 26, 27, 28, 29, 30, 31, 32,
			1, 2, 3, 4, 5, 6, 7, 8,   9, 10, 11, 12, 13, 14, 15, 16,   17, 18, 19, 20, 21, 22, 23, 24,   25, 26, 27, 28, 29, 30, 31, 32,
			1, 2, 3, 4, 5, 6, 7, 8,   9, 10, 11, 12, 13, 14, 15, 16,   17, 18, 19, 20, 21, 22, 23, 24,   25, 26, 27, 28, 29, 30, 31, 32,
			1, 2, 3, 4, 5, 6, 7, 8,   9, 10, 11, 12, 13, 14, 15, 16,   17, 18, 19, 20, 21, 22, 23, 24,   25, 26, 27, 28, 29, 30, 31, 32,
		},
	)

	// Max length
	maxLengthBytes := make([]byte, 65535, 65535)
	maxLengthBytes[65533] = 2
	maxLengthBytes[1] = 3
	exp := append([]byte{255, 255}, maxLengthBytes...)

	act, Err := Serializer{}.AddBytes2([]byte{}, maxLengthBytes)
	if Err != nil { t.Fatal(Err) }
	if len(act) != len(exp) {
		t.Fatalf("\nlen(exp)=%v \nlen(act)=%v", len(exp), len(act))
	}
	if string(act) != string(exp) {
		var errMsg string
		for i:=0; i<len(exp); i++ {
			if exp[i] != act[i] {
				t.Fatalf("\nexp%v \nact%v", exp, act)
				errMsg += fmt.Sprintf("byte mismatch: exp[%v]=%v act[%v]=%v\n", i, exp[i], i, act[i])
			}
		}
		t.Fatal("Fail:\n" + errMsg)
	}
}

func TestAddBytes2(t *testing.T) {
	tooLongBytes := make([]byte, 65535+1, 65535+1)
	_, Err := Serializer{}.AddBytes2([]byte{}, tooLongBytes)
	if Err == nil { t.Error("Expected an error") }
	expErr := "internal server error"
	actErr := Err.Error()
	if expErr != actErr {
		t.Fatal(expErr, actErr)
	}
}

func TestEatBytes2(t *testing.T) {
	// Length bytes not present
	_, _, Err := Serializer{}.EatBytes2([]byte{})
	if Err == nil { t.Error("Expected an error") }
	expErr := "internal server error"
	actErr := Err.Error()
	if expErr != actErr {
		t.Fatal(expErr, actErr)
	}

	// Length bytes not present
	_, _, Err = Serializer{}.EatBytes2([]byte{0})
	if Err == nil { t.Error("Expected an error") }
	expErr = "internal server error"
	actErr = Err.Error()
	if expErr != actErr {
		t.Fatal(expErr, actErr)
	}

	// Data bytes not present
	_, _, Err = Serializer{}.EatBytes2([]byte{0, 1})
	if Err == nil { t.Error("Expected an error") }
	expErr = "internal server error"
	actErr = Err.Error()
	if expErr != actErr {
		t.Fatal(expErr, actErr)
	}
}

