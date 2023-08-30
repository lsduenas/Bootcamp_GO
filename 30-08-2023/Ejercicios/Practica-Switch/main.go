package main

import (
	"fmt"
	
)

func main() {
	meses := []string{"Enero", "Febrero", "Junio", "Agosto", "Abril"}
	fmt.Println("Utilizando condicionales")
	ObtenerEstacion(meses)
	fmt.Println("Utilizando switch case")
	ObtenerEstacionSwitch(meses)
	
	if v:=3; v<5{
		println(v)
	}

	//fallthrough ejecuta el otro case por mas que no sea cierto !!!
	//Cuidado!!
	switch dias:="Lunes";dias{

	case "Lunes","Martes","Miercoles":
		fmt.Println("Son los primeros 3 dias")
		fallthrough
	case "Jueves","Viernes":
		fmt.Println("Son lo 2 ultimos dias")

	
	default:
		fmt.Println("Es Finde de semana")
	}
}

func ObtenerEstacion(meses []string) {
	for _, mes := range meses {
		if mes == "Enero" || mes == "Febrero" || mes == "Marzo" {
			fmt.Println("Verano")
		} else if mes == "Abril" || mes == "Mayo" || mes == "Junio" {
			fmt.Println("Otoño")
		} else if mes == "Julio" || mes == "Agosto" || mes == "Septiembre" {
			fmt.Println("Invierno")
		} else if mes == "Octubre" || mes == "Noviembre" || mes == "Diciembre" {
			fmt.Println("Primavera")
		} else {
			fmt.Println("No existe este mes")
		}
	}
}

func ObtenerEstacionSwitch(meses []string) {
	for _, value := range meses {
		switch value {
		case "Enero", "Febrero", "Marzo":
			fmt.Println("dfsdf")
			
		
		case "Abril", "Mayo", "Junio":
			fmt.Println("Otoño")

		case "Julio", "Agosto", "Septiembre":
			fmt.Println("Invierno")

		case "Octubre", "Noviembre", "Diciembre":
			fmt.Println("Invierno")
		default:
			fmt.Println("Tropical")
	}
}
}