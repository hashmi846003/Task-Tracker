package utils

import "math/rand"

func GenerateID() int {
	return rand.Intn(1000000)
}
