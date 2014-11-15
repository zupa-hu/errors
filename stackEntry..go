
package errors

type stackEntry struct {
	file string
	line int
	funcName string
	serverNote string
	clientNote string
}

func newStackEntry(file string, line int, funcName, serverNote, clientNote string) (stackEntry) {
	return stackEntry{
		file: file,
		line: line,
		funcName: funcName,
		serverNote: serverNote,
		clientNote: clientNote,
	}
}

