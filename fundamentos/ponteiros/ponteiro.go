package main

import "fmt"

func main() {
	/*
	   64 bits ou 8 bytes na memória, pode-se compartilhar endereço dela na memória
	*/

	// Ponteiro é uma referência de memória
	i := 1 // Go não tem aritiméticas de ponteiros

	var p *int = nil // nil == null

	// Pegar endereço de memória da váriavel e traibuir ao ponteiro
	// pegando o endereço da váriavel
	p = &i // p agora guarda o endereço de memória que aponta para mesmo local que i aponta

	*p++ /*Desreferênciando o ponteiro*/ // valor para o que o endereço aponta,

	// endereço de memória que ponteiro aponta
	i++

	// p++ (Go não permite aritimética de ponteiro)
	fmt.Println(p, *p, i, &i)
}
