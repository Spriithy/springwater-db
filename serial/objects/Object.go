package objects

import "github.com/Spriithy/SerialBits/serial"

type object struct {
	containerType byte
	nameLength    uint16
	name          []byte
	size          int

	// TODO Implement hashtable for object offsets in byte array
	fields        []*field
	arrays        []*array
}

func SerialObject(name string) *object {
	obj := new(object)
	obj.containerType = serial.ObjectContainer
	obj.SetName(name)
	obj.size = 7
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

func (obj *object) GetSize() int {
	return obj.size
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

	return ptr
}