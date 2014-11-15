
package errors

import (
	"testing"
)

func TestInstanceIsServerError(t *testing.T) {
	instance := &Instance{}

	instance.internal = false
	if instance.IsServerError() { t.Fatal(instance.IsServerError()) }
	
	instance.internal = true
	if ! instance.IsServerError() { t.Fatal(instance.IsServerError()) }
}

