package objects

import "github.com/Spriithy/SerialBits/serial"

type field struct {
	containerType, dataType byte
	nameLength              uint16

	name                    []byte
	data                    serial.Data
}

// Exported Field methods

func (f *field) ValueOf() interface{} {
	switch f.dataType {
	case serial.Boolean:
		return f.data.ReadBoolean(0)
	case serial.Int8:
		return f.data.ReadInt8(0)
	case serial.Int16:
		return f.data.ReadInt16(0)
	case serial.Int32:
		return f.data.ReadInt32(0)
	case serial.Int64:
		return f.data.ReadInt64(0)
	case serial.UInt8:
		return f.data.ReadUInt8(0)
	case serial.UInt16:
		return f.data.ReadUInt16(0)
	case serial.UInt32:
		return f.data.ReadUInt32(0)
	case serial.UInt64:
		return f.data.ReadUInt64(0)
	case serial.Float32:
		return f.data.ReadFloat32(0)
	case serial.Float64:
		return f.data.ReadFloat64(0)
	case serial.Complex64:
		return f.data.ReadComplex64(0)
	case serial.Complex128:
		return f.data.ReadComplex128(0)
	}

	panic("Parsing unknown data-typed field!")

	return nil
}

func (f *field) SetName(name string) {
	f.nameLength = (uint16)(len(name))
	f.name = ([]byte)(name)
}

func (f *field) GetBytes(d serial.Data, ptr int) int {
	ptr = d.WriteByte(ptr, f.containerType)
	ptr = d.WriteUInt16(ptr, f.nameLength)
	ptr = d.WriteBytes(ptr, f.name)
	ptr = d.WriteByte(ptr, f.dataType)
	ptr = d.WriteBytes(ptr, f.data)
	return ptr
}

func FieldFromBytes(data serial.Data, offset int) *field {
	var ptr int = offset

	containerType := data.ReadByte(ptr); ptr++
	if containerType != serial.FieldContainer {
		panic("Unexpected non-field wrapper in data stream: " + serial.ContainerName[containerType])
	}

	l, name := data.ReadString(ptr)
	ptr += l + 2
	println("\t" + name)

	dt := data.ReadByte(ptr); ptr++

	switch dt {
	case serial.Boolean:
		val := data.ReadBoolean(ptr)
		return BooleanField(name, val)
	case serial.Int8:
		val := data.ReadInt8(ptr)
		return Int8Field(name, val)
	case serial.Int16:
		val := data.ReadInt16(ptr)
		return Int16Field(name, val)
	case serial.Int32:
		val := data.ReadInt32(ptr)
		return Int32Field(name, val)
	case serial.Int64:
		val := data.ReadInt64(ptr)
		return Int64Field(name, val)
	case serial.UInt8:
		val := data.ReadUInt8(ptr)
		return UInt8Field(name, val)
	case serial.UInt16:
		val := data.ReadUInt16(ptr)
		return UInt16Field(name, val)
	case serial.UInt32:
		val := data.ReadUInt32(ptr)
		return UInt32Field(name, val)
	case serial.UInt64:
		val := data.ReadUInt64(ptr)
		return UInt64Field(name, val)
	case serial.Float32:
		val := data.ReadFloat32(ptr)
		return Float32Field(name, val)
	case serial.Float64:
		val := data.ReadFloat64(ptr)
		return Float64Field(name, val)
	case serial.Complex64:
		val := data.ReadComplex64(ptr)
		return Complex64Field(name, val)
	case serial.Complex128:
		val := data.ReadComplex128(ptr)
		return Complex128Field(name, val)
	}

	panic("Unknown field type encountered while deserializing data stream")

	return nil
}

func (f *field) GetSize() int {
	return 4 + (int)(f.nameLength) + len(f.data) /* 4 = 2*(1) byte + 1*(2) uint16 + */
}

// General Purpose Fields "Constructors"

