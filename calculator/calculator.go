package calculator

import (
	"BasicCalculatorApp/mathops"
	"errors"
	"math"
)

// Function to calculate Addition
func Add(a float64, b float64) float64 {
	return mathops.Addition(a, b)
}

// Function to calculate Subtraction
func Subtract(a float64, b float64) float64 {
	return mathops.Subtraction(a, b)
}

// Function to calculate Multiplication
func Multiply(a float64, b float64) float64 {
	return mathops.Multiplication(a, b)
}

// Function to calculate Division
func Divide(a float64, b float64) (float64, error) {
	return mathops.Division(a, b)
}

// Function to calculate The Power Of
func PowerOf(a float64, b float64) (float64, error) {
	// Error checking, the exponent number can't be negative
	if b < 0 {
		// Return an error
		return 0, errors.New("invalid power")
	}

	return math.Pow(a, b), nil
}

// Function to calculate Square Root
func SquareRoot(a float64) (float64, error) {
	// Error checking, the square root can't be negative
	if a < 0 {
		// Return an error
		return 0, errors.New("can't square root a negative number")
	}

	return math.Sqrt(a), nil
}
