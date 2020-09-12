package main

import "fmt"

//ESTA FUNCION ES UN SLICE PREDEFINIDO QUE NOS PERMITE METER MUCHOS DATOS DEL TIPO QUE SE ESPECIFICO
//Es decir nos permite pasar una cantidad variable del tipo que definimos, es decir la variable se comporta como slice
func saludarVarios(edad uint8, nombres ...string) {
	fmt.Printf("%T\n", nombres)
	for _, v := range nombres {
		fmt.Println("Hola", v, "tienes", edad, "años de edad")
	}
}


func main() {
        //usamos la funcion y la pasamos varios string a la variatica
        saludarVarios(20, "Alexys", "Pedro", "Daniel", "Alvaro", "Diana")

        //si queremos desplegar un slice para la variable nombres debemos usar ...
        x := []string{"uno", "dos", "tres"}
        saludarVarios(20, x...)
}
