package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	var s = "Este texto es copiado hasta la ultima linea."

	r := strings.NewReader(s)

	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}
