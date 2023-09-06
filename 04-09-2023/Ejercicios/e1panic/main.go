package main

import (
	"fmt"
	"os"
)

func main() {
	defer func(){
		err := recover() // recupera el error de panic 
		fmt.Println(err)
		fmt.Println("ejecución finalizada")
	}()
	_, err := os.Open("customers.txt")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
 }
 