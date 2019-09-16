package main
import "fmt"

func main()  {
	//GO PUEDE ELEGIR EL TIPO DE DATO POR NOSOTROS (OBVIANDO LA SIGUIENTE LINEA)
	//var nombre, apellido string

	// := nos permite hacer que go elija el tipo de dato por nosotros
	//nombre := "Alvaro"
	//apellido := "Diaz"

	//go nos permite declarar y asignar de la siguiente manera con :=
	//(:= debe tener una variable nueva, obligatoriamente)
	nombre, apellido := "Alvaro", "Diaz"
	fmt.Print(nombre+" "+apellido)
}