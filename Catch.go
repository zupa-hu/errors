
package errors

var ErrCatch = New()

func Catch(cb func()) (Err Error) {
	defer func() {
		err := recover()
		if err == nil { return }

		var ok bool
		Err, ok = err.(Error)
		if ok { return }
		
		Err = ErrCatch.ServerErrorf("Uncaught error: %v", err)
	}()

	cb()

	return nil
}

