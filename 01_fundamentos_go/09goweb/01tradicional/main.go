package main

import (
	"log"
	"net/http"
)

func main() {
	//asi se crea uin servidor para una web estatica servida desde go
	/*
		//RUTA DE DONDE SERVIRSE
		http.Handle("/", http.FileServer(http.Dir("public")))
		log.Println("Ejecutando server en http://localhost:8080")
		//ESTE LE DICE POR QUE TARJETA Y PUERTO VA A SERVIRSE
		log.Println(http.ListenAndServe(":8085", nil))
	*/

	//UN MULTIPLEXOR ES PARA MEJORAR LA EFICIENCIA DE MI SERVIDOR
	//mux es un multiplexor
	mux := http.NewServeMux()
	//guardamos la ruta de donde se va a servir
	fs := http.FileServer(http.Dir("public"))
	//EL HANDLE ES UN MANEJADOR Y SE LE PASA LA RUTA Y EL FILESERVER
	mux.Handle("/", fs)
	log.Println("Ejecutando server en http://localhost:8086")
	log.Println(http.ListenAndServe(":8086", mux))

}

/*
package main

import (
	"net/http"
	"io"
)

funchandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hola Servidor Go!")
}

funcmain() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/ii", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Route: ii!")
	})
	http.ListenAndServe(":8000", nil)
}
*/
