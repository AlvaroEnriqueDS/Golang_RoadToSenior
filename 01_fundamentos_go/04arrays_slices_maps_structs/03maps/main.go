package main

import "fmt"

func main() {
	/*
		//MAPS DECLARAR PARA AGREGARLES LUEGO EL VALOR
		//NOMBRE DE VARIABLE, := , MAKE, TIPO MAP, [TIPO DE LLAVE], TIPO DE DATO ALMACENADO
		idiomas := make(map[string]string)
		idiomas["es"] = "Español"
		idiomas["en"] = "Inglés"
		idiomas["fr"] = "Francés"
	*/

	/*
		//MAPS ASIGNANDOLES DIRECTAMENTE EL VALOR
		idiomas := map[string]string{
			"es": "Español",
			"en": "Inglés",
			"fr": "Francés",
			"pt": "Portugués",
		}
	*/
	/*
		fmt.Println(idiomas)

		//CON ESTA FUNCION SE ELIMINA LAS POSICIONES DE UN MAPA
		delete(idiomas, "en")

		fmt.Println(idiomas)
	*/

	//EVALUAMOS SI UNA POSICION EXISTE EN EL MAPA MEDIANTE UN IF
	/*
				if idioma, ok := idiomas["pt"]; ok {
		        		fmt.Println("La posición pt sí existe y vale", idioma)
		        	} else {
					fmt.Println("La posición pt NO existe")
				}
	*/

	numeros := map[uint16]string{
		1:    "Uno es un número chiquito",
		2016: "Esto es otro número",
		4:    "Aquí la llave es negativa",
	}

	fmt.Println(numeros)
}
