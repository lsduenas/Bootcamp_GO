package e5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnimalDog(t *testing.T) {
	t.Run("Calculate amount dog", func(t *testing.T) {
		// Arrange
		amount := 6000.0
		expectedResult := 600.0

		// Act
		obtainedResult := AnimalDog(float64(amount))

		// Assert
		assert.Equal(t, expectedResult, obtainedResult)
	})
	t.Run("Calculate amount cat", func(t *testing.T) {
		// Arrange
		amount := 5000.0
		expectedResult := 1000.0

		// Act
		obtainedResult := AnimalCat(float64(amount))

		// Assert
		assert.Equal(t, expectedResult, obtainedResult)
	})
	t.Run("Calculate amount hamster", func(t *testing.T) {
		// Arrange
		amount := 23345.0
		expectedResult := 93.38

		// Act
		obtainedResult := AnimalHamster(float64(amount))

		// Assert
		assert.Equal(t, expectedResult, obtainedResult)
	})
	t.Run("Calculate amount tarantula", func(t *testing.T) {
		// Arrange
		amount := 22000.0
		expectedResult := 146.66666666666666

		// Act
		obtainedResult := AnimalTarantula(float64(amount))

		// Assert
		assert.Equal(t, expectedResult, obtainedResult)
	})

}
