package objects

import "github.com/Spriithy/SerialBits/serial"

type array struct {
	containerType, dataType byte
	nameLength              uint16

	count                   int
	name                    []byte
	data                    serial.Data
}

// Exported Field methods
func (a *array) SetName(name string) {
	a.nameLength = (uint16)(len(name))
	a.name = ([]byte)(name)
}

func (a *array) GetBytes(d serial.Data, ptr int) int {
	ptr = d.WriteByte(ptr, a.containerType)
	ptr = d.WriteUInt16(ptr, a.nameLength)
	ptr = d.WriteBytes(ptr, a.name)
	ptr = d.WriteByte(ptr, a.dataType)
	ptr = d.WriteInt32(ptr, int32(a.count))
	ptr = d.WriteBytes(ptr, a.data)
	return ptr
}

func (f *array) GetSize() int {
	return 8 + (int)(f.nameLength) + len(f.data) /* 4 = 2*(1) byte + 1*(2) uint16 + */
}

func BooleanArray(name string, data []bool) *array {
	a := new(array)
	a.containerType = serial.ArrayContainer
	a.dataType = serial.Boolean
	a.SetName(name)
	a.count = len(data)
	a.data = make(serial.Data, a.count * serial.GetSize(a.dataType))
	a.data.WriteBooleanArray(0, data)
	return a
}

func Int8Array(name string, data []int8) *array {
	a := new(array)
	a.containerType = serial.ArrayContainer
	a.dataType = serial.Int8
	a.SetName(name)
	a.count = len(data)
	a.data = make(serial.Data, a.count * serial.GetSize(a.dataType))
	a.data.WriteInt8Array(0, data)
	return a
}

func Int16Array(name string, data []int16) *array {
	a := new(array)
	a.containerType = serial.ArrayContainer
	a.dataType = serial.Int16
	a.SetName(name)
	a.count = len(data)
	a.data = make(serial.Data, a.count * serial.GetSize(a.dataType))
	a.data.WriteInt16Array(0, data)
	return a
}

func Int32Array(name string, data []int32) *array {
	a := new(array)
	a.containerType = serial.ArrayContainer
	a.dataType = serial.Int32
	a.SetName(name)
	a.count = len(data)
	a.data = make(serial.Data, a.count * serial.GetSize(a.dataType))
	a.data.WriteInt32Array(0, data)
	return a
}

func Int64Array(name string, data []int64) *array {
	a := new(array)
	a.containerType = serial.ArrayContainer
	a.dataType = serial.Int32
	a.SetName(name)
	a.count = len(data)
	a.data = make(serial.Data, a.count * serial.GetSize(a.dataType))
	a.data.WriteInt64Array(0, data)
	return a
}

func UInt8Array(name string, data []uint8) *array {
	a := new(array)
	a.containerType = serial.ArrayContainer
	a.dataType = serial.UInt8
	a.SetName(name)
	a.count = len(data)
	a.data = make(serial.Data, a.count * serial.GetSize(a.dataType))
	a.data.WriteUInt8Array(0, data)
	return a
}

func UInt16Array(name string, data []uint16) *array {
	a := new(array)
	a.containerType = serial.ArrayContainer
	a.dataType = serial.UInt16
	a.SetName(name)
	a.count = len(data)
	a.data = make(serial.Data, a.count * serial.GetSize(a.dataType))
	a.data.WriteUInt16Array(0, data)
	return a
}

func UInt32Array(name string, data []uint32) *array {
	a := new(array)
	a.containerType = serial.ArrayContainer
	a.dataType = serial.UInt32
	a.SetName(name)
	a.count = len(data)
	a.data = make(serial.Data, a.count * serial.GetSize(a.dataType))
	a.data.WriteUInt32Array(0, data)
	return a
}

func UInt64Array(name string, data []uint64) *array {
	a := new(array)
	a.containerType = serial.ArrayContainer
	a.dataType = serial.UInt64
	a.SetName(name)
	a.count = len(data)
	a.data = make(serial.Data, a.count * serial.GetSize(a.dataType))
	a.data.WriteUInt64Array(0, data)
	return a
}

func Float32Array(name string, data []float32) *array {
	a := new(array)
	a.containerType = serial.ArrayContainer
	a.dataType = serial.Float32
	a.SetName(name)
	a.count = len(data)
	a.data = make(serial.Data, a.count * serial.GetSize(a.dataType))
	a.data.WriteFloat32Array(0, data)
	return a
}

func Float64Array(name string, data []float64) *array {
	a := new(array)
	a.containerType = serial.ArrayContainer
	a.dataType = serial.Float64
	a.SetName(name)
	a.count = len(data)
	a.data = make(serial.Data, a.count * serial.GetSize(a.dataType))
	a.data.WriteFloat64Array(0, data)
	return a
}

func Complex64Array(name string, data []complex64) *array {
	a := new(array)
	a.containerType = serial.ArrayContainer
	a.dataType = serial.Complex64
	a.SetName(name)
	a.count = len(data)
	a.data = make(serial.Data, a.count * serial.GetSize(a.dataType))
	a.data.WriteComplex64Array(0, data)
	return a
}

func Complex128Array(name string, data []complex128) *array {
	a := new(array)
	a.containerType = serial.ArrayContainer
	a.dataType = serial.Complex128
	a.SetName(name)
	a.count = len(data)
	a.data = make(serial.Data, a.count * serial.GetSize(a.dataType))
	a.data.WriteComplex128Array(0, data)
	return a
}

func ByteArray(name string, data []byte) *array {
	a := new(array)
	a.containerType = serial.ArrayContainer
	a.dataType = serial.Byte
	a.SetName(name)
	a.count = len(data)
	a.data = make(serial.Data, a.count * serial.GetSize(a.dataType))
	a.data.WriteByteArray(0, data)
	return a
}
func RuneArray(name string, data []rune) *array {
	a := new(array)
	a.containerType = serial.ArrayContainer
	a.dataType = serial.Rune
	a.SetName(name)
	a.count = len(data)
	a.data = make(serial.Data, a.count * serial.GetSize(a.dataType))
	a.data.WriteRuneArray(0, data)
	return a
}

