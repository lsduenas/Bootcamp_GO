package e2

import (
	"github.com/stretchr/testify/assert" // Se importa testify
	"testing"
)

func TestCalculateMean(t *testing.T) {
	// input - output
	notes := []float64{5.0, 4.5, 3.0, 5}
	expectedResult := 4.375

	// Test
	result, _ := CalculateMean(notes...)

	// validate results
	assert.Equal(t, expectedResult, result, "It should be equals")
}
