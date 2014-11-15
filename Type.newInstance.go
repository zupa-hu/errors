
package errors

import (
	"runtime"
)

// Create a new error instance the current stack recorded.
// Add a message at the topmost stack entry.
// This is an internal method, skip `skip` number of stack entries
// to ignore stack entries internal to this package.
func (t *Type) newInstance(skip int, internal bool, serverNote, clientNote string) (*Instance) {
	// Record program counters. This enforces a max stack depth of 256.
	// Not a big issue as it is quite big, plus at the time of this writing,
	// Golang itself imposes a lower limit, 32.
	pcs := make([]uintptr, 256)
	n := runtime.Callers(skip+2, pcs)

	// Record each stack entry in a stackEntry object
	stack := make([]stackEntry, n)
	var fn *runtime.Func
	var file, funcName string
	var line int
	var emptyNote = ""
	for i:=0; i<n; i++ {
		fn = runtime.FuncForPC(pcs[i])
		funcName = fn.Name()
		file, line = fn.FileLine(pcs[i])
		
		stack[i] = newStackEntry(file, line, funcName, emptyNote, emptyNote)
	}

	// Decorate the top stack entry with a message
	stack[0].serverNote = serverNote
	stack[0].clientNote = clientNote

	// Return an error instance
	instance := Instance{
		typ: t,
		internal: internal,
		stack: stack,
	}
	
	return &instance
}

