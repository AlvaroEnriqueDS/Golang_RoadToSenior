package main

import "fmt"

func main() {
	//el for es:
	//sentancia de inicio
		//i := 0;
	//la expresion de ocndicion
		// i < 5;
	//instruccion postejecucion
		// i++
	/*
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
	*/

	//DENTRO DE LOS FOR PUEDEN JHABER CONDICIONES QUE ALTEREN SUS COMPORTAMIENTO
	/*
		for i := 4; i >= 0; i-- {
			//ESTA CONDICION HACE QUE SALTE LA EVALUACION DEL 3
			if i == 3 {
				continue
			}
			fmt.Println(i)
		}
	*/

	//CON ESTE FOR SI LLEGAMOS AL NUMERO 2 ROMPE EL CICLO
	/*
		for i := 0; i < 5; i++ {
			if i == 2 {
				fmt.Println("i vale 2 y se rompe el ciclo")
				break
			}
			fmt.Println(i)
		}
	*/


	//hacemos el ejercicio de llenar un array bidimensional
	// var matriz [3][3]
	matriz := [3][3]int{}
	valor := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			valor++
			matriz[i][j] = valor
		}
	}

	for fila := 0; fila < 3; fila++ {
		for columna := 0; columna < 3; columna++ {
			fmt.Println(matriz[fila][columna])
		}
	}

}