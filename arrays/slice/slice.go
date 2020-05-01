package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice

	fmt.Println(a1)
	fmt.Println(s1)

	fmt.Println(reflect.TypeOf(a1))
	fmt.Println(reflect.TypeOf(s1))

	// slice não é um array, slice define um pedaço de um array
	a2 := [5]int{1, 2, 3, 4, 5}

	s2 := a2[1:3] // index 1 até 3 não incluindo o 3

	fmt.Println(s2) // ele não cria um array cria um pedaçõ do original, ponteiro

	s3 := a2[:2] // novo slice, mas aponta para o mesmo array
	fmt.Println(a2, s3)

	// você pode imaginar um slice como: tamanho e um ponteiro para um elemento de um array
	s4 := s2[:1] // slice de um slice
	fmt.Println(s4)
}
