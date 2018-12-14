
package errors

import ()

func (Serializer) AddInt32(b []byte, n int32) []byte {
	nbyte := []byte{byte(n>>24), byte(n>>16), byte(n>>8), byte(n)}
	b = append(b, nbyte...)
	return b
}
func (Serializer) EatInt32(b []byte) ([]byte, int32, Error) {
	if len(b) < 4 {
		return b, 0, Serverf("input too short [%v]", len(b))
	}
	
	// value
	un := uint32(b[0])<<24 + uint32(b[1])<<16 + uint32(b[2])<<8 + uint32(b[3])
	n := int32(un)
	
	return b[4:], n, nil
}

