package main

import "fmt"

type Bootcamp struct {
	Organizacion string
	Lenguaje     string
	Alumnos      []Alumno
}

type Alumno struct {
	Nombre       string
	FechaIngreso string
}

func main() {

	BootcampGO := crearBootcap("MELI", "GO")

	alm1 := Alumno{
		Nombre:       "Agustin",
		FechaIngreso: "18/02/22",
	}
	alm2 := Alumno{
		Nombre:       "Vincent",
		FechaIngreso: "18/02/22",
	}
	alm3 := Alumno{
		Nombre:       "Sol",
		FechaIngreso: "18/02/22",
	}

	BootcampGO.guardarAlumno(alm1)
	BootcampGO.guardarAlumno(alm2)
	BootcampGO.guardarAlumno(alm3)

	fmt.Println(BootcampGO)
}

func crearBootcap(org, lng string) Bootcamp {
	boot := Bootcamp{
		Organizacion: org,
		Lenguaje:     lng,
	}

	return boot
}

func (b *Bootcamp) guardarAlumno(alumno Alumno) {
	b.Alumnos = append(b.Alumnos, alumno)
}
