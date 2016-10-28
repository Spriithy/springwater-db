package objects

import (
	"github.com/Spriithy/SerialBits/serial"
)

var CollectionHeaderName []byte = ([]byte)("SBC")

const (
	CollectionHeaderMajor byte = 0x00
	CollectionHeaderMinor byte = 0x01
)

func ValidateHeader(header []byte) {
	hn := CollectionHeaderName
	maj := CollectionHeaderMajor
	min := CollectionHeaderMinor
	if ! (header[0] == hn[0] && header[1] == hn[1] && header[2] == hn[2] && header[3] == maj && header[4] == min) {
		panic("Loading unapropriate data from file. Invalid file header!")
	}
}

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

func (col *collection) GetSize() int {
	return col.size - 2
}

func (col *collection) GetBytes(d serial.Data, ptr int) int {
	ptr = d.WriteBytes(ptr, CollectionHeaderName)
	ptr = d.WriteByte(ptr, CollectionHeaderMajor)
	ptr = d.WriteByte(ptr, CollectionHeaderMinor)
	ptr = d.WriteByte(ptr, col.containerType)
	ptr = d.WriteUInt16(ptr, col.nameLength)
	ptr = d.WriteBytes(ptr, col.name)

	ptr = d.WriteUInt16(ptr, uint16(len(col.objects)))
	for _, obj := range col.objects {
		ptr = obj.GetBytes(d, ptr)
	}

	return ptr
}

func CollectionFromBytes(data serial.Data) *collection {
	var ptr int = 0

	ValidateHeader(data[:2 + len(CollectionHeaderName)])
	ptr += 2 + len(CollectionHeaderName)

	containerType := data.ReadByte(ptr); ptr++
	if containerType != serial.CollectionContainer {
		panic("Unexpected non-collection top level object in data stream: " + serial.ContainerName[containerType])
	}

	l, name := data.ReadString(ptr)
	ptr += l + 2

	col := SerialCollection(name)

	count := (int)(data.ReadUInt16(ptr)); ptr += 2
	for i := 0; i < count; i++ {
		col.AddObject(ObjectFromBytes(data, ptr))
	}


	return col
}