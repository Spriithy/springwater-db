package main

import (
	"github.com/Spriithy/SerialBits/serial"
	"github.com/Spriithy/SerialBits/serial/objects"
)

func main() {
	var data serial.Data
	data = make(serial.Data, 32)
	f := objects.Complex128Field("Test", 145.2 + 2.43i)
	f.GetBytes(data, 0)
	println(f.String())
	serial.PrintData(data)
}

