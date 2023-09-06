package main

import (
	"io"
	"log"
	"os"
)

func main() {

	var s = "Este texto es copiado hasta la ultima linea."

	_, err := io.WriteString(os.Stdout, s)
	if err != nil {
		log.Fatal(err)
	}
}
