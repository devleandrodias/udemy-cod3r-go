package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// numéricos inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é:", reflect.TypeOf(32000))
	fmt.Println("Literal inteiro é:", reflect.TypeOf(-32000))

	// sem sinal (só positivos) uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é:", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("O valor máximo de int é:", i1)
	fmt.Println("O valor máximo de int é:", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é:", reflect.TypeOf(i2))
	fmt.Println("O rune é:", i2)

	// número reais (float32, float64)
	var x float32 = 49.99
	// var y = float32(49.99)
	fmt.Println("O tipo de x é:", reflect.TypeOf(x))
	fmt.Println("O do literar 49.99 é:", reflect.TypeOf(49.99)) // Padrão float64

	// booleano
	bo := true
	fmt.Println("O tipo de bo é:", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Olá meu nome é Grood"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho de s1 é:", len(s1))

	// string com multiplas linhas
	s2 := `Olá meu nome
  é
  Grood`

	fmt.Println(s2)

	// char ?? (Não existe)
	char := 'a'
	fmt.Println(char)
	fmt.Println("Tipo de char?:", reflect.TypeOf(char))
}
