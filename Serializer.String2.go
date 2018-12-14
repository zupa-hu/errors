
package errors

import ()

func (ser Serializer) AddString2(b []byte, s string) ([]byte, Error) {
	return ser.AddBytes2(b, []byte(s))
}
func (ser Serializer) EatString2(b []byte) ([]byte, string, Error) {
	b, b2, err := ser.EatBytes2(b)
	return b, string(b2), err
}

