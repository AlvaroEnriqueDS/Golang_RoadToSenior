package main

import "fmt"

func main() {
	//TAMBIEN CONOCIDOS COMO WHILE (MIENTRAS)
	a := 5
	//esta solo contiene una condiucion que sera evaluada antes de ejecutar toda la sentencia
	for a > 0 {
		a--
		fmt.Println(a)
	}
}
