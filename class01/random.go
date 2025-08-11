package main

import (
	"errors"
	"math/rand"
)

func randomArrayGenerator(size int, max int) ([]int, error) {
	if size <= 0 {
		return nil, errors.New("array size must postive")
	}
	if max <= 0 {
		return nil, errors.New("max number must be positive")
	}

	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(max)
	}
	return arr, nil
}
