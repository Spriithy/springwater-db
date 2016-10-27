package main

import (
	"github.com/Spriithy/SerialBits/serial"
	"github.com/Spriithy/SerialBits/serial/io"
	"github.com/Spriithy/SerialBits/serial/objects"
)

func main() {
	var data serial.Data

	collection := objects.SerialCollection("Players Info")

	xpos := objects.Int32Field("xpos", 32)
	ypos := objects.Int32Field("ypos", 78)

	player1 := objects.SerialObject("Player1")
	player1.AddField(xpos)
	player1.AddField(ypos)

	collection.AddObject(player1)
	println(collection.GetSize())
	data = make(serial.Data, collection.GetSize())
	collection.GetBytes(data, 0)

	serialio.SaveDataToFile("serial.bin", data)

}

