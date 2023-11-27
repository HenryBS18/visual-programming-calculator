package calculator

import (
	"math"
	"testing"

	"BasicCalculatorApp/error"

	"github.com/stretchr/testify/assert"
)

// Addition Test
func TestAddition(t *testing.T) {
	result := Add(30, 20)

	assert.Equal(t, float64(50), result)
}

// Subtraction Test
func TestSubtraction(t *testing.T) {
	result := Subtract(30, 20)

	assert.Equal(t, float64(10), result)
}

// Multiplication Test
func TestMultiplication(t *testing.T) {
	result := Multiply(5, 4)

	assert.Equal(t, float64(20), result)
}

// Division Test
func TestDivision(t *testing.T) {
	result, err := Divide(100, 20)
	error.IsError(err)

	assert.Equal(t, float64(5), result)

	// Division Sub Test, test number divided by 0
	t.Run("divided by 0", func(t *testing.T) {
		result, err := Divide(20, 0)
		error.IsError(err)

		assert.Equal(t, math.Inf(1), result)
	})

	// test 0 divided by any number
	t.Run("0 divided by any", func(t *testing.T) {
		result, err := Divide(0, 1)
		error.IsError(err)

		assert.Equal(t, float64(0), result)
	})
}

// To The Power Of Test
func TestPowerOf(t *testing.T) {
	result, err := PowerOf(2, 3)
	error.IsError(err)

	assert.Equal(t, float64(8), result)

	// To The Power Of Sub Test
	t.Run("to the power of 0", func(t *testing.T) {
		result, err := PowerOf(2, 0)
		error.IsError(err)

		assert.Equal(t, float64(1), result)
	})
}

// Square Root Test
func TestSquareRoot(t *testing.T) {
	result, err := SquareRoot(100)
	error.IsError(err)

	assert.Equal(t, float64(10), result)

	// Square Root Sub Test
	t.Run("square root of 0", func(t *testing.T) {
		result, err := SquareRoot(0)
		error.IsError(err)

		assert.Equal(t, float64(0), result)
	})
}
