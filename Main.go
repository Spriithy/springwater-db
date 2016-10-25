package main

import (
	"github.com/Spriithy/SerialBits/serial"
)

func main() {
	var data serial.Data = make(serial.Data, 16)
	data.WriteComplex128(0, complex(1.3, 2.0))
	serial.PrintData(data)
	println(data.ReadFloat64(0))
}

