package main

import (
	"fmt"

	"github.com/alvaroenriqueds/golang-es/edcurso-go/06funciones/07paquetes/despedida"
	"github.com/alvaroenriqueds/golang-es/edcurso-go/06funciones/07paquetes/saludar"
)

func main() {
	nombre := "gente del futuro"

	saludar.Saludar(nombre)

	saludar.MeVes = "Esto es un texto asignado desde el main"
	fmt.Println(saludar.MeVes)

	despedida.Despedirse(nombre)
}
