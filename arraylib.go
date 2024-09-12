package arraylib

import (
	"math/rand"
	"time"
)

func ArrayGenerate(size int) []int {
	rand.Seed(time.Now().UnixNano())
	array := make([]int, size)
	for i := 0; i < size; i++ {
		array[i] = rand.Intn(10000) + 1
	}
	return array
}
