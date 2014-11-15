
package errors

import (
	"sync"
)

var nextId = 1
var idMutex sync.Mutex

func New() (*Type) {
	idMutex.Lock()
	defer idMutex.Unlock()

	id := nextId
	nextId++
	return &Type{ id:id }
}

