package serial

type Serializable interface {
	Serialize() []byte
	Deserialize([]byte) Serializable
}
