package main

import "fmt"

func main() {
	imprimirResultado(4.6)
	imprimirResultado(9.6)
}

func imprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota:", nota)
	} else {
		fmt.Println("Reprovado com nota:", nota)
	}
}
