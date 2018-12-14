
package errors

import ()

func (Serializer) AddBytes4(b []byte, m []byte) ([]byte, Error) {
	length := uint32(len(m))
	if length > uint32(16*MB) {
		return b, Serverf("input too long [%v]", length)
	}
	
	b = append(b, byte(length>>24), byte(length>>16), byte(length>>8), byte(length))
	b = append(b, m...)
	return b, nil
}
func (Serializer) EatBytes4(b []byte) ([]byte, []byte, Error) {
	if len(b) < 4 {
		return b, []byte{}, Serverf("input too short [%v]", len(b))
	}
	length := uint32(b[0])<<24 + uint32(b[1])<<16 + uint32(b[2])<<8 + uint32(b[3])
	if len(b) < int(4+length) {
		return b, []byte{}, Serverf("input too short [%v]", len(b))
	}
	return b[4+length:], b[4:4+length], nil
}

