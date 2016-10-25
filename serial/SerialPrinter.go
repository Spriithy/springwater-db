package serial

import (
	"fmt"
)

func PrintHex(value interface{}) {
	fmt.Printf("0x%02X\n", value)
}

func PrintBin(value interface{}) {
	fmt.Printf("0b%b\n", value)
}

func PrintData(data Data) {
	for i := 0; i < len(data); i++ {
		fmt.Printf("0x%02X ", data[i])
	}
}