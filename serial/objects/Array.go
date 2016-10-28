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
		dat := make([]bool, count)
		for i := 0; i < count; i++ {
			dat[i] = data.ReadBoolean(ptr); ptr++
		}
		return BooleanArray(name, dat)
	case serial.Int8:
		dat := make([]int8, count)
		for i := 0; i < count; i++ {
			dat[i] = data.ReadInt8(ptr); ptr++
		}
		return Int8Array(name, dat)
	case serial.Int16:
		dat := make([]int16, count)
		for i := 0; i < count; i++ {
			dat[i] = data.ReadInt16(ptr); ptr++
		}
		return Int16Array(name, dat)
	case serial.Int32:
		dat := make([]int32, count)
		for i := 0; i < count; i++ {
			dat[i] = data.ReadInt32(ptr); ptr++
		}
		return Int32Array(name, dat)
	case serial.Int64:
		dat := make([]int64, count)
		for i := 0; i < count; i++ {
			dat[i] = data.ReadInt64(ptr); ptr++
		}
		return Int64Array(name, dat)
	case serial.UInt8:
		dat := make([]uint8, count)
		for i := 0; i < count; i++ {
			dat[i] = data.ReadUInt8(ptr); ptr++
		}
		return UInt8Array(name, dat)
	case serial.UInt16:
		dat := make([]uint16, count)
		for i := 0; i < count; i++ {
			dat[i] = data.ReadUInt16(ptr); ptr++
		}
		return UInt16Array(name, dat)
	case serial.UInt32:
		dat := make([]uint32, count)
		for i := 0; i < count; i++ {
			dat[i] = data.ReadUInt32(ptr); ptr++
		}
		return UInt32Array(name, dat)
	case serial.UInt64:
		dat := make([]uint64, count)
		for i := 0; i < count; i++ {
			dat[i] = data.ReadUInt64(ptr); ptr++
		}
		return UInt64Array(name, dat)
	case serial.Float32:
		dat := make([]float32, count)
		for i := 0; i < count; i++ {
			dat[i] = data.ReadFloat32(ptr); ptr++
		}
		return Float32Array(name, dat)
	case serial.Float64:
		dat := make([]float64, count)
		for i := 0; i < count; i++ {
			dat[i] = data.ReadFloat64(ptr); ptr++
		}
		return Float64Array(name, dat)
	case serial.Complex64:
		dat := make([]complex64, count)
		for i := 0; i < count; i++ {
			dat[i] = data.ReadComplex64(ptr); ptr++
		}
		return Complex64Array(name, dat)
	case serial.Complex128:
		dat := make([]complex128, count)
		for i := 0; i < count; i++ {
			dat[i] = data.ReadComplex128(ptr); ptr++
		}
		return Complex128Array(name, dat)
	}

	panic("Unknown field type encountered while deserializing data stream")

	return nil
}