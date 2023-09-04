package main

import "fmt"

func main() {
	v := 5

	var p *int
	p = &v

	fmt.Println("variable:", v)

	fmt.Println("direccion de la variable:", &v)

	//deberia ser igual a la direccion de memoria asignada a la variable
	fmt.Println("direccion a la que apunta el puntero:", p)
	//deberia ser igual al valor de la variable
	fmt.Println("valor que contiene la direccion que contiene el puntero (desreferencia):", *p)
}