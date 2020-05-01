package main

import "fmt"

func main() {
	fmt.Println(tipo(2.4))
	fmt.Println(tipo("Bom dia..."))
	fmt.Println(tipo(1))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(true))
}

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "não reconhecido"
	}
}
