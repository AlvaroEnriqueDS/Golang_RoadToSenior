package main

import "fmt"

//LOS PANIC HACEN UNA ESPECIE DE SEGUIMIENTO DE LOS ERRORES
//HACEN UN TIPO TRACK O SIGUE DE DONDE VIENE O VIENEN LOS ERRORES
func main() {
	division(3, 0)
}


func division(dividendo, divisor int) {
	//DEFER ENTRA A LA COLA DE ESPERA
	defer fmt.Println("Esto se ejecutará pase lo que pase")
	if divisor == 0 {
		//LA FUNCION PANIC NOS DICE DE DONDE VIENE EL ERROR (PANICO)
		panic("No se puede dividir por cero")
	}
	fmt.Println(dividendo / divisor)
}