package main

import (
	"fmt"
	"net/http"
	"time"
)

func main()  {
	http.HandleFunc("/saludar", saludar)

	server := http.Server{
		Addr:              ":8080",
		Handler:           nil,
		TLSConfig:         nil,
		ReadTimeout:       0*time.Second,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0*time.Second,
		IdleTimeout:       0,
		MaxHeaderBytes:    1 << 20,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	server.ListenAndServe()
}

func saludar(w http.ResponseWriter, r *http.Request)  {

	//w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "Hola mundo")
}