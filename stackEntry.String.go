
package errors

import(
	"strconv"
)

func (se *stackEntry) String() (string) {
	s := "fn:  "+se.funcName+"\n"
	s += "src: "+se.file+":"+strconv.Itoa(se.line)+"\n"

	if se.serverNote != "" {
		s += "srv: "+se.serverNote+"\n"
	}
	if se.clientNote != "" {
		s += "cli: "+se.clientNote+"\n"
	}
	
	s += "\n"

	return s
}

