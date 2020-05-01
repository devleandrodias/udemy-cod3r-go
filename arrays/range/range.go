package main

import "fmt"

func main() {
	// Você tem um array de tamanho fixo de tamnho de elementos que você passar na inicialização
	numeros := [...]int{5, 3, 1, 3, 26}

	// i = indicie atual, element é o elemento que estou percorrendo no array, numeros é o array
	for i, element := range numeros {
		fmt.Printf("[%d] %d\n", i, element)
	}

	for _, num := range numeros {
		fmt.Println(num)
	}
}