func BooleanField(name string, val bool) *field {
	f := new(field)
	f.containerType = serial.FieldContainer
	f.dataType = serial.Boolean
	f.SetName(name)
	f.data = make(serial.Data, serial.GetSize(serial.Boolean))
	f.data.WriteBoolean(0, val)
	return f
}

func Int8Field(name string, val int8) *field {
	f := new(field)
	f.containerType = serial.FieldContainer
	f.dataType = serial.Int8
	f.SetName(name)
	f.data = make(serial.Data, serial.GetSize(serial.Int8))
	f.data.WriteInt8(0, val)
	return f
}

func Int16Field(name string, val int16) *field {
	f := new(field)
	f.containerType = serial.FieldContainer
	f.dataType = serial.Int16
	f.SetName(name)
	f.data = make(serial.Data, serial.GetSize(serial.Int16))
	f.data.WriteInt16(0, val)
	return f
}

func Int32Field(name string, val int32) *field {
	f := new(field)
	f.containerType = serial.FieldContainer
	f.dataType = serial.Int32
	f.SetName(name)
	f.data = make(serial.Data, serial.GetSize(serial.Int32))
	f.data.WriteInt32(0, val)
	return f
}

func Int64Field(name string, val int64) *field {
	f := new(field)
	f.containerType = serial.FieldContainer
	f.dataType = serial.Int64
	f.SetName(name)
	f.data = make(serial.Data, serial.GetSize(serial.Int64))
	f.data.WriteInt64(0, val)
	return f
}

func UInt8Field(name string, val uint8) *field {
	f := new(field)
	f.containerType = serial.FieldContainer
	f.dataType = serial.UInt8
	f.SetName(name)
	f.data = make(serial.Data, serial.GetSize(serial.UInt8))
	f.data.WriteUInt8(0, val)
	return f
}

func UInt16Field(name string, val uint16) *field {
	f := new(field)
	f.containerType = serial.FieldContainer
	f.dataType = serial.UInt16
	f.SetName(name)
	f.data = make(serial.Data, serial.GetSize(serial.UInt16))
	f.data.WriteUInt16(0, val)
	return f
}

func UInt32Field(name string, val uint32) *field {
	f := new(field)
	f.containerType = serial.FieldContainer
	f.dataType = serial.UInt32
	f.SetName(name)
	f.data = make(serial.Data, serial.GetSize(serial.UInt32))
	f.data.WriteUInt32(0, val)
	return f
}

func UInt64Field(name string, val uint64) *field {
	f := new(field)
	f.containerType = serial.FieldContainer
	f.dataType = serial.UInt64
	f.SetName(name)
	f.data = make(serial.Data, serial.GetSize(serial.UInt64))
	f.data.WriteUInt64(0, val)
	return f
}

func Float32Field(name string, val float32) *field {
	f := new(field)
	f.containerType = serial.FieldContainer
	f.dataType = serial.Float32
	f.SetName(name)
	f.data = make(serial.Data, serial.GetSize(serial.Float32))
	f.data.WriteFloat32(0, val)
	return f
}

func Float64Field(name string, val float64) *field {
	f := new(field)
	f.containerType = serial.FieldContainer
	f.dataType = serial.Float64
	f.SetName(name)
	f.data = make(serial.Data, serial.GetSize(serial.Float64))
	f.data.WriteFloat64(0, val)
	return f
}

func Complex64Field(name string, val complex64) *field {
	f := new(field)
	f.containerType = serial.FieldContainer
	f.dataType = serial.Complex64
	f.SetName(name)
	f.data = make(serial.Data, serial.GetSize(serial.Complex64))
	f.data.WriteComplex64(0, val)
	return f
}

func Complex128Field(name string, val complex128) *field {
	f := new(field)
	f.containerType = serial.FieldContainer
	f.dataType = serial.Complex128
	f.SetName(name)
	f.data = make(serial.Data, serial.GetSize(serial.Complex128))
	f.data.WriteComplex128(0, val)
	return f
}
