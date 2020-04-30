package main

import "fmt"

func main() {
	imprimirResultado(4.6)
	imprimirResultado(9.6)
	imprimirResultado(10.0)
}

func imprimirResultado(nota float64) {
	if nota < 7 {
		fmt.Println("Reprovado com nota:", nota)
	} else if nota >= 7 && nota <= 9.9 {
		fmt.Println("Aprovado com nota:", nota)
	} else {
		fmt.Println("Aprovado com execelêcia:", nota)
	}
}
