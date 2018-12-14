
package errors

import ()

func (Serializer) AddBool(b []byte, v bool) []byte {
	if v == false {
		b = append(b, 0)
	} else {
		b = append(b, 1)
	}
	return b
}
func (Serializer) EatBool(b []byte) ([]byte, bool, Error) {
	var v bool
	if len(b) < 1 {
		return b, false, Serverf("input too short [%v]", len(b))
	}
	if b[0] == 0 {
		v = false
	} else {
		v = true
	}
	return b[1:], v, nil
}

