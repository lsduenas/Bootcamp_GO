package main

import (
	"encoding/json"
	"fmt"
)

type hobbies struct {
	comidas string `json:"comidas"`
	peliculas string `json:"peliculas"`
	series string `series:series"`
	animes string `animes:"animes"`
	deporte string `deporte:"deporte"`
}

type persona struct {
	nombre string `json:"nombre"`
	apellido string `json:"apellido"`
	edad int16 `json:"edad"`
	gustos hobbies `json:"gustos"`
}

func main(){
	// Instanciación
	hobbiesPepito := hobbies {
		"Arroz con pollo",
		"Inception",
		"Friends",
		"Inuyasha",
		"Natación",
	}

	pepito := persona{
		"Pepito",
		"Perez",
		24,
		hobbiesPepito,
	}
	fmt.Println("Datos de", pepito.nombre)
	fmt.Println(pepito)
	pepito.nombre = "Ronaldo"
	fmt.Println("Datos de", pepito.nombre)
	fmt.Println(pepito)

	// JSON
	miJSON, err := json.Marshal(pepito)
	fmt.Println("Estrcutura JSON")
	fmt.Println(string(miJSON))
	fmt.Println(err)
}