package main

import (
	"errors"
	"fmt"
)

type Myerror struct {
}

func (e *Myerror) Error() string {

	return fmt.Sprintf("Error: el salario es menor a 10.000")
}

func ValidateSalary(val int) (err error) {
	er2 := Myerror{}

	if val < 10000 {
		return &er2
	} else {
		return nil
	}
}

func main() {

	salary := 180

	err := ValidateSalary(salary)

	if err!=nil && errors.Is(err, &Myerror{}) {
		fmt.Println(err.Error())
	} 
	fmt.Println("Fin")
}