package main

import "fmt"

func main() {
	f()
}

func f() {
	defer func() { //CON ESTO LE DICES A LA FUNCON DEFER QUE SE AUTOEJECUTE
					//PARA QUE AVELUE Y RECUPERE EL PANIC PARA QUE SIGA EL PROGRAMA
		if r := recover(); r != nil {
			fmt.Printf("%T\n", r)
			fmt.Println("Recuperado en f:", r)
		}
	}()

	fmt.Println("Llamando g.")
	g(3)
	fmt.Println("Retorna normalmente desde g")
}

func g(i int) {
	if i > 3 {
		fmt.Println("AAAAAAAAaaaaaa")
		panic("no puede ser mayor a 3")
	}
	defer fmt.Println("Defer en la función g")
	fmt.Println("Imprimiendo en g", i)
}