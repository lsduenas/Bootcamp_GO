package e4

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestOpMin(t *testing.T){
	t.Run("Calculate minimum value from an slice", func(t *testing.T) {
		// Arrange
		notes:= []float64{2, 3, 3, 4, 10, 2, 4, 5}
		expectedResult:= 2.0

		// Act 
		obtainedResult:= OpMin(notes...)

		//Assert
		assert.Equal(t, expectedResult, obtainedResult)
	})
	t.Run("Calculate maximum value from an slice", func(t *testing.T) {
		// Arrange
		notes:= []float64{2, 3, 3, 4, 10, 2, 4, 5}
		expectedResult:= 10.0

		// Act 
		obtainedResult:= OpMax(notes...)

		//Assert
		assert.Equal(t, expectedResult, obtainedResult)
	})
	t.Run("Calculate average/mean value from an slice", func(t *testing.T) {
		// Arrange
		notes:= []float64{2, 3, 3, 4, 10, 2, 4, 5}
		expectedResult:= 4.125

		// Act 
		obtainedResult:= OpAverage(notes...)

		//Assert
		assert.Equal(t, expectedResult, obtainedResult)
	})
}