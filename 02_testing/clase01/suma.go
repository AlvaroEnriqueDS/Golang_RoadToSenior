package clase01

func Add(a, b int) int {
        return a+b
}

func AddAcum(numbers ...int) int {
        resp := 0
        for _, n := range numbers {
                resp += n
        }
        return resp
}
