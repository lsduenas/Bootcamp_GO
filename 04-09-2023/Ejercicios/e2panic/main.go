package main

import (
	"bufio"
	"fmt"
	"os"
)

func openFile(file os.File){
	defer func(){
		err := recover() // recupera el error de panic 
		fmt.Println(err)
		fmt.Println("ejecución finalizada")
	}()
	file, err := os.Open("customers.txt")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
}

func readFile(file ){
	defer func(){
		err := recover() // recupera el error de panic 
		fmt.Println(err)
		fmt.Println("ejecución finalizada")
	}()
	scanner, err := bufio.NewScanner(file)
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
	for scanner.Scan() {
		//Print the line
		fmt.Println(scanner.Text())
	}
}

func main() {
	
	
	
 }
 