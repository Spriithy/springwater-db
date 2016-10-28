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
func (a *array) ValueOf() interface{} {
	switch a.dataType {
	case serial.Boolean:
		return a.data.ReadBooleanArray(0, a.count)
	case serial.Int8:
		return a.data.ReadInt8Array(0, a.count)
	case serial.Int16:
		return a.data.ReadInt16Array(0, a.count)
	case serial.Int32:
		return a.data.ReadInt32Array(0, a.count)
	case serial.Int64:
		return a.data.ReadInt64Array(0, a.count)
	case serial.UInt8:
		return a.data.ReadUInt8Array(0, a.count)
	case serial.UInt16:
		return a.data.ReadUInt16Array(0, a.count)
	case serial.UInt32:
		return a.data.ReadUInt32Array(0, a.count)
	case serial.UInt64:
		return a.data.ReadUInt64Array(0, a.count)
	case serial.Float32:
		return a.data.ReadFloat32Array(0, a.count)
	case serial.Float64:
		return a.data.ReadFloat64Array(0, a.count)
	case serial.Complex64:
		return a.data.ReadComplex64Array(0, a.count)
	case serial.Complex128:
		return a.data.ReadComplex128Array(0, a.count)
	}

	panic("Parsing unknown data-typed field!")

	return nil
}

func (a *array) SetName(name string) {
	a.nameLength = (uint16)(len(name))
	a.name = ([]byte)(name)
}

func (a *array) GetBytes(d serial.Data, ptr int) int {
	ptr = d.WriteByte(ptr, a.containerType)
	ptr = d.WriteUInt16(ptr, a.nameLength)
	ptr = d.WriteBytes(ptr, a.name)
	ptr = d.WriteByte(ptr, a.dataType)
	ptr = d.WriteInt32(ptr, (int32)(a.count))
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

func ArrayFromBytes(data serial.Data, offset int) *array {
	var ptr int = offset

	containerType := data.ReadByte(ptr); ptr++
	if containerType != serial.ArrayContainer {
		panic("Unexpected non-array wrapper in data stream: " +
			serial.ContainerName[containerType])
	}

	l, name := data.ReadString(ptr)
	ptr += l + 2
	println(name)

	dt := data.ReadByte(ptr); ptr++
	count := (int)(data.ReadUInt32(ptr)); ptr += 2

	switch dt {
	case serial.Boolean:
		dat := data.ReadBooleanArray(ptr, count)
		return BooleanArray(name, dat)
	case serial.Int8:
		dat := data.ReadInt8Array(ptr, count)
		return Int8Array(name, dat)
	case serial.Int16:
		dat := data.ReadInt16Array(ptr, count)
		return Int16Array(name, dat)
	case serial.Int32:
		dat := data.ReadInt32Array(ptr, count)
		return Int32Array(name, dat)
	case serial.Int64:
		dat := data.ReadInt64Array(ptr, count)
		return Int64Array(name, dat)
	case serial.UInt8:
		dat := data.ReadUInt8Array(ptr, count)
		return UInt8Array(name, dat)
	case serial.UInt16:
		dat := data.ReadUInt16Array(ptr, count)
		return UInt16Array(name, dat)
	case serial.UInt32:
		dat := data.ReadUInt32Array(ptr, count)
		return UInt32Array(name, dat)
	case serial.UInt64:
		dat := data.ReadUInt64Array(ptr, count)
		return UInt64Array(name, dat)
	case serial.Float32:
		dat := data.ReadFloat32Array(ptr, count)
		return Float32Array(name, dat)
	case serial.Float64:
		dat := data.ReadFloat64Array(ptr, count)
		return Float64Array(name, dat)
	case serial.Complex64:
		dat := data.ReadComplex64Array(ptr, count)
		return Complex64Array(name, dat)
	case serial.Complex128:
		dat := data.ReadComplex128Array(ptr, count)
		return Complex128Array(name, dat)
	}

	panic("Unknown field type encountered while deserializing data stream")

	return nil
}