package saludar

import "fmt"

//CUANDO PONES UNA MAYUSCULA AL INICIO SIGNIFICA QUE VA A SER EXPORTADA
func Saludar(nombre string)  {
	fmt.Println("Hola", nombre)
}

//MeVes es para utilizar desde otro paquete
var MeVes string
//este no se puede utilizar desde otro paquete
var noMeVes string
