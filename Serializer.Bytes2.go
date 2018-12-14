
package errors

import ()

func (Serializer) AddBytes2(b []byte, m []byte) ([]byte, Error) {
	length := len(m)
	if 65535 < length {
		return b, Serverf("input too long [%v]", length)
	}
	
	// NOTE: there was a zero-day bug here until 2017-09-10: the length was shifted '<<' instead of '>>'
	b = append(b, byte(uint16(length)>>8), byte(uint16(length)))
	b = append(b, m...)
	return b, nil
}
func (Serializer) EatBytes2(b []byte) ([]byte, []byte, Error) {
	if len(b) < 2 {
		return b, []byte{}, Serverf("input too short [%v]", len(b))
	}
	length := uint16(b[0])<<8 + uint16(b[1])
	if len(b) < int(2+length) {
		return b, []byte{}, Serverf("input too short [%v]", len(b))
	}
	return b[2+length:], b[2:2+length], nil
}

