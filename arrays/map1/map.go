package main

import "fmt"

func main() {
	//var aprovados = map[int]string // mapas devem ser inicializados, [chave] valor
	aprovados := make(map[int]string)

	aprovados[12345678911] = "Maria"
	aprovados[42237642873] = "João"
	aprovados[34598347548] = "Thaísa"

	// se você passar duas chaves iguais permanece o último valor

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[34598347548])
	delete(aprovados, 12345678911)
	fmt.Println(aprovados)
}
