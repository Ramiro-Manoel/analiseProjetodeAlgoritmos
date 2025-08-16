package random

import (
	"errors"
	"math/rand"
)

func RandomArrayGenerator(size int) ([]int, error) {
	if size <= 0 {
		return nil, errors.New("array size must postive")
	}

	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Int()
	}
	return arr, nil
}
