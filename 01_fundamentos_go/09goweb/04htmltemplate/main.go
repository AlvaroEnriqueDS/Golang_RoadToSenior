package main

import (
	"html/template"
	"log"
	"net/http"
)

// Persona contiene los datos de una persona
type Persona struct {
	Nombre string
	Edad   uint8
}

//una hanldefunc
func renderizar(w http.ResponseWriter, r *http.Request) {
	//asignamos valores a la estructura persona
	p := &Persona{"Daniel", 26}
	//vamos a analizar archivos o renderizado
	//te devuelve el tample y el error
	t, err := template.ParseFiles("./views/index.html")
	//si tiene un error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	//sitodo salio bien le dedimos al template que se eejecute
	// se le puede pasar nil en vez de p
	t.Execute(w, p)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", renderizar)

	log.Println("Ejecutando server en http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", mux))
	log.Println("Ejecución terminada, ya no está corriendo el servidor")
}
