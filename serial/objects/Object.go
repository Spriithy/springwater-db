package objects

import (
	"github.com/Spriithy/SerialBits/serial"
)

type object struct {
	containerType byte
	nameLength    uint16
	name          []byte

	size          int

	// TODO Implement hashtable for object offsets in byte array
	fields        []*field
	arrays        []*array
	bstrings      []*bytestring
	rstrings      []*runestring
}

func SerialObject(name string) *object {
	obj := new(object)
	obj.containerType = serial.ObjectContainer
	obj.SetName(name)
	return obj
}

func (obj *object) SetName(name string) {
	if obj.size != 0 {
		obj.size -= (int)(obj.nameLength)
	}

	obj.nameLength = (uint16)(len(name))
	obj.name = ([]byte)(name)
	obj.size += (int)(obj.nameLength)
}

func (obj *object) AddField(f *field) {
	obj.fields = append(obj.fields, f)
	obj.size += f.GetSize()
}

func (obj *object) AddArray(a *array) {
	obj.arrays = append(obj.arrays, a)
	obj.size += a.GetSize()
}

func (obj *object) AddByteString(str *bytestring) {
	obj.bstrings = append(obj.bstrings, str)
	obj.size += str.GetSize()
}

func (obj *object) AddRuneString(str *runestring) {
	obj.rstrings = append(obj.rstrings, str)
	obj.size += str.GetSize()
}

func (obj *object) GetSize() int {
	return obj.size + 4
}

func (obj *object) GetBytes(d serial.Data, ptr int) int {
	ptr = d.WriteByte(ptr, obj.containerType)
	ptr = d.WriteUInt16(ptr, obj.nameLength)
	ptr = d.WriteBytes(ptr, obj.name)

	ptr = d.WriteUInt16(ptr, uint16(len(obj.fields)))
	for _, f := range obj.fields {
		ptr = f.GetBytes(d, ptr)
	}

	ptr = d.WriteUInt16(ptr, uint16(len(obj.arrays)))
	for _, a := range obj.arrays {
		ptr = a.GetBytes(d, ptr)
	}

	ptr = d.WriteUInt16(ptr, uint16(len(obj.bstrings)))
	for _, s := range obj.bstrings {
		ptr = s.GetBytes(d, ptr)
	}

	ptr = d.WriteUInt16(ptr, uint16(len(obj.rstrings)))
	for _, s := range obj.rstrings {
		ptr = s.GetBytes(d, ptr)
	}

	return ptr
}

func ObjectFromBytes(data serial.Data, offset int) *object {
	var ptr int = offset

	containerType := data.ReadByte(ptr); ptr++
	if containerType != serial.ObjectContainer {
		panic("Unexpected non-object wrapper in data stream: " + serial.ContainerName[containerType])
	}

	l, name := data.ReadString(ptr)
	ptr += l + 2
	println(name)
	obj := SerialObject(name)

	fc := (int)(data.ReadUInt16(ptr)); ptr += 2
	for i := 0; i < fc; i++ {
		obj.AddField(FieldFromBytes(data, ptr))
		ptr += obj.fields[i].GetSize()
	}

	ac := (int)(data.ReadUInt16(ptr)); ptr += 2
	for i := 0; i < ac; i++ {
		obj.AddArray(nil)
		ptr += obj.arrays[i].GetSize()
	}

	bc := (int)(data.ReadUInt16(ptr)); ptr += 2
	for i := 0; i < bc; i++ {
		obj.AddByteString(nil)
		ptr += obj.bstrings[i].GetSize()
	}

	rc := (int)(data.ReadUInt16(ptr)); ptr += 2
	for i := 0; i < rc; i++ {
		obj.AddRuneString(nil)
		ptr += obj.rstrings[i].GetSize()
	}

	return obj
}