package util

import "math/rand"

func RandomNlengthInteger(max, min int) int {
	return rand.Intn((max - min + 1) + min)
}
