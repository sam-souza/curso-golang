package main

import "fmt"

func main() {
	// var aprovados map[int]string     <--- Não funciona. Maps devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[123] = "Maria"
	aprovados[456] = "Ana"
	aprovados[789] = "Pedro"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[123])
	delete(aprovados, 123)
	fmt.Println(aprovados[123])
}
