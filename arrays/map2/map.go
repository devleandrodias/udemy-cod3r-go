package main

import "fmt"

func main() {
	// inicializar um map com valores literais
	funcsESalarios := map[string]float64{
		"José João":      123124.45,
		"Gabriela Silva": 3424325.23,
		"Pedro Júnior":   24234.52,
	}

	funcsESalarios["Rafael Filho"] = 243425.32
	delete(funcsESalarios, "Inexistente...") // não tem problema excluir algo que não existe

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
