package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 //inferido

	area := PI * math.Pow(raio, 2) // := cria e inicializa a variavel
	fmt.Println(area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	g, h, i := 2, false, "opa"

	fmt.Println(e, f, g, h, i)
}
