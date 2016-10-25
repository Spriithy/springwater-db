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
	ptr = d.WriteBytes(ptr, a.data)
	return ptr
}

func (f *array) GetSize() int {
	return 4 + (int)(f.nameLength) + len(f.data) /* 4 = 2*(1) byte + 1*(2) uint16 + */
}

func (a *array) String() string {
	return (string)(a.name) + ": " + serial.ContainerName[a.containerType] + " - " + serial.TypeName[a.dataType]
}