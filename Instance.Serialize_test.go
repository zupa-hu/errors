
package errors

import (
	"testing"
)

func TestInstanceSerialize(t *testing.T) {
	instance := &Instance{
		typ: &Type{ id: 123 },
		internal: true,
		stack: []stackEntry{
			newStackEntry("file1", 1, "funcName1", "serverNote1", "clientNote1"),
			newStackEntry("file2", 2, "funcName2", "serverNote2", "clientNote2"),
		},
	}

	exp := []byte{
		0, 0, 0, 123,
		1,
		0, 0, 0, 104,
			0, 0, 0, 48,
				0, 5, 'f', 'i', 'l', 'e', '1',
				0, 0, 0, 1,
				0, 9, 'f', 'u', 'n', 'c', 'N', 'a', 'm', 'e', '1',
				0, 11, 's', 'e', 'r', 'v', 'e', 'r', 'N', 'o', 't', 'e', '1',
				0, 11, 'c', 'l', 'i', 'e', 'n', 't', 'N', 'o', 't', 'e', '1',
			0, 0, 0, 48,
				0, 5, 'f', 'i', 'l', 'e', '2',
				0, 0, 0, 2,
				0, 9, 'f', 'u', 'n', 'c', 'N', 'a', 'm', 'e', '2',
				0, 11, 's', 'e', 'r', 'v', 'e', 'r', 'N', 'o', 't', 'e', '2',
				0, 11, 'c', 'l', 'i', 'e', 'n', 't', 'N', 'o', 't', 'e', '2',
	}

	// Serialize
	act, Err := instance.Serialize()
	if Err != nil { t.Fatal(Err) }
	if string(exp) != string(act) {
		t.Fatal(exp, act)
	}

	// Restore
	restoredInstance, Err := Deserialize(act)
	if Err != nil { t.Fatal(Err.Debug()) }

	// Re-serialize -- that's the easiest way to make sure we get
	// what we started with, as Instance.typ uses a pointer
	act2, Err := restoredInstance.Serialize()
	if Err != nil { t.Fatal(Err) }
	if string(exp) != string(act2) {
		t.Fatal(exp, act2)
	}
}

