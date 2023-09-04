package e3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateSalary(t *testing.T) {
	t.Run("Calculate total salary when category is A", func(t *testing.T) {
		// input - output
		minutesPerMonth := 3000.0
		category := "A"
		expectedResult := 225000.0

		obtainedResult := CalculateSalary(minutesPerMonth, category)

		// the message is showed if the test fails
		assert.Equal(t, expectedResult, obtainedResult, "It can't be that value")
	})
	t.Run("Calculate total salary when category is B", func(t *testing.T) {
		// input - output
		minutesPerMonth := 4000.0
		category := "B"
		expectedResult := 120000.0

		obtainedResult := CalculateSalary(minutesPerMonth, category)

		// the message is showed if the test fails
		assert.Equal(t, expectedResult, obtainedResult, "It can't be that value")
	})
	t.Run("Calculate total salary when category is C", func(t *testing.T) {
		// input - output
		minutesPerMonth := 2500.0
		category := "C"
		expectedResult := 41666.666666666664

		obtainedResult := CalculateSalary(minutesPerMonth, category)

		// the message is showed if the test fails
		assert.Equal(t, expectedResult, obtainedResult, "It can't be that value")
	})
}
