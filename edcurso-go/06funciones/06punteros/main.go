package main

import "fmt"

func main() {
	//UN PUNTERO ES LA DIRTERCCION DE MEMORIA DE UNA VARIABLE
	//SE USA & ANTES DE LA VARIABLE PARA ACCEDER A LA DIRECCION DE MEMORIA
	a := 3
	fmt.Println("Antes de duplicar, a =", a)
	//AL PASARLE EL PUNTERO MODIFICA DIRECTAMENTE ESTA VARIABLE
	duplicar(&a)
	//SE MODIFICA DIRECTAMENTE EL ORIGINAL EN DIRECCION DE MEMORIA
	fmt.Println("Después de duplicar, a =", a)
}

//esta funcion va a modificar una variable internamente
//* ES EL PUNTERO
func duplicar(puntero *int) {
	*puntero *= 2
	fmt.Println("Dentro de la función duplicar a =", *puntero)
}

/*
package main

import"fmt"

func pointerTest(){
	a := 100
	var b * int
	b = &a
	fmt.Println("Sin modificar")
	fmt.Println(a, *b)
	fmt.Println(&a, b)
	pointerModify(b)
	fmt.Println("Con una modificación")
	fmt.Println(a, *b)
	fmt.Println(&a, b)
}

func pointerModify(c  *int) {
	*c = 10
}

func main(){
	pointerTest()
}
*/

/*
package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	v := new(Vertex)
	fmt.Println(v)
	v.X, v.Y = 11, 9
	fmt.Println(v)

	var t *Vertex = new(Vertex)
	fmt.Println(t)
	t.X, t.Y = 1, 2
	fmt.Println(t)

}
*/
