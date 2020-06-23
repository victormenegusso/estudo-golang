package main

import (
	"fmt"
	"reflect"
)

func main() {
	var notas [3]float64

	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 9.0, 10.0

	fmt.Println(notas)

	for i, nota := range notas {
		fmt.Printf("indice: %d nota: %f \n", i, nota)
	}

	// slice
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	// slice não é um array, ele é um pedaço da array
	s2 := a2[1:3] //

	fmt.Println(a2, s2)

	// slice com make

	sm := make([]int, 18)
	sm[9] = 17
	fmt.Println(sm)

	// tamanho do slice, tamanho do array interno
	sm = make([]int, 10, 20)
	fmt.Println(sm, len(sm), cap(sm))

	sm = append(sm, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(sm, len(sm), cap(sm))

	// var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[98765432100] = "Pedro"
	aprovados[95135745682] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[95135745682])
	delete(aprovados, 95135745682)
	fmt.Println(aprovados[95135745682])
}
