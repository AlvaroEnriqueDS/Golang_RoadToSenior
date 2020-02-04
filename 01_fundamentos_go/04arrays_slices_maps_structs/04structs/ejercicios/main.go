package main

import "fmt"

type Persona struct {
	Nombre   string
	Apellido string
	sabores  []string
}

func main() {
	p1 := Persona{
		Nombre:   "Alvaro",
		Apellido: "Diaz",
		sabores:  []string{"uno", "dos"},
	}

	p2 := Persona{
		Nombre:   "Alvaro2",
		Apellido: "Diaz2",
		sabores:  []string{"uno2", "dos2"},
	}

	m := map[string]Persona{
		p1.Apellido: p1,
		p2.Apellido: p2,
	}

	fmt.Print(p1.Nombre, " ")
	fmt.Println(p1.Apellido)

	for i, valor := range p1.sabores {
		fmt.Println("\t", i, ": ", valor)
	}

	fmt.Print(p2.Nombre, " ")
	fmt.Println(p2.Apellido)

	for i, valor := range p2.sabores {
		fmt.Println("\t", i, ": ", valor)
	}

	//llave k, valor v
	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v)
		//Sabores de v
		for _, i := range v.sabores {
			fmt.Println(i)
		}
	}
}
