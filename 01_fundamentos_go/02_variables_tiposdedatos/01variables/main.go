package main

import "fmt"

func main() {
	//se declaran las variables de la siguiente manera
	//palabra reservada | nombre(s) de variable(s) | tipo de dato
	//var z int
	//var nombre, apellido string

	//GO PUEDE ELEGIR EL TIPO DE DATO POR NOSOTROS
	// := nos permite hacer que go elija el tipo de dato por nosotros, ejemplo:
	//nombre := "Alvaro"
	//apellido := "Diaz"

	//go nos permite declarar y asignar de la siguiente manera con :=
	//:= debe tener una variable nueva, obligatoriamente
	nombre, apellido := "Alvaro", "Diaz"
	fmt.Println(nombre + " " + apellido)
	fmt.Println(z)
}

//var tambien nos permite crear una variable con scope global con:
var z = 41
