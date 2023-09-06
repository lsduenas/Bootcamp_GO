package main

import (
	"fmt"
)

func main() {
	n := 3

	ch := make(chan int)

	go multiplicarPorDos(n, ch)

	fmt.Println(<-ch)
}

func multiplicarPorDos(num int, ch chan int) {
	res := num * 2
	ch <- res
}
