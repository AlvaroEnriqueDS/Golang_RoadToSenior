package main

import "fmt"

func main() {
	//switch o segun
	//esta estructura es para evaluar una variable

	//DECLARAMOS LA VARIABLE A
	var a uint8
	a = 1

	//segun case y sin brake en GO
	switch a {
	case 1:
		fmt.Printf("1 = %d \n", a)
	case 2:
		fmt.Printf("2 = %d \n", a)
	default:
		fmt.Println("ESTA EVALUACION YA NO TIENE SOLUCION")
	}

	//switch puede seguir pasando la evaluacion con la palabra fallthrough
	//esta palabra nos permite tener un mejor control y pasar las evaluciones hasta entontrar una valida (break)
	switch a {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fmt.Println("hola yo si pude")
	case 4:
		fallthrough
	case 5:
		fmt.Println("hola yo tambien")
	default:
		fmt.Println("hubo algun error")
	}

	//ESTA OTRA FORMA DE HACER UN SWITCH DONDE SE DECLARA LA VARIABLE EN EL INICIO
	//se tiene que evaluar

	/*
		switch a := 36; a {
		case 5:
			fmt.Println("Estás entre semana")
		default:
			fmt.Println("No es un día válido")
		}
	*/

	//EN ESTE CASO ESTOY DECLARANDO LA VARIABLE SWITCH Y EVALUANDO EN CADA CASO
	switch a := 7; {
	case a >= 0 && a <= 5:
		fmt.Println("Estás entre semana")
	case a >= 6 && a <= 7:
		fmt.Println("Ests en fin de semana :)")
	default:
		fmt.Println("No es un día válido")
	}

	// OOOOOOOOOOOOOO
	/*
		a := 7
		switch{
		case a == 7:
			fmt.Println("Este si es el numero")
		default:
			fmt.Println("Este no es el numero")
		}
	*/
}
