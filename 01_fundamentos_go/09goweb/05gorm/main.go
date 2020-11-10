package main

import (
	"fmt"

	//go get -u github.com/jinzhu/gorm
	//go get -u github.com/lib/pq
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Producto contiene los datos de un artículo
type Producto struct {
	gorm.Model   //ID, CreatedAt, UpdatedAt, DeletedAt
	CodigoBarras string
	Precio       uint
}

func main() {
	db, err := gorm.Open("postgres", "user=postgres password=asd020 dbname=lacoraza sslmode=disable")
	if err != nil {
		panic("Error en la conexión a la base de datos")
	}
	defer db.Close()
	fmt.Println("Se conectó a la base de datos")

	/*
		01_db.CreateTable(&Producto{})
		fmt.Println("Se creó la tabla productos en la base de datos")
	*/

	/*
		01_db.Create(&Producto{
			CodigoBarras: "otroCodigoDebarras",
			Precio:       5000,
		})

	*/
	/*
		var p Producto
		01_db.First(&p, 2)
		fmt.Println("Precio del producto consultado:", p.Precio)

		p.Precio = 4500
		01_db.Save(&p)
		fmt.Println("El nuevo precio del producto es", p.Precio)
	*/
}
