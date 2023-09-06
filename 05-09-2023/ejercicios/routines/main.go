package main

import (
	"fmt"
	"time"
)

func main() {
	go contador("Oveja")
	go contador("Pez")

	time.Sleep(time.Second * 2)
}

func contador(elem string) {
	for i := 1; true; i++ {
		fmt.Println(i, elem)
		time.Sleep(time.Second / 2)
	}
}
