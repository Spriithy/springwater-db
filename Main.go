package main

import (
	"github.com/Spriithy/SerialBits/serial"
	"github.com/Spriithy/SerialBits/serial/io"
	"github.com/Spriithy/SerialBits/serial/objects"
	"math/rand"
)

func main() {
	var data serial.Data
	data = make(serial.Data, 4 * 50 + 64)
	stream := make([]int32, 50)
	for i := 0; i < 50; i++ {
		stream[i] = rand.Int31()
	}
	array := objects.Int32Array("RandomNumbers", stream)

	object := objects.SerialObject("Entity")
	object.AddArray(array)
	object.AddField(objects.Int8Field("Integer", 8))
	object.GetBytes(data, 0)

	serialio.SaveDataToFile("serial.bin", data)

}

