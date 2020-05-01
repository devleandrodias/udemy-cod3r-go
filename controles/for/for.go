package main

import (
	"fmt"
	"time"
)

// Não tem while em go
func main() {
	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for j := 1; j <= 20; j++ {
		fmt.Println(j)
	}

	for k := 1; k <= 10; k++ {
		if k%2 == 0 {
			fmt.Print("Par")
		} else {
			fmt.Print("Impar")
		}
	}

	for {
		// laço infinito
		fmt.Println("Pra sempre...")
		time.Sleep(time.Second)
	}
}
