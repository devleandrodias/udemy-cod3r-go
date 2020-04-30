package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// Cuidado...
	fmt.Println("Teste " + string(97))

	// Int para String
	fmt.Println("Teste " + strconv.Itoa(123))

	// String para Inteiro
	num, _ := strconv.Atoi("123") // Retorna dois valores (número, erro)
	fmt.Println("Número:", num-23)

	// String para Booleano
	bo, _ := strconv.ParseBool("true")
	fmt.Println("Booleando:", bo)
	fmt.Println("Booleando:", reflect.TypeOf(bo))
}
