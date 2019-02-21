
package errors

import ()

type Context struct {
	Data *ContextData
}
type ContextData struct {
	Err Error
}

type _Catch struct {}

var ErrContextUncaught = Type("ErrContextUncaught")
var ErrNoteOnNoError = Type("ErrNoteOnNoError")

func InContext(cb func(e Context)) (Err Error) {
	e := Context{
		Data: &ContextData{},
	}

	defer func() {
		err := recover()
		if err == nil { return }

		_, ok := err.(_Catch)
		if ok {
			Err = e.Data.Err
			return
		}

		caughtErr, ok := err.(Error)
		if ! ok {
			Err = ErrContextUncaught.ServerErrorf("Uncaught error: %v", err)
			return
		}

		Err = ErrContextUncaught.ServerErrorf("Uncaught error: [%v]", caughtErr.Debug())
	}()

	cb(e)

	return nil
}

func (e *Context) Throw(Err Error) {
	e.Data.Err = Err
	panic(_Catch{})
}
func (e *Context) Handle(cb func()) {
	if e.Data.Err != nil {
		cb()
	}
}

func (e *Context) ClientNote(msg string) {
	if e.Data.Err == nil {
		Err := ErrNoteOnNoError.ServerError("unexpected client note on non-existent error")
		Err.ClientNote(msg)
		e.Throw(Err)
		return
	}
	e.Data.Err = e.Data.Err.ClientNote(msg)
}
func (e *Context) ClientNotef(template string, args... interface{}) {
	if e.Data.Err == nil {
		Err := ErrNoteOnNoError.ServerError("unexpected client note on non-existent error")
		Err.ClientNotef(template, args...)
		e.Throw(Err)
		return
	}
	e.Data.Err = e.Data.Err.ClientNotef(template, args...)
}
func (e *Context) ServerNote(msg string) {
	if e.Data.Err == nil {
		Err := ErrNoteOnNoError.ServerError("unexpected server note on non-existent error")
		Err.ServerNote(msg)
		e.Throw(Err)
		return
	}
	e.Data.Err = e.Data.Err.ServerNote(msg)
}
func (e *Context) ServerNotef(template string, args... interface{}) {
	if e.Data.Err == nil {
		Err := ErrNoteOnNoError.ServerError("unexpected server note on non-existent error")
		Err.ServerNotef(template, args...)
		e.Throw(Err)
		return
	}
	e.Data.Err = e.Data.Err.ServerNotef(template, args...)
}

func (e *Context) Client(clientNote string) {
	e.Throw(ErrGeneric.ClientError(clientNote))
}
func (e *Context) Clientf(template string, args ...interface{}) {
	e.Throw(ErrGeneric.ClientErrorf(template, args...))
}
func (e *Context) Server(serverNote string) {
	e.Throw(ErrGeneric.ServerError(serverNote))
}
func (e *Context) Serverf(template string, args ...interface{}) {
	e.Throw(ErrGeneric.ServerErrorf(template, args...))
}

