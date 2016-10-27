package objects

import "github.com/Spriithy/SerialBits/serial"

const (
	ByteStringKind byte = 0x00
	RuneStringKind byte = 0xFF
)

type bytestring struct {
	containerType byte

	nameLength    uint16
	name          []byte

	size, length  int
	string        []byte
}

type runestring struct {
	containerType byte

	nameLength    uint16
	name          []byte

	size, length  int
	string        []rune
}

func ByteString(name, content string) *bytestring {
	str := new(bytestring)
	str.containerType = serial.StringContainer
	str.SetName(name)
	str.length = len(content)
	str.string = ([]byte)(content)
	str.size += 11
	return str
}

func (str *bytestring) SetName(name string) {
	if str.size != 0 {
		str.size -= (int)(str.nameLength)
	}

	str.nameLength = (uint16)(len(name))
	str.name = ([]byte)(name)
	str.size += (int)(str.nameLength)
}

func (str *bytestring) GetBytes(d serial.Data, ptr int) int {
	ptr = d.WriteByte(ptr, str.containerType)
	ptr = d.WriteByte(ptr, ByteStringKind)
	ptr = d.WriteUInt16(ptr, str.nameLength)
	ptr = d.WriteBytes(ptr, str.name)
	ptr = d.WriteInt32(ptr, (int32)(str.length))
	ptr = d.WriteBytes(ptr, str.string)
	return ptr
}

func (str *bytestring) GetSize() int {
	return str.size + (int)(str.nameLength) + str.length
}

func RuneString(name, content string) *runestring {
	str := new(runestring)
	str.containerType = serial.StringContainer
	str.SetName(name)
	str.length = len(content)
	str.string = ([]rune)(content)
	str.size += 11
	return str
}

func (str *runestring) SetName(name string) {
	if str.size != 0 {
		str.size -= (int)(str.nameLength)
	}

	str.nameLength = (uint16)(len(name))
	str.name = ([]byte)(name)
	str.size += (int)(str.nameLength)
}

func (str *runestring) GetBytes(d serial.Data, ptr int) int {
	ptr = d.WriteByte(ptr, str.containerType)
	ptr = d.WriteByte(ptr, RuneStringKind)
	ptr = d.WriteUInt16(ptr, str.nameLength)
	ptr = d.WriteBytes(ptr, str.name)
	ptr = d.WriteInt32(ptr, (int32)(str.length))
	ptr = d.WriteRuneArray(ptr, str.string)
	return ptr
}

func (str *runestring) GetSize() int {
	return str.size + (int)(str.nameLength) + serial.GetSize(serial.Rune) * str.length
}