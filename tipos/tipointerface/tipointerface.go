package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	type dinamica interface{}

	var coisa2 dinamica = "Opa"
	fmt.Println(coisa2)
}
