package main

import (
	"fmt"
	ma "math"
)

// Em go você não precisa passar tipo da váriavel, ele infere os tipos

func main() {
	const PI float64 = 3.14 // constante
	var raio = 3.2

	// Forma reduzida de criar uma váriavel!
	area := PI * ma.Pow(raio, 2)

	// Em go se você define uma váriavel e não utilize ele gera um erro de compilação

	fmt.Println("A área da cincurferência é:", area)

	const (
		a = 1
		b = 4
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false

	fmt.Println(e)
	fmt.Println(f)

	// Inferência de tipos
	g, h, i := 2, false, "epa!"

	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
}
