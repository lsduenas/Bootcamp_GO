package e1

import "fmt"

func main() {
	salary := 80000.0
	fmt.Println("Taxes", CalculateTaxes(salary))
}

func CalculateTaxes(salary float64) (taxes float64) {
	if salary > 50000 {
		taxes = salary * 0.17
		if salary > 150000 {
			taxes += salary * 0.10
		}
	} else {
		taxes = 0.0
	}
	return
}
