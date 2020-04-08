package main

import (
	"fmt"
	// "time"
)

type structTeste struct {
	a string
}

func main() {
	// fmt.Print("Mesma")
	// fmt.Print(" linha.")

	// fmt.Println(" Nova")
	// fmt.Println("linha.")

	// x := 3.141516

	// // fmt.Println("O valor de x é " + x)
	// xs := fmt.Sprint(x)
	// fmt.Println("O valor de x é " + xs)
	// fmt.Println("O valor de x é", x, "!!!!!")

	// fmt.Printf("O valor de x é %.2f.", x)

	// a := 1
	// b := 1.9999
	// c := false
	// d := "opa"
	// fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	// fmt.Printf("\n%v %v %v %v", a, b, c, d)

	// teste := "2019-12-28"

	// matched, err := time.Parse("2006-01-02", teste)
	// fmt.Println(matched, err)

	// fmt.Println(time.Now().Sub(matched).Hours() / 24)

	// fmt.Printf(time.Now().Format("2006-01-02"))

	// a := structTeste{}

	// if a.a == "" {
	// 	fmt.Printf("Teste")
	// }

	// teste := make(map[string]struct{}, 2)

	// teste["t"] = struct{}{}

	// _, ok := teste["b"]

	// fmt.Println(ok)

	// teste := "2019-12-28"

	// matched, err := time.Parse("2006-01-02", teste)
	// fmt.Println(matched, err)

	// pattern := "January 2 15:04:05 2006 MST"

	// t := matched.Format(pattern)

	// fmt.Println(t)

	type Month int

	const (
		January Month = iota + 1
		February
		March
		April
		May
		June
		July
		August
		September
		October
		November
		December
	)

	fmt.Printf("%d", January)
}
