package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415

	// tipo (float64 definido por inferencia do compilador)
	var raio = 3.2

	// forma reduzida de criar var declara e atribui valor
	area := PI * math.Pow(raio, 2)
	fmt.Println("A area da circusferencia é", area)
}
