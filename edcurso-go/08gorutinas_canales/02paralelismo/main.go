package main

import (
	"fmt"
	"sync"
	"runtime"
)

func main() {
	//ESTE PAQUETE ES PARA USAR CIERTOS PROCESADORES
	runtime.GOMAXPROCS(4)

	//VARIABLE WAITGROUG
	var wg sync.WaitGroup

	numbers := []uint32{125424, 235, 1241877, 32135647, 6544774, 417485, 96547854, 11113131, 124578963, 11477741}
	//se agrega cuantos rutinas va a lanzar y espere
	wg.Add(len(numbers))

	//n := []uint32{1,1,1,1,1,1,1,1,1,1,1}

	fmt.Println("Comienza el proceso...")
	for _, v := range numbers {
		fmt.Println("////")
		//GO RUTINA CON FUNCION ANONIMA
		go func(a uint32){
			//esto quita las rutinas
			defer wg.Done()
			fmt.Println(a, EsPrimo(a))
		}(v) //aqui le estan pasando el valor que capta de numbers

	}

	//CON ESTO ESTA ESPERANDO ANTES DE TEMRINAR
	wg.Wait()
	fmt.Println("Terminó el proceso")
}

func EsPrimo(a uint32) bool {
	c := 0
	var i uint32
	for i = 1; i <= a; i++ {
		if a%i == 0 {
			c++
		}
	}
	if c == 2 {
		return true
	}
	return false
}