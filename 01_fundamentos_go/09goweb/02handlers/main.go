package main

import (
	"net/http"
	"fmt"
	"log"
)
//un handler es una interface que contiene un metodo que recibe como parametro
//un response writer(devolver informacion) y tambien recibe un request (solicitud del usuario)
type messageHandlerPrueba struct {
	message string
}

//metodo de messaeHandler
//es un metodo de la interface Handler y se llama ServeHTTP
//recibe como parametro el repsonse (para devolver) y el request (peticion, es un puntero)
func (m messageHandlerPrueba) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//UN MENSAJE CON EL WRITER "W", Y LO QUE VENGA EN EL MENSAJE.
	fmt.Fprintf(w, m.message)
}

func main() {
	mux := http.NewServeMux()
	m1 := &messageHandlerPrueba{message: "<h1>Hola desde un handler</h1>"}
	m2 := &messageHandlerPrueba{message: "<h1>Estos son Perritos</h1>"}
	m3 := &messageHandlerPrueba{message: "<h1>Estos son gatitos</h1>"}

	// / es la ruta
	mux.Handle("/saludar", m1)
	mux.Handle("/perros", m2)
	mux.Handle("/gatos", m3)

	log.Println("Ejecutando server en http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", mux))
}