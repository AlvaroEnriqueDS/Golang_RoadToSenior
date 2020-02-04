package main

import (
	"fmt"
	"strings"
)

func main() {
	//strings
	text := "HELLO! MY NAME IS ALVARO, IAM FROM PeRuVian, I LIVE IN LIMA. HELLO!"

	fmt.Println(strings.ToUpper(text))
	fmt.Println(strings.ToLower(text))
	fmt.Println(strings.Replace(text, "HELLO", "HOLA", -1))
}
