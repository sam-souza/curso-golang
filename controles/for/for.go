package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	for i <= 10 { // não tem while em go
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	for {
		// laço infinito
		fmt.Println("Para sempre...")
		time.Sleep(time.Second)
	}
}
