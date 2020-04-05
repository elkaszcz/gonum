package calc

import (
	"errors"

	"github.com/fatih/color"
)

// Add returns sum of given integers
func Add(numbers ...int) (int, error) {
	sum := 0

	if len(numbers) < 3 {
		errorMessage := color.RedString("Provide more than 3 numbers")
		return sum, errors.New(errorMessage)
	}
	for _, num := range numbers {
		sum = sum + num
	}

	return sum, nil
}
