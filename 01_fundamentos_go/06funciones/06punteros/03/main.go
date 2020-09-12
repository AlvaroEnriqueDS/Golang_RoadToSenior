package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
        //declaramos variables
	var order int = 0
	var file bool = true
	fmt.Println(order,&order)
        fmt.Println(file,&file)
        fmt.Println("///////////")
	prueba1(&order, &file)

        fmt.Println(order,&order)
        fmt.Println(file,&file)
        fmt.Println("///////////")
        prueba1(&order, &file)
}

func prueba1(order *int, file *bool) {
        fmt.Println(*order,order,&order)
        fmt.Println(*file,file,&file)
        fmt.Println("///////////")
        prueba2(order,file)
}

func prueba2(order *int, file *bool) {
        if *file {
                order = &[]int{5}[0]
                fmt.Println(&[]int{5}[0])
                fmt.Println(&[]int{5}[0])
                fmt.Println(order)

                x := &[]int{5}[0]
                *order = *x
                fmt.Println(order)


                *file = false
                fmt.Println(*order,order,&order)
                fmt.Println(*file,file,&file)
                fmt.Println("FILEEEEEEEEEEEEEE true")
        }else {
                if !*file{
                        //*order = 6
                        //*file = false
                        fmt.Println(*order,order,&order)
                        fmt.Println(*file,file,&file)
                        fmt.Println("FILEEEEEEEEEEEEEE false")
                }
        }
}
