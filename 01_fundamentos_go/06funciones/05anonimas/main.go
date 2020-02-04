package main

import "fmt"

func main() {
	/*
		//UNA VARIABLE ES UNA FUNCION
		anonima := func() {
			fmt.Println("Me imprimo estando en una variable llamada anonima")
		}
		anonima()
	*/

	//SE MANTIENE ACTIVA PARA SECUENCIA1
	secuencia1 := secuencia()
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())

	fmt.Println("----------")

	//SE MANTIENE ACTIVA PARA SECUENCIA2

	secuencia2 := secuencia()
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
}

//ESTE TIPO DE FUNCIONES SE MANTIENEN VIVAS PARA CADA OBJETO O VARIABLE INSTANCIADA
func secuencia() func() int32 {
	var x int32
	return func() int32 {
		x++
		return x
	}
}
