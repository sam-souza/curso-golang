package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2

	// := forma reduzida de criar e inicializar a variavel
	area := PI * math.Pow(raio, 2)
	fmt.Println("A área da circunferência é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = false, true
	fmt.Println(e, f)

	g, h, i := false, "String", 4.56
	fmt.Println(g, h, i)
}
