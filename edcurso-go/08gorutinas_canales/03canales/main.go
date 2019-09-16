package main

import (
	"fmt"
	"time"
)
//LOS CANALES PERMITE COMUNICAR ENTRE SI A LAS GO RUTINAS
func main() {
	bodegaOrigen := []string{"php", "javascript", "html", "css", "java", "bases de datos", "git"}
	bodegaDestino := []string{}

	//LOS CANALES SE DECLARAN CON MAKE
	//ESTE CANAL SE USA CHAN (CANAL) Y EL TIPO
	fotocopiadora := make(chan string)
	fotocopiado := make(chan string)


	go func() {
		for _, libro := range bodegaOrigen {
			fotocopiadora <- libro
		}
	}()

	// Se deja en la bodega destino
	go func() {
		for {
			// Se entrega el contenido de la fotocopiadora a la variable libro
			libro := <-fotocopiadora

			fotocopiado <- libro

			// Agregar el libro a la bodega destino
			bodegaDestino = append(bodegaDestino, libro)

			//SI LLEGAN A SER IGUALES ENTONCES VAN A CERRAR EL CANAL
			if len(bodegaOrigen) == len(bodegaDestino) {
				// Cerrar el canal fotocopiado
				close(fotocopiado)
				fotocopiado <- "CTMREEEEEEEEE"
				break
			}

		}
		fmt.Println("TERMINE EN EL FOR INFITITO")
	}()

	fmt.Println("** Listado de libros fotocopiados **")
	for {
		//el canal esta retornando el string y el true o false
		if libro, ok := <-fotocopiado; ok {
			fmt.Println(libro)
		} else {
			break
		}
	}
	//ESTO ES SOLO PARA LA PRUEBA DONDE NOS ASEGURAMOS QUE TEMRINA EL PRIMER FOR INFINITO
	time.Sleep(time.Second*5)
}

/////////////////////////////////////7
/*
package main

import"fmt"

func main () {
	ch := make(chanstring)

	gofunc () {
		deferclose(ch)
		ch <- "Hola Channel"
	}()

	fmt.Println(<- ch)

	ch1 := make(chanint)

	gofunc() {
		deferclose(ch1)
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
		ch1 <- 4
		ch1 <- 5
	}()

	for n := range ch1 {
		fmt.Printf("%d\n", n)
	}

	ch2 := make(chanint, 2)
	ch2 <- 1
	ch2 <- 2
	//ch2 <- 3
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
	ch2 <- 3
	fmt.Println(<-ch2)
}
 */