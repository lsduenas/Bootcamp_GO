package main

import "fmt"

func main() {
	var (
		nombre   string
		apellido string
		telefono int
	)

	fmt.Scanf("%s %s %d", &nombre, apellido, telefono)

	fmt.Println(nombre, apellido, telefono)
}
