package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José":       123.0,
		"Marcio":     1513.0,
		"Alessandra": 1354.0,
	}

	funcsESalarios["Marcela"] = 123.6

	fmt.Println(funcsESalarios)

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
