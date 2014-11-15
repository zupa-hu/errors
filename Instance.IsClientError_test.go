
package errors

import (
	"testing"
)

func TestInstanceIsClientError(t *testing.T) {
	instance := &Instance{}

	instance.internal = false
	if ! instance.IsClientError() { t.Fatal(instance.IsClientError()) }
	
	instance.internal = true
	if instance.IsClientError() { t.Fatal(instance.IsClientError()) }
}

