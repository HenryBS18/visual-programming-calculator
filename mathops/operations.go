package mathops

import (
	"errors"
	"math"
)

// Function to calculate Addition
func Addition(a float64, b float64) float64 {
	return a + b
}

// Function to calculate Subtraction
func Subtraction(a float64, b float64) float64 {
	return a - b
}

// Function to calculate Multiplication
func Multiplication(a float64, b float64) float64 {
	return a * b
}

// Function to calculate Division
func Division(a float64, b float64) (float64, error) {
	if b == 0 {
		return float64(math.Inf(1)), errors.New("cannot divide by zero")
	}

	return a / b, nil
}
