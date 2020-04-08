package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"A": {
			"Andre":    123.0,
			"Ana":      91279.4,
			"Anderson": 456456.7,
		},
		"B": {
			"Bruno": 1234.5,
		},
		"C": {
			"Carla": 24564.4,
		},
	}

	fmt.Println(funcsPorLetra)

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra)
		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}
	}
}
