package main

import (
	"fmt"
)

func main() {
	fmt.Println("Conectando a la base de datos...")
	//CON DEFER TE ASEGURAS QUE SI O SI SE EJECUTE
	//defer es una "sala de espera"
	defer cerrarDB()
	fmt.Println("Consultamos información, set de datos")
	//pase lo que pase defer ejecutara el codigo para que no caiga la aplicacion
	defer cerrarSetDeDatos()

	//CON LA PALABRA RETURN SE ACABA LA FUNCIION
	//O LLEGO HASTA LA LLAVE DE CIERRE
	//AL LLEGAR A ESTOS 2 O ENTRA EN PANICO SE EJECUTA LOS DEFER

	panic("ya fue todo")
	fmt.Println("NO DEBERIA LLEGAR HASTA AQUÍ")
}

func cerrarDB() {
	fmt.Println("Cerrar la base de datos")
}

func cerrarSetDeDatos() {
	fmt.Println("Cerra el set de datos")
}