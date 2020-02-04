package main

import "fmt"

func main() {
	// Arrays
	// Todos los valores deben ser del mismo tipo de dato
	// Tamaño fijo.

	//var nombres [3]string
	/*
		nombres[0] = "Daniel"
		nombres[1] = "Alvaro"
		nombres[2] = "Alexys"
		// nombres[3] = "Pedro"
	*/

	//DE ESTA MANERA TAMBIEN DE PUEDE DECLARAR Y LLENAR UN ARREGLO
	nombres := [3]string{
		"Daniel",
		"Alvaro",
		"Alexys",
	}
	//en esta linea estamos guardando el dato en una variable
	//nombre := nombres[1]
	//en esta linea estamos guardando el tamaño del arreglo con la propiedad len
	size := len(nombres)

	fmt.Println("El tamaño del arreglo es de:", size)
	fmt.Println(nombres)

	// fmt.Println(nombres[inicio:final(excluyente)])
	// fmt.Println(nombres[1:2])
}
