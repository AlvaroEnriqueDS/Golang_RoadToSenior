package main

import "fmt"

func main() {
	saludarVarios(20, "Alexys", "Pedro", "Daniel", "Alvaro", "Diana")
}

//ESTA FUNCION ES UN SLICE PREDEFINIDO QUE NOS PERMITE METER MUCHOS DATOS DEL TIPO QUE SE ESPECIFICO
func saludarVarios(edad uint8, nombres ...string) {
	fmt.Printf("%T\n", nombres)
	for _, v := range nombres {
		fmt.Println("Hola", v, "tienes", edad, "años de edad")
	}
}