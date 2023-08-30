package main 

import (
	"fmt"
)

func main() {
	var numEntrada int32 = 3
	var meses = map[int32]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril", 5: "Mayo", 6: "Junio", 7: "Julio", 8: "Agosto", 9: "Septiembre", 10: "Octubre", 11: "Noviembre", 12: "Diciembre"}
	if numEntrada > 0 && numEntrada <13 {
		fmt.Println(numEntrada,",", meses[numEntrada])
	}
}