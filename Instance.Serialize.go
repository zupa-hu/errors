
package errors

import ()

// Serialize serializes an error instance that can be later deserialized
// WITHIN THE SAME running program. Because error types use internal IDs
// that are non-deterministically assigned when initializing the program,
// deserializing an error instance will break the error type.
func (instance *Instance) Serialize() (b []byte, Err Error) {
	// typ
	b, Err = Serializer{}.AddString2(b, string(instance.typ))
	if Err != nil { return nil, Err }

	// internal
	b = Serializer{}.AddBool(b, instance.internal)

	// stack
	var stackBytes []byte
	for i:=0; i<len(instance.stack); i++ {
		stackEntry := instance.stack[i]

		var stackEntryBytes []byte

		// file string
		stackEntryBytes, Err = Serializer{}.AddString2(stackEntryBytes, stackEntry.file)
		if Err != nil { return nil, Err }
		// line int
		stackEntryBytes = Serializer{}.AddInt32(stackEntryBytes, int32(stackEntry.line))
		// funcName string
		stackEntryBytes, Err = Serializer{}.AddString2(stackEntryBytes, stackEntry.funcName)
		if Err != nil { return nil, Err }
		// serverNote string
		stackEntryBytes, Err = Serializer{}.AddString2(stackEntryBytes, stackEntry.serverNote)
		if Err != nil { return nil, Err }
		// clientNote string
		stackEntryBytes, Err = Serializer{}.AddString2(stackEntryBytes, stackEntry.clientNote)
		if Err != nil { return nil, Err }

		// The entire stack entry
		stackBytes, Err = Serializer{}.AddBytes4(stackBytes, stackEntryBytes)
		if Err != nil { return nil, Err }
	}

	// All stack entries
	b, Err = Serializer{}.AddBytes4(b, stackBytes)
	if Err != nil { return nil, Err }

	return b, nil
}

func Deserialize(b []byte) (instance *Instance, Err Error) {
	// typ
	b, typStr, Err := Serializer{}.EatString2(b)
	if Err != nil { return nil, Err }

	// internal
	b, internal, Err := Serializer{}.EatBool(b)
	if Err != nil { return nil, Err }

	// stack
	b, stackBytes, Err := Serializer{}.EatBytes4(b)
	if Err != nil { return nil, Err }
	if len(b) != 0 {
		return nil, Clientf("can't deserialize error, unexpected bytes found")
	}

	stack := []stackEntry{}
	for 0 < len(stackBytes) {
		var stackEntryBytes []byte
		stackBytes, stackEntryBytes, Err = Serializer{}.EatBytes4(stackBytes)
		if Err != nil { return nil, Err }

		// file string
		stackEntryBytes, file, Err := Serializer{}.EatString2(stackEntryBytes)
		if Err != nil { return nil, Err }
		// line int
		stackEntryBytes, line, Err := Serializer{}.EatInt32(stackEntryBytes)
		if Err != nil { return nil, Err }
		// funcName string
		stackEntryBytes, funcName, Err := Serializer{}.EatString2(stackEntryBytes)
		if Err != nil { return nil, Err }
		// serverNote string
		stackEntryBytes, serverNote, Err := Serializer{}.EatString2(stackEntryBytes)
		if Err != nil { return nil, Err }
		// clientNote string
		stackEntryBytes, clientNote, Err := Serializer{}.EatString2(stackEntryBytes)
		if Err != nil { return nil, Err }

		if len(stackEntryBytes) != 0 {
			return nil, Clientf("can't deserialize error, invalid stack entry found")
		}

		stack = append(stack, newStackEntry(file, int(line), funcName, serverNote, clientNote))
	}

	instance = &Instance{
		typ: Type(typStr),
		internal: internal,
		stack: stack,
	}

	return instance, nil
}

