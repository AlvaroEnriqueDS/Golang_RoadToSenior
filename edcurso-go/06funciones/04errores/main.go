package main

import (
	"errors"
	"fmt"
)

func main() {
	r, err := division(100, 0)
	//NIL ES EL NULL DEL TIPO ERROR
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(r)
}

//ERROR ES UN TIPO DE DATO
func division(dividendo, divisor float64) (resultado float64, err error) {
	if divisor == 0 {
		//CON ERRORS.NEW LE ESTAS DICIENDO QUE ES DE TIPO ERROR Y EL MENSAJE
		err = errors.New("No se puede dividir por cero")
	} else {
		resultado = dividendo / divisor
	}
	return
}