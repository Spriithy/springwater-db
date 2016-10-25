package serial

import "os"

// Just here to assert we don't overflow arrays when writing in Data chunks
func assertBounds(b bool) {
	switch b {
	case false:
		println("Accessing out of bounds data!")
		os.Exit(1)
	default:
		return
	}
}
