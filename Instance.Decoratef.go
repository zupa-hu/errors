
package errors

// Decoratef is to Decorate what fmt.Printf is to fmt.Print.
func (instance *Instance) Decoratef(template string, args ...interface{}) (Error) {
	return instance
}

