package e1

import "testing"

func TestCalculateTaxes(t *testing.T) {
	// test 1---------
	salary1 := 50000.0
	expectedResult1 := 0

	// test
	result1 := CalculateTaxes(salary1)

	// validate results
	if expectedResult1 != int(result1) {
		t.Errorf("Unexpected result %v, expected was %v", result1, expectedResult1)
	}

	// test 2 -----------
	salary2 := 80000.0
	expectedResult2 := 13600.000000000002

	// test
	result2 := CalculateTaxes(salary2)

	// validate results
	if expectedResult2 != result2 {
		t.Errorf("Unexpected result %v, expected was %v", result2, expectedResult2)
	}

	// test 3 -------

	salary3 := 1000000.0
	expectedResult3 := 270000

	// test
	result3 := CalculateTaxes(salary3)

	// validate results
	if expectedResult3 != int(result3) {
		t.Errorf("Unexpected result %v, expected was %v", result3, expectedResult3)
	}

}
