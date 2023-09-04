package e3

import "fmt"

func main() {
	fmt.Println("The salary is", CalculateSalary(2500.0, "C"))
}

func CalculateSalary(minutesPerMonth float64, category string) (salary float64) {
	hours := minutesPerMonth / 60
	switch category {
	case "C":
		{
			salary = hours * 1000
		}
	case "B":
		{
			salary = hours * 1500
			salary += salary * 0.20
		}
	case "A":
		{
			salary = hours * 3000
			salary += salary * 0.50
		}
	}
	return
}
