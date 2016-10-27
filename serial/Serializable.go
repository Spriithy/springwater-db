package serial


type Serializable interface {
	GetBytes(Data, int) int
	GetSize() int
}
