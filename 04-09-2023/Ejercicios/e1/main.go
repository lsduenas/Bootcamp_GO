package main

import (
	"fmt"
)

type myError struct {
   msg string
}

//hacemos que nuestro type struct implemente el método Error()
func (e *myError) Error() string {
   return fmt.Sprintf("Error: el salario ingresado no alcanza el mínimo imponible%s", e.msg)
}


func main(){
	var salary int64
	salary = 12300000.0
	e:= myError {
		msg: "nuevo error",
	}
	if salary < 150000.0 {
		fmt.Printf(e.Error())
		return
	}
	fmt.Println("Debe pagar impuesto")

}