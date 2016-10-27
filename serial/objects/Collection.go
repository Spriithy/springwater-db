package objects

import "github.com/Spriithy/SerialBits/serial"

var CollectionHeaderName []byte = ([]byte)("SBC")

const (
	CollectionHeaderMajor byte = 0x00
	CollectionHeaderMinor byte = 0x01
)

type collection struct {
	containerType byte
	nameLength    uint16
	name          []byte

	size          int

	objects       []*object
}

func SerialCollection(name string) *collection {
	col := new(collection)
	col.containerType = serial.CollectionContainer
	col.SetName(name)
	col.size += 7 + len(CollectionHeaderName) + 2
	return col
}

func (col *collection) SetName(name string) {
	if col.size != 0 {
		col.size -= (int)(col.nameLength)
	}

	col.nameLength = (uint16)(len(name))
	col.name = ([]byte)(name)
	col.size += (int)(col.nameLength)
}

func (col *collection) AddObject(obj *object) {
	col.objects = append(col.objects, obj)
	col.size += obj.GetSize()
}

func (col *collection) GetBytes(d serial.Data, ptr int) int {
	ptr = d.WriteBytes(ptr, CollectionHeaderName)
	ptr = d.WriteByte(ptr, CollectionHeaderMajor)
	ptr = d.WriteByte(ptr, CollectionHeaderMinor)
	ptr = d.WriteByte(ptr, col.containerType)
	ptr = d.WriteInt32(ptr, (int32)(col.size))
	ptr = d.WriteUInt16(ptr, col.nameLength)
	ptr = d.WriteBytes(ptr, col.name)

	ptr = d.WriteUInt16(ptr, uint16(len(col.objects)))
	for _, obj := range col.objects {
		ptr = obj.GetBytes(d, ptr)
	}

	return ptr
}

func (col *collection) GetSize() int {
	return col.size + 2
}