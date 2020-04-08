package main

import "fmt"

func inc1(n int) {
	n++
}

func inc2(n *int) {
	*n++
}

func main() {
	x := 1

	inc1(x)
	fmt.Println(x)

	inc2(&x)
	fmt.Println(x)
}
