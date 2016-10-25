package serial

import (
	"math"
)

type Data []byte


func (d Data) WriteBytes(ptr int, a []byte) int {
	for _, v := range a {
		ptr = d.WriteByte(ptr, v)
	}
	return ptr
}

// 1 bit boolean
func (d Data) WriteBoolean(ptr int, v bool) int {
	if v {
		return d.WriteByte(ptr, 1)
	}
	return d.WriteByte(ptr, 0)
}

// 8 bits integer
func (d Data) WriteInt8(ptr int, v int8) int {
	return d.WriteByte(ptr, byte(v))
}

func (d Data) WriteUInt8(ptr int, v uint8) int {
	return d.WriteByte(ptr, byte(v))
}

// 16 bits integer
func (d Data) WriteInt16(ptr int, v int16) int {
	d[ptr] = byte((v >> 8) & 0xFF); ptr++
	d[ptr] = byte((v >> 0) & 0xFF); ptr++
	return ptr
}

func (d Data) WriteUInt16(ptr int, v uint16) int {
	d[ptr] = byte((v >> 8) & 0xFF); ptr++
	d[ptr] = byte((v >> 0) & 0xFF); ptr++
	return ptr
}

// 32 bits integer
func (d Data) WriteInt32(ptr int, v int32) int {
	d[ptr] = byte((v >> 24) & 0xFF); ptr++
	d[ptr] = byte((v >> 16) & 0xFF); ptr++
	d[ptr] = byte((v >> 8) & 0xFF); ptr++
	d[ptr] = byte((v >> 0) & 0xFF); ptr++
	return ptr
}

func (d Data) WriteUInt32(ptr int, v uint32) int {
	d[ptr] = byte((v >> 24) & 0xFF); ptr++
	d[ptr] = byte((v >> 16) & 0xFF); ptr++
	d[ptr] = byte((v >> 8) & 0xFF); ptr++
	d[ptr] = byte((v >> 0) & 0xFF); ptr++
	return ptr
}

// 64 bits integer
func (d Data) WriteInt64(ptr int, v int64) int {
	d[ptr] = byte((v >> 56) & 0xFF); ptr++
	d[ptr] = byte((v >> 48) & 0xFF); ptr++
	d[ptr] = byte((v >> 40) & 0xFF); ptr++
	d[ptr] = byte((v >> 32) & 0xFF); ptr++
	d[ptr] = byte((v >> 24) & 0xFF); ptr++
	d[ptr] = byte((v >> 16) & 0xFF); ptr++
	d[ptr] = byte((v >> 8) & 0xFF); ptr++
	d[ptr] = byte((v >> 0) & 0xFF); ptr++
	return ptr
}

func (d Data) WriteUInt64(ptr int, v uint64) int {
	d[ptr] = byte((v >> 56) & 0xFF); ptr++
	d[ptr] = byte((v >> 48) & 0xFF); ptr++
	d[ptr] = byte((v >> 40) & 0xFF); ptr++
	d[ptr] = byte((v >> 32) & 0xFF); ptr++
	d[ptr] = byte((v >> 24) & 0xFF); ptr++
	d[ptr] = byte((v >> 16) & 0xFF); ptr++
	d[ptr] = byte((v >> 8) & 0xFF); ptr++
	d[ptr] = byte((v >> 0) & 0xFF); ptr++
	return ptr
}

// 32 bits floating point
func (d Data) WriteFloat32(ptr int, v float32) int {
	u := math.Float32bits(v)
	return d.WriteUInt32(ptr, u)
}

func (d Data) WriteFloat64(ptr int, v float64) int {
	u := math.Float64bits(v)
	return d.WriteUInt64(ptr, u)
}

// 64 bits floating point
func (d Data) WriteComplex64(ptr int, v complex64) int {
	ptr = d.WriteFloat32(ptr, real(v))
	return d.WriteFloat32(ptr, imag(v))
}

func (d Data) WriteComplex128(ptr int, v complex128) int {
	ptr = d.WriteFloat64(ptr, real(v))
	return d.WriteFloat64(ptr, imag(v))
}

// Aliases
func (d Data) WriteByte(ptr int, v byte) int {
	d[ptr] = v; ptr++
	return ptr
}

func (d Data) WriteRune(ptr int, v rune) int {
	return d.WriteInt32(ptr, int32(v))
}

func (d Data) WriteInt(ptr, v int) int {
	return d.WriteInt32(ptr, int32(v))
}

// Strings
// Written using length indicator (int) at ptr destination
func (d Data) WriteString(ptr int, v string) int {
	ptr = d.WriteUInt16(ptr, uint16(len(v)))
	return d.WriteBytes(ptr, ([]byte)(v))
}