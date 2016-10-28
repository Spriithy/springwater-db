package serialio

import (
	"io/ioutil"
	"github.com/Spriithy/SerialBits/serial"
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

func LoadDataFromFile(path string) (serial.Data, error) {
	data, err := ioutil.ReadFile(path)
	return (serial.Data)(data), err
}
