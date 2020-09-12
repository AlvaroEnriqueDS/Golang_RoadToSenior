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
