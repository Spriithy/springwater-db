package main

import (
	"github.com/Spriithy/SerialBits/serial"
	"github.com/Spriithy/SerialBits/serial/objects"
)

func main() {
	var data serial.Data
	data = make(serial.Data, 64)
	a := objects.BooleanArray("Booleans", []bool{true, false, false, true, true})
	println(a.String())
	a.GetBytes(data, 0)
	serial.PrintData(data)
	s := data.ReadString(1)
	println(s)
}

