package main

import (
	"github.com/Spriithy/SerialBits/serial"
	"github.com/Spriithy/SerialBits/serial/io"
	"github.com/Spriithy/SerialBits/serial/objects"
)

func serializeTest() {
	var data serial.Data

	collection := objects.SerialCollection("Players Info")

	name := objects.ByteString("name", "Foo bar Baz")
	xpos := objects.Int32Field("xpos", 32)
	ypos := objects.Int32Field("ypos", 78)

	player1 := objects.SerialObject("Player12")
	player1.AddByteString(name)
	player1.AddField(xpos)
	player1.AddField(ypos)

	collection.AddObject(player1)
	data = make(serial.Data, collection.GetSize())
	collection.GetBytes(data, 0)

	serialio.SaveDataToFile("serial.bin", data)
}

func main() {
	serializeTest()

	data, err := serialio.LoadDataFromFile("serial.bin")
	if err != nil {
		panic(err)
	}

	serial.PrintData(data)

	objects.CollectionFromBytes(data)

}

