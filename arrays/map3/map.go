package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 12312.53,
			"Guga Pereira":   3454.23,
		},
		"J": {
			"José João": 34534.34,
		},
		"L": {
			"Leandro Dias": 23423523.23,
		},
	}

	for letra, funct := range funcsPorLetra {
		fmt.Println(letra, funct)
	}
}
