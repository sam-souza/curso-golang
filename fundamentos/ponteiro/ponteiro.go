package main

import "fmt"

func main() {
	i := 1

	// Go não tem aritmética de ponteiros. Ex: p++
	var p *int = nil
	p = &i
	fmt.Println("Endereço de memória i é:", p)
	fmt.Println("Acessando o valor de i por p:", *p)

	*p++ // atualizando o valor de i por p
	fmt.Println("Acessando o valor de i por p atualizado:", *p)
	fmt.Println("Acessando o valor de i por i atualizado:", i)
}
