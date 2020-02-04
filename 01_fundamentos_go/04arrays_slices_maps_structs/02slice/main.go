package main

import "fmt"

func main() {
	// Slices
	// 1. Apuntador a un array
	// 2. Tamaño (no es fijo)
	// 3. Capacidad

	//ASI SE CREAN LOS SLICE
	// var nombres []string

	// make (tipo de dato del slice, tamaño inicial, capacidad inicial)
	// nombres := make([]string, 0)

	nombres := []string{
		"Alvaro",
		"Daniel",
		"Alexys",
	}

	/*
		fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
		nombres = append(nombres, "Daniel")
		fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
		nombres = append(nombres, "Alvaro")
		fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
		nombres = append(nombres, "Alexys")
		fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
		nombres = append(nombres, "Pedro")
		fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
		nombres = append(nombres, "Juan")
		fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
		nombres[3] = "Juan"
	*/
	fmt.Println(nombres)

	// con : puedes dividir por partes un slice y puedes meterlo dentro de otro
	// x = append(x[:2], x[:4]...)
	//añadirle append
	//eliminar se usa (...) algoritmos...

	//slice multibidimensionales
	et := []string{"uno", "dos", "tres"}
	fmt.Println(et)

	vp := [][]string{et, nombres}
	fmt.Println(vp)
	fmt.Println(vp[1][2])
}
