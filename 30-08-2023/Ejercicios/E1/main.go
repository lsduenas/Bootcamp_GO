package main

import (
	"fmt"
	"strings" 
)

func main() {
	var phrase string = "Hola mundo"
	length:= len(phrase)
	
	fmt.Println("la cantidad de letras de esta palabra es ",length)
	for _, letter:= range phrase{
		fmt.Println("La letra es ", letter)
		fmt.Println("La letra es ", string(letter))
	}	
	for _, letter:= range strings.Split(phrase, ""){
		fmt.Println("La letra es ", letter)
		
	}	
}