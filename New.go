
package errors

import (
	"sync"
)

var nextId int32 = 1
var idMutex sync.Mutex

// Create a new error type
func New() (*Type) {
	idMutex.Lock()
	defer idMutex.Unlock()

	id := nextId
	nextId++
	return &Type{ id:id }
}

