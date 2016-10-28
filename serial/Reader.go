package serial

import "math"

func (d Data) ReadBoolean(ptr int) bool {
	return d[ptr] != 0x00
}

func (d Data) ReadInt8(ptr int) int8 {
	return (int8)(d[ptr])
}

func (d Data) ReadUInt8(ptr int) uint8 {
	return (uint8)(d[ptr])
}

func (d Data) ReadInt16(ptr int) int16 {
	return (int16)(d[ptr] << 8) | (int16)(d[ptr + 1] << 0)
}

func (d Data) ReadUInt16(ptr int) uint16 {
	return (uint16)(d[ptr] << 8) | (uint16)(d[ptr + 1] << 0)
}

func (d Data) ReadInt32(ptr int) int32 {
	return (int32)(d[ptr]) << 24 |
		(int32)(d[ptr + 1]) << 16 |
		(int32)(d[ptr + 2]) << 8 |
		(int32)(d[ptr + 3]) << 0
}

func (d Data) ReadUInt32(ptr int) uint32 {
	return (uint32)(d[ptr]) << 24 |
		(uint32)(d[ptr + 1]) << 16 |
		(uint32)(d[ptr + 2]) << 8 |
		(uint32)(d[ptr + 3]) << 0
}

func (d Data) ReadInt64(ptr int) int64 {
	return (int64)(d[ptr]) << 56 |
		(int64)(d[ptr + 1]) << 48 |
		(int64)(d[ptr + 2]) << 40 |
		(int64)(d[ptr + 3]) << 32 |
		(int64)(d[ptr + 4]) << 24 |
		(int64)(d[ptr + 5]) << 16 |
		(int64)(d[ptr + 6]) << 8 |
		(int64)(d[ptr + 7]) << 0
}

func (d Data) ReadUInt64(ptr int) uint64 {
	return (uint64)(d[ptr]) << 56 |
		(uint64)(d[ptr + 1]) << 48 |
		(uint64)(d[ptr + 2]) << 40 |
		(uint64)(d[ptr + 3]) << 32 |
		(uint64)(d[ptr + 4]) << 24 |
		(uint64)(d[ptr + 5]) << 16 |
		(uint64)(d[ptr + 6]) << 8 |
		(uint64)(d[ptr + 7]) << 0
}

func (d Data) ReadFloat32(ptr int) float32 {
	return math.Float32frombits(d.ReadUInt32(ptr))
}

func (d Data) ReadFloat64(ptr int) float64 {
	return math.Float64frombits(d.ReadUInt64(ptr))
}

func (d Data) ReadComplex64(ptr int) complex64 {
	return complex(d.ReadFloat32(ptr), d.ReadFloat32(ptr + 4))
}

func (d Data) ReadComplex128(ptr int) complex128 {
	return complex(d.ReadFloat64(ptr), d.ReadFloat64(ptr + 8))
}

// Aliases
func (d Data) ReadByte(ptr int) byte {
	return d[ptr]
}

func (d Data) ReadRune(ptr int) rune {
	return d.ReadInt32(ptr)
}

// Strings
func (d Data) ReadString(ptr int) (int, string) {
	l := (int)(d.ReadUInt16(ptr))
	return l, (string)(d[ptr + 2:ptr + 2 + l])
}