package objects

import "go/types"

const (
	Unknown byte = iota

	Boolean

	Int8
	Int16
	Int32
	Int64

	UInt8
	UInt16
	UInt32
	UInt64

	Float32
	Float64

	Complex64
	Complex128

	Byte = UInt8
	Rune = Int32
)

var typeSize = map[byte]byte{
	Unknown:    0,

	Boolean:    1,

	Int8:       1,
	Int16:      2,
	Int32:      4,
	Int64:      8,

	UInt8:      1,
	UInt16:     2,
	UInt32:     4,
	UInt64:     8,

	Float32:    4,
	Float64:    8,

	Complex64:  8,
	Complex128: 16,
}

func GetSize(t byte) int {
	types.Byte
	return typeSize[t]
}