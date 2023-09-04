package main

import "fmt"

// documentation in go: https://go.dev/doc/comment
// godoc -> file

// Person is a struct that represents a person
// for our example
type Person struct {
	// ID is the unique identifier for a person
	ID int

	// Name is the full name of a person
	Name string

	// DateOfBirth is the birthday of a person
	DateOfBirth string // package time
}

func (e Employee) PrintEmployee(){
	fmt.Println(
		"ID", e.ID,
		"Position", e.Position,
		"Person", e.Person,
	)
}

// Employee is a struct of a person who is a employee
type Employee struct {
	Person //composition

	// ID
	ID int
	Position string

}

func main(){
	person := Person {
		ID: 123,
		Name: "John Doe",
		DateOfBirth: "05-11-1999",
	}

	employee := Employee{
		Person: person,
		ID: 896,
		Position: "Software Engineer",
	}

	fmt.Println(employee.Name)
	employee.PrintEmployee()
}

