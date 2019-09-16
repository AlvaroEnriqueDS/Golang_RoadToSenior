package main

import "fmt"

// Persona es una estructura
//para crear una estructura se usa type (palabra reservada) y struct (es como una clase)
type Persona struct {
	//se declaran los campos que almacenan valores de tipo
	Nombre string
	Edad   uint8
	Emails []string
}

func main() {
	//ACA VAMOS A "INSTANCIAR" LA ESTRUCTURA PARA ACCEDER A ELLA MEDIANTE UN OBJETO
	/*
		    var persona1 Persona
			persona1.Nombre = "Pedro"
			persona1.Edad = 20
			fmt.Println(persona1.Nombre)
	*/

	emails := []string{"a@b.com", "z@b.com"}

	//SE CREA UN NUEVO OBJETO DE LA ESTREUCTURA PERSONA
	persona2 := Persona{
		"Pablo",
		33,
		emails,
	}
	fmt.Println(persona2)

}