
package errors

import ()

func (Serializer) AddInt64(b []byte, n int64) []byte {
	nbyte := []byte{byte(n>>56), byte(n>>48), byte(n>>40), byte(n>>32), byte(n>>24), byte(n>>16), byte(n>>8), byte(n)}
	b = append(b, nbyte...)
	return b
}
func (Serializer) EatInt64(b []byte) ([]byte, int64, Error) {
	if len(b) < 8 {
		return b, 0, Serverf("input too short [%v]", len(b))
	}
	
	// value
	un := uint64(b[0])<<56 + uint64(b[1])<<48 + uint64(b[2])<<40 + uint64(b[3])<<32 + uint64(b[4])<<24 + uint64(b[5])<<16 + uint64(b[6])<<8 + uint64(b[7])
	n := int64(un)
	
	return b[8:], n, nil
}

