package calc

import "errors"

// Add returns sum of given integers
func Add(numbers ...int) (int, error) {
	sum := 0

	if len(numbers) < 3 {
		return sum, errors.New("Provide more than 3 numbers")
	} else {
		for _, num := range numbers {
			sum = sum + num
		}
	}
	return sum, nil
}
