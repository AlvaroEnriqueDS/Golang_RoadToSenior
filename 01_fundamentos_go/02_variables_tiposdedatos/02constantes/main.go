package main

import "fmt"

func main() {
	//constante de tipo String sin necesidad de declararla
	const name = "s"
	fmt.Println(name)

	//tipos de declaracion
	var a int8
	var b int
	n := "Pedro"
	//asignar de esta manera
	a = 123
	b = 127
	//se asigna con := y se hace un CAST con int(a)
	c := b + int(a)
	//esto imprime en consola
	fmt.Println(c)
	//esto le da un formato es decir el %s le estoy pasando una variavle al final
	fmt.Printf("hola %s cómo estás %d? \n ", n, a)
	//esto te permite saber qué tipo de dato es
	fmt.Printf("c es de tipo: %T \n", c)

	//VALOR CERO
	var nombre string
	var numero int
	var entiendes bool
	fmt.Println(nombre, numero, entiendes)

}

//HAY 2 TIPOS DE COMENTARIOS EN GO
/*
	en bloque para comentar muchas lineas a la vez (comentar codigo, metodos , etc)
*/

//comentarios en linea, se utilizan para documentar el código(optimizado)
