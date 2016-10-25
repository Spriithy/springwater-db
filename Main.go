package main

import (
	"github.com/Spriithy/SerialBits/serial"
)

func main() {
	var data serial.Data = make(serial.Data, 16)
	data.WriteString(0, "Hello!")
	println(data.ReadString(0))
	serial.PrintData(data)
}

