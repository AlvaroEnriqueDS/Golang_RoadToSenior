//iniciando esta repo de golang haremos el clásico Hola Mundo

//Todo programa en Go está hecho de paquetes.
//Golang trabaja con paquetes, el programa principal siempre esta contenido en el paquete "main"
//Los programas se empiezan a ejecutar en el paquete main.
package main

//Este programa utiliza los paquetes con la ruta "fmt"
import "fmt"

//la funcion main es la que se ejecuta solo en el paquete main
func main() {
	//La convención es que el nombre del paquete es el mismo que el último elemento de su ruta.
	//en este caso importamos "fmt" y usamos el paquete con ese nombre
	fmt.Print("HELLO WORLD")
}
