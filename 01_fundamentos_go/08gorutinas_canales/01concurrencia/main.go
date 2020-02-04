package main

import (
	"fmt"
	"time"
)

//TOMAR VARIAS PARTES E IDEPENDISARLAS Y QUE SE COMUNIQUEN
//PARA QUE SIMULEN QUE ESTAN TRABAJANDO AL MISMO TIEMPO
func main() {
	var h string
	//CON LA PALABRA "GO" estamos creando una gorutina
	//le decimos que ejecute esto de manera idependiente a lo demás
	go MostrarNumeros()
	fmt.Println("Digita algo")
	//SCAN PIDE QUE INGRESES UIN VALOR
	fmt.Scan(&h)
	fmt.Println("Digitaste", h)
	time.Sleep(time.Second * 10)
}

func MostrarNumeros() {
	//for clasico
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		//ESTE PAQUETE "TIME" Y AL FUNCION "SLEEP"
		//esto nos solicita que opcion tiene para detenerse
		time.Sleep(time.Second)
	}
}
