package objects

type Field struct {
	nlen    uint16
	elcount int32
	typ     byte

	name    []byte
	data    []byte
}

func (f *Field) SetName(name string) {
	f.nlen = (uint16)(len(name))
	f.name = ([]byte)(name)
}