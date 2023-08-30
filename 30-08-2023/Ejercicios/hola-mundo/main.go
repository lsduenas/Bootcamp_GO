package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
	// Tipos de datos basicos
	// string => cadena de caracteres 
	// int => Numeros enteros 123 - uint
	// float32 => Numero decimal, precisión simple (7 digitos de precisión)
	// float64 => Numero decimal, precisión doble (15 digitos de precisión)

	var username string = "admin"
	// var username string
	// username = "admin"
	// var username = "admin"
    // username := "name" // operador de asignación y declaración
	// var (
	//	username = "admin"
	//	age = 30
	// )
	fmt.Println("Hello, " + username)
	

}