package main

import "fmt"

func main() {
	//condicionales en if es
	// or - ||
	// and - &&
	// not !

	//tambien se peude incorporar una expresion inicial y evaluarla luego de ;
	//se peude asignar variables así
	if isValid, nombre := true, "alvaro"; isValid && nombre == "alvaro" {
		fmt.Println("esto es correcto")
		fmt.Printf("%T %s", isValid, nombre)

	}
	fmt.Println("\nEl programa termino")
}
func h() string {
	return "alvaroo"
}
