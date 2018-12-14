
package errors

import ()

func (ser Serializer) AddString4(b []byte, s string) ([]byte, Error) {
	return ser.AddBytes4(b, []byte(s))
}
func (ser Serializer) EatString4(b []byte) ([]byte, string, Error) {
	b, b2, err := ser.EatBytes4(b)
	return b, string(b2), err
}

