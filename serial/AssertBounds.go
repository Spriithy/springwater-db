package serial

import "os"

// Just here to assert we don't overflow arrays when writing in Data chunks
func assertBounds(b bool) {
	if !b {
		panic("Accessing (write/read) out of bounds data!")
		os.Exit(1)
	}
}
