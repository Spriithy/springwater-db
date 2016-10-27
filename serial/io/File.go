package serialio

import (
	"github.com/Spriithy/SerialBits/serial"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func SaveDataToFile(path string, data serial.Data) {
	err := ioutil.WriteFile(path, data, 0644)
	check(err)
}