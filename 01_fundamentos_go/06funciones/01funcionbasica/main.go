package main

import "fmt"

//A UNA FUNCION SE EL PUEDEN PASAR PARAMETROS
func saludar(nombre string, edad uint8) {
	fmt.Printf("Hola %s tienes %d años de edad\n", nombre, edad)
	//ESTA CONDICION SIRVE PARA RETORNAR EL VALOR
	if edad > 30 {
		return
	}
	fmt.Println("Eres joven")
}

func main() {
	saludar("Alexys", 25)
}
