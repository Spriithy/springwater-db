package main

import (
	"github.com/Spriithy/SerialBits/serial"
)

func main() {
	var data serial.Data = make(serial.Data, 16)
	data.WriteString(0, "Cherno")
	serial.PrintData(data)
}

