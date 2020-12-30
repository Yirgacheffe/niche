package calc

import "errors"

// ErrorNeedParams - not enough inputs
var ErrorNeedParams = errors.New("please provide more than 2 numbers")

// Add return sum of input numbers
func Add(numbers ...int) (int, error) {
	sum := 0

	if len(numbers) < 2 {
		return sum, ErrorNeedParams
	}

	for _, n := range numbers {
		sum += n
	}

	return sum, nil // return the sum of the input numbers
}
