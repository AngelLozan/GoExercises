package helpers

import "math/rand"

func RandomNumber(n int) int {
	value:= rand.Intn(n) // n is how big pool of numbers
	return value
}
