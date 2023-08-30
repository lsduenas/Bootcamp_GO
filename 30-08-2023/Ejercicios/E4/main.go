package main 

import (
	"fmt"
)

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println("Benjamin age", employees["Benjamin"])
	counter:=0
	for _, value:= range employees {
		
		if value>22 {
			counter++
		}
	}
	
	fmt.Println("Amount of employees older than 22", counter)

	// new employee
	employees["Federico"] = 25
	fmt.Println(employees)

	// delete 
	delete(employees, "Pedro")
	fmt.Println(employees)
}