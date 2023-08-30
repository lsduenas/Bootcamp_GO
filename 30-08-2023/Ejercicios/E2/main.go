package main 

import (
	"fmt"
)

func main() {
	var (
		edad, antiguedad int = 22, 2
		empleado bool = true
		sueldo float64 = 150000.0
	)
	switch {
	case edad > 22:
		fmt.Println("Usted es menor de 23 años")
	case !empleado: 
		fmt.Println("Usted esta desempleado")
	case antiguedad < 1:
		fmt.Println("Su antiguedad es menor a un año")
	case sueldo < 100000:
		fmt.Println("El sueldo no es suficiente")
	default:
		fmt.Println("Usted puede recibir el prestamo")
	}
}