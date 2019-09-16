package main

import (
	"github.com/alvaroenriqueds/golang-es/edcurso-go/07metodos_interfaces_defer_panic_recover/02interfaces/animales"
)

func main() {
	//variables de instancias de las estructuras perro y gato
	var p animales.Perro
	var g animales.Gato
	//se le pone nombre al objeto
	p.Nombre = "Firulais"
	g.Nombre = "Manchas"
	// AdoptarPerro(p)
	// AdoptarGato(g)
	AdoptarMascota(p)
	AdoptarMascota(g)
}

/*
func AdoptarPerro(p animales.Perro) {
	p.Comunicarse()
}
func AdoptarGato(g animales.Gato) {
	g.Comunicarse()
}
*/

//PARA IOMPLEMENTAR UINA IUNTERFACE EL METODO NO PUEDE RECIBIR UN PUNTERO
func AdoptarMascota(m animales.Mascota) {
	m.Comunicarse()
}