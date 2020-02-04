package main

import "fmt"

//para crear una estructura se usa type (palabra reservada) y struct (es como una clase)
//lo que hacemos es crear un tipo, en este paso en tipo persona y podemos crear valores de el
// Persona es una estructura
type Persona struct {
	//se declaran los campos que almacenan valores de tipo
	Nombre string
	Edad   uint8
	Emails []string
}

type Amigos struct {
	Persona
	identidicador string
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
		//[]string{"", ""},
		emails,
	}
	fmt.Println(persona2)

	//EJEMPLO DE STRUCT EMBEBIDOS
	amg1 := Amigos{
		Persona: Persona{
			Nombre: "Alvaro",
			Edad:   22,
			Emails: emails,
		},
		identidicador: "UNICO",
	}
	fmt.Println(amg1)
	//los datos persona se promueven a amigos, es decir con amg1.nombre accedo al valor del tipo.
	fmt.Println(amg1.Nombre)

	//ESTRUCTURAS ANONIMAS
	//son tipos anonimos, es decir no tiene nombre ni campos definidos, sino lo crea directamente
	anomi := struct {
		campo1 string
		campo2 int
	}{
		campo1: "PRUEBA",
		campo2: 1,
	}
	fmt.Println(anomi)

}
