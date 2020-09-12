package main

import (
        "fmt"
        "github/alvaroenriqueds/Golang_RoadToSenior/01_fundamentos_go/07metodos_interfaces_defer_panic_recover/02interfaces/animales"
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

	//interface vacia
        AdoptarMascotaIV(p)
        AdoptarMascotaIV(g)
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
        //no podemos ingresar directamente a las variables de m porque es de tipo mascota y este solo contiene metodos
        //pero si podemos imprimir todo el objeto
        fmt.Println(m)
        //pero podemos usar esto para poder lograr entrar en este objeto
        //lo que hacemos es buscar en que tipo es m (mascota)
        switch m.(type) {
        //lo que aremos en ambos casos es decirle al compilador que sabemos que tipo es y debe ejecutarlo por afirmacion
        //IDENTIFICADOR.(TIPO DE DATO).CAMPO
        case animales.Gato:
                fmt.Println("IMPRIMO ESTO ENTRANDO AL TIPO(GATO): ", m.(animales.Gato).Nombre)
        case animales.Perro:
                fmt.Println("IMPRIMO ESTO ENTRANDO AL TIPO(PERRO): ", m.(animales.Perro).Nombre)
        }
        //polimorfismo
        //hacemos que el tipo ejecute el metodo que implementa la interfaz
	m.Comunicarse()
}

//interfaz vacia (todos los tipos implementan la interface vacia)
func AdoptarMascotaIV(m interface{})  {
        fmt.Println(m)
        switch m.(type) {
        case animales.Gato:
                fmt.Println("IMPRIMO ESTO ENTRANDO AL TIPO(GATO VACIO): ", m.(animales.Gato).Nombre)
                m.(animales.Gato).Comunicarse()
        case animales.Perro:
                fmt.Println("IMPRIMO ESTO ENTRANDO AL TIPO(PERRO VACIO): ", m.(animales.Perro).Nombre)
                m.(animales.Perro).Comunicarse()
        }
        //aqui no se puede llamar al metodo Comunicarse porque no hay una interfaz que implementar con ese metodo
        //pero aun asi puedo llamar a los metodos de los tipos por afirmacion
}
