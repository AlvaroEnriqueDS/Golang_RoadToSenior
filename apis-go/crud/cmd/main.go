package main

import (
	"apis-go/crud/authorization"
	"apis-go/crud/handler"
	"apis-go/crud/storage"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main()  {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	err = authorization.LoadFiles("crud/cmd/certificates/app.rsa", "crud/cmd/certificates/app.rsa.pub")
	if err != nil {
		log.Fatal("No se pudieron cargar los certificados: ", err)
	}

	store := storage.NewMongo()
	mux := http.NewServeMux()


	handler.RoutePerson(mux, &store)
	handler.RouteLogin(mux, &store)

	log.Println("Servidor iniciado en el puerto http://127.0.0.1:8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Println("Error en el servidor: %v\n", err)
	}
}