package main

import "fmt"

func main() {
	x := 1
	y := 2

	// apenas postfix
	x++ // x = x + 1
	fmt.Println(x)

	// não existe ++x

	y--
	fmt.Println(y)

	// fmt.Println(x == y--) não permitido, misturar na expressão
}
