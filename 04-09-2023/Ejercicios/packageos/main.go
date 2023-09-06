package main

import (
	"fmt"
	"os"
)

func main() {
	createFile()
}

func createFile() {

	fmt.Println("El programa esta comenzando")

	os.Exit(0)

	fmt.Println("El programa esta terminando")
}
