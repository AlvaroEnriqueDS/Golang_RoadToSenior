package main

import "fmt"


//una funcion se crea con la palabra func (funcion)
//A UNA FUNCION SE EL PUEDEN PASAR PARAMETROS

//receptor -> es un vlor que recibe una funcion | funcion adjutnada (se le llama metodo)
//identificador -> nombre de la funcion
//parametros -> valor que se le pasa a una funcion, son argumentos
//retorno -> lo que la funcion retorna, puede tener multiples retornos
//codigo -> algoritmo, codigo de la funcion
// func (r receptor) identificador(parametros) retorno(s) { código}


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
