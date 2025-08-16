package prime

import "math"

func GetPrimeNumbers(array []int) []int {

	var result = []int{}

	for i := range array {
		if isPrimeNumber(array[i]) {
			result = append(result, array[i])
		}
	}
	return result
}

func isPrimeNumber(number int) bool {

	if number < 2 {
		return false
	}

	sqrt := int(math.Sqrt(float64(number)))
	for i := 2; i < sqrt+1; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
