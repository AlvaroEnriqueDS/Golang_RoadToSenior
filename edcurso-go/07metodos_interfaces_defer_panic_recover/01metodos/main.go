package main

import "fmt"

//Esta es uan estructura o clase
type Persona struct {
	nombre string
	edad int8
}

//NO ES NECESARIO CREAR LA ESTRUCTURA PARA ASIGNARLE UN METODO
type Numero int

//esta asociada a una estructura y la firma es asi:
//este es un metodo de la estruct (clase) persona
func (p Persona) saludar() {
	fmt.Println("Hola soy una persona")
}

func (n Numero) presentarse() {
	fmt.Println("Hola, yo soy un número")
}

//aca modificamos con punteros la informacion de la estructura original
func (p *Persona) asignarEdad(i int8) {
	if i >= 0 {
		p.edad = i
	} else {
		fmt.Println("No es válida la edad negativa")
	}
}

func main() {
	var persona Persona
	// var numero Numero
	persona.saludar()
	// numero.presentarse()
	persona.asignarEdad(-35)
	fmt.Println(persona)
}