package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")
	fmt.Println("esse virá na próxima linha.")

	x := 3.141516
	xs := fmt.Sprint(x)

	fmt.Println("O valor de x é: ", x)
	fmt.Print("O valor de x é: " + xs)
	fmt.Printf("O valor de x é: %.2f", x) // Imprimir valor formatado

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	/*
		\n pular linha
		%d tipo inteiro
		%f tipo decimal
		%t tipo boleano
		%s tipo string
		%v tipo genérico
		%.1f tipo decimal formatado
	*/

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v %v", a, b, b, c, d)
}
