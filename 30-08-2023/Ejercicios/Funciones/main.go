package main

import "fmt"
const (
	Suma   = "+"
	Resta  = "-"
	Multip = "*"
	Divis  = "/"
 )
 
func main() {
	
	fmt.Println(operacionAritmetica(6, 2, Suma))
   	fmt.Println(operacionAritmetica(6, 2, Resta))
   	fmt.Println(operacionAritmetica(6, 2, Multip))
   	fmt.Println(operacionAritmetica(6, 2, Divis))
}

func operacionAritmetica(valor1, valor2 float64, operador string) float64 {
   switch operador {
   case Suma:
       return valor1 + valor2
   case Resta:
       return valor1 - valor2
   case Multip:
       return valor1 * valor2
   case Divis:
       if valor2 != 0 {
           return valor1 / valor2
       }
   }
   return 0
}
