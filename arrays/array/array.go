package main

import "fmt"

// homogênea (mesmo tipo) e estática (fixo), estrutura indexada
func main() {
	var notas [3]float64

	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 3.5, 8.6

	total := 0.0

	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("Média %.2f\n", media)
}
