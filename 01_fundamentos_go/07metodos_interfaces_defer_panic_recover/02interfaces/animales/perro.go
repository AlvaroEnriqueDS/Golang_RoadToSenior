package animales

import "fmt"

//estructura o clase
type Perro struct {
	Nombre string
}

//metodo de la estructura perro
func (p Perro) Comunicarse() {
	fmt.Println("woffff")
}
