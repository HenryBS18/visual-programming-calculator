package mathops

import (
	"BasicCalculatorApp/error"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Addition Test
func TestAddition(t *testing.T) {
	result := Addition(30, 20)

	assert.Equal(t, float64(50), result)
}

// Subtraction Test
func TestSubtraction(t *testing.T) {
	result := Subtraction(30, 20)

	assert.Equal(t, float64(10), result)
}

// Multiplication Test
func TestMultiplication(t *testing.T) {
	result := Multiplication(5, 4)

	assert.Equal(t, float64(20), result)
}

// Division Test
func TestDivision(t *testing.T) {
	result, err := Division(100, 20)
	error.IsError(err)

	assert.Equal(t, float64(5), result)

	// Division Sub Test, test number divided by 0
	t.Run("divided by 0", func(t *testing.T) {
		result, err := Division(20, 0)
		error.IsError(err)

		assert.Equal(t, math.Inf(1), result)
	})

	// test 0 divided by any number
	t.Run("0 divided by any", func(t *testing.T) {
		result, err := Division(0, 1)
		error.IsError(err)

		assert.Equal(t, float64(0), result)
	})
}
