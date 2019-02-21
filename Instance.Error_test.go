
package errors

import (
	"testing"
)

func TestInstanceError(t *testing.T) {
	// Server error, no client notes
	instance := &Instance{
		internal: true,
		stack: []stackEntry{
			stackEntry{},
		},
	}
	actual := instance.Error()
	if actual != InternalServerError { t.Fatal(actual) }


	// Server error, client notes on instance
	instance = &Instance{
		internal: true,
		stack: []stackEntry{
			stackEntry{},
		},
		clientNotes: "client-notes",
	}
	actual = instance.Error()
	if actual != InternalServerError+"\nclient-notes" { t.Fatal(actual) }


	// Server error, client note exists on stack and on instance
	instance = &Instance{
		internal: true,
		stack: []stackEntry{
			stackEntry{ clientNote:"clientNote1" },
			stackEntry{ clientNote:"" },
			stackEntry{ clientNote:"clientNote3" },
			stackEntry{ clientNote:"" },
			stackEntry{ clientNote:"clientNote5" },
		},
		clientNotes: "clientNote-extra",
	}
	expected := ""+
		"internal server error\n"+
		"clientNote1\n"+
		"clientNote3\n"+
		"clientNote5\n"+
		"clientNote-extra";
	actual = instance.Error()
	if actual != expected { t.Fatalf("\n[%v]\n[%v]", expected, actual) }


	// Client error
	instance = &Instance{
		internal: false,
		stack: []stackEntry{
			stackEntry{ clientNote:"clientNote1" },
		},
	}
	actual = instance.Error()
	if actual != "clientNote1" { t.Fatal(actual) }
}

