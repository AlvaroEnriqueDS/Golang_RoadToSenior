package main

import "apis-go/middlewares/funciones"

func execute(name string, f funciones.MyFunction) {
	f(name)
}

func main() {
	name := "Comunidad EDteam"
	execute(name, funciones.MiddlewareLog(funciones.Saludar))
	execute(name, funciones.MiddlewareLog(funciones.Despedirse))
}