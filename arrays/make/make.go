package main

import "fmt"

func main() {
	s := make([]int, 10) // crie um slice do tipo inteiro de 10 posições
	s[4] = 12
	fmt.Println(s)

	s = make([]int, 10, 20)        // também é um elemento de inteiro, só que o array interno vai ter 20 elementos
	fmt.Println(s, len(s), cap(s)) // cap é capacidade

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(s, len(s), cap(s)) // cap é capacidade

	s = append(s, 1, 2, 3)
	fmt.Println(s, len(s), cap(s)) // se passar a capacidade máxima ele dobra tamanho do array interno
}
