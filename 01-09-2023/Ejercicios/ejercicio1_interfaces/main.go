package main

import "fmt"

type Alumno struct {
	Nombre string
	Apellido string
	DNI int64
	Fecha string
}

type Operaciones interface {
	obtenerDetalles()
}

func (a Alumno) obtenerDetalles(Alumno){  //definición del método de la interfaz
	fmt.Println(a)
}

func main(){
	// Instanciación
	a:= Alumno{
		Nombre: "Laura",
		Apellido: "Dueñas",
		DNI: 123456789,
		Fecha: "01-09-2023",
	}
	a.obtenerDetalles(a)

}