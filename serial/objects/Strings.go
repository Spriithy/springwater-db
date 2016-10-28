package objects

import "github.com/Spriithy/SerialBits/serial"

type strings struct {
	containerType byte
	kind          byte
	nameLength    uint16
	name          []byte

	size, length  int
	bytes         []byte
	runes         []rune
}

func ByteString(name, content string) *strings {
	str := new(strings)
	str.containerType = serial.StringContainer
	str.kind = serial.Byte
	str.SetName(name)
	str.length = len(content)
	str.bytes = ([]byte)(content)
	return str
}

func RuneString(name, content string) *strings {
	str := new(strings)
	str.containerType = serial.StringContainer
	str.kind = serial.Rune
	str.SetName(name)
	str.length = len(content)
	str.runes = ([]rune)(content)
	return str
}

func (str *strings) SetName(name string) {
	if str.size != 0 {
		str.size -= (int)(str.nameLength)
	}

	str.nameLength = (uint16)(len(name))
	str.name = ([]byte)(name)
	str.size += (int)(str.nameLength)
}

func (str *strings) GetBytes(d serial.Data, ptr int) int {
	ptr = d.WriteByte(ptr, str.containerType)
	ptr = d.WriteByte(ptr, str.kind)
	ptr = d.WriteUInt16(ptr, str.nameLength)
	ptr = d.WriteBytes(ptr, str.name)
	ptr = d.WriteInt32(ptr, (int32)(str.length))

	switch str.kind {
	case serial.Byte:
		ptr = d.WriteBytes(ptr, str.bytes)
	case serial.Rune:
		ptr = d.WriteRuneArray(ptr, str.runes)
	}
	return ptr
}

func (str *strings) GetSize() int {
	return 8 + len(str.name) + str.length * serial.GetSize(str.kind)
}

func StringFromBytes(data serial.Data, offset int) *strings {
	var ptr int = offset

	containerType := data.ReadByte(ptr); ptr++
	if containerType != serial.StringContainer {
		panic("Unexpected non-string wrapper in data stream: " + serial.ContainerName[containerType])
	}

	kind := data.ReadByte(ptr); ptr++

	l, name := data.ReadString(ptr)
	ptr += l + 2
	print("\t" + name + ": ")

	switch kind {
	case serial.Byte:
		l = (int)(data.ReadInt32(ptr)); ptr += 4
		str := make([]byte, l)
		for i := 0; i < l; i++ {
			str[i] = data.ReadByte(ptr)
			ptr++
		}
		println("\"" + (string)(str) + "\"")
		return ByteString(name, (string)(str))
	case serial.Rune:
		l = (int)(data.ReadInt32(ptr)); ptr += 4
		str := make([]rune, l)
		for i := 0; i < l; i++ {
			str[i] = data.ReadRune(ptr)
			ptr += 4
		}
		println("\"" + (string)(str) + "\"")
		return RuneString(name, (string)(str))
	}

	return nil
}