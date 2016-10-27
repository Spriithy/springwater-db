package serial

const (
	UnknownContainer    byte = 0x00
	FieldContainer      byte = 0x01
	ArrayContainer      byte = 0x02
	ObjectContainer     byte = 0x03
	StringContainer     byte = 0x04
	CollectionContainer byte = 0x05
)

var (
	ContainerName = map[byte]string{
		UnknownContainer    : "Unknown",
		FieldContainer      : "Field",
		ArrayContainer      : "Array",
		ObjectContainer     : "Object",
		StringContainer     : "String",
		CollectionContainer : "Collection",
	}
)