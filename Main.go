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
	ypos := objects.Float32Field("ypos", 1256.317)

	items := objects.Int8Array("Items", []int8{ 10, 11, 125 })

	player1 := objects.SerialObject("Player12")
	player1.AddString(name)
	player1.AddField(xpos)
	player1.AddField(ypos)
	player1.AddArray(items)

	collection.AddObject(player1)
	data = make(serial.Data, collection.GetSize())

	collection.GetBytes(data, 0)

	serialio.SaveDataToFile("serial.bin", data)
}

func deserializeTest() {
	data, err := serialio.LoadDataFromFile("serial.bin")
	if err != nil {
		panic(err)
	}

	serial.PrintData(data)

	col := objects.CollectionFromBytes(data)
	println(col.GetObject("Player12").GetField("ypos").ValueOf().(float32))
	println(col.GetObject("Player12").GetArray("Items").ValueOf().([]int8)[2])
}

func main() {
	serializeTest()
	deserializeTest()
}

