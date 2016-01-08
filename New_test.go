
package errors

import (
	"testing"
)

func TestNew(t *testing.T) {
	nextId = 64727644
	typ := New()
	if nextId != 64727645 { t.Fatal(nextId) }
	if typ.id != 64727644 { t.Fatal(typ.id) }
}

