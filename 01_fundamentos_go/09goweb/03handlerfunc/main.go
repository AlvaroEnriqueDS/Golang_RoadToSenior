package main

import (
	"net/http"
	"fmt"
	"log"
)

//esta firma es para el handlerfunc
func messageHandlerFunc(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "<h1>Hola este es un handlerfunc</h1>")
}

//esta funcion me devuelve un mensaje en un handlerfunc
func messageHFCustom(message string) http.Handler {
	//vamos a retornar un hanlderfunc
	return http.HandlerFunc(
		//una funcion anonima que tiene la firma requerida
		func(w http.ResponseWriter, r *http.Request) {
			//escribe en writer y el message
			fmt.Fprintf(w, message)
		},
	)
}


func main() {
	//multiplexor
	mux := http.NewServeMux()
	//handlefunc nos permite meter una funcion, se le pasa la ruta y la funcion
	mux.HandleFunc("/", messageHandlerFunc)
	//aca solo se usa handle, pues la funcion persdonalizada no tiene la firma de handlefunc
	mux.Handle("/saludar", messageHFCustom("<h1>Hola amigos</h1>"))
	mux.Handle("/despedirse", messageHFCustom("<h1>Chao amigos</h1>"))

	log.Println("Ejecutando server en http://localhost:8080")
	log.Println(http.Client{
                Transport: http.RoundTripper(),
                CheckRedirect: nil,
                Jar:           nil,
                Timeout:       0,
        })
	log.Println(http.ListenAndServe(":8080", mux))
}
