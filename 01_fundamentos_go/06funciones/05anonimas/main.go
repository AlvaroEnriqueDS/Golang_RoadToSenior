package main

import "fmt"

func main() {
        //En las funciones anonimas

	//UNA VARIABLE PUEDE SER UNA FUNCION
	//a esto se le llama expresion funcion
	anonima := func(x int) {
		fmt.Println("Me imprimo estando en una variable llamada anonima, ", x)
	}
	anonima(2)

	//o

	//Se pueden ejecutar directamente
	func() {
		fmt.Println("Me imprimo estando en una variable llamada anonima")
	}()
	//tambien pueden recibir parametros
	nombre := "alvaro"
        func(x string) {
                fmt.Println("Imprimo mi parametro ", x )
        }(nombre)


	//RETORNAR UNA FUNCION
	//al darle una funcion a la variable lo que hacemos es retornar una funcion
	//para poder usarla se usa lavariable creada con ()
	//ahora la variable se comporta como la funcion retornada

	//podria ser
	variable := secuencia()
        fmt.Println(variable())
        fmt.Println(variable())
        fmt.Println("----------")

	//o ejecutamos diractamente lo que retorna
        fmt.Println(secuencia()())
        fmt.Println(secuencia()())
        fmt.Println("----------")


	//SE MANTIENE ACTIVA PARA SECUENCIA1
	secuencia1 := secuencia()
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())

	fmt.Println("----------")

	//SE MANTIENE ACTIVA PARA SECUENCIA2
	secuencia2 := secuencia()
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
}

//ESTE TIPO DE FUNCIONES SE MANTIENEN VIVAS PARA CADA OBJETO O VARIABLE INSTANCIADA
func secuencia() func() int32 {
	var x int32
	return func() int32 {
		x++
		return x
	}
}
