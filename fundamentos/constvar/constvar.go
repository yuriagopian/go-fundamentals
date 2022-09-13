package main

import (
	"fmt"
	m "math" // podemos passar como alias o nome do import
)

func main() {
	const PI float64 = 3.1415

	// tipo (float64 definido por inferencia do compilador)
	var raio = 3.2

	// forma reduzida de criar var declara e atribui valor (recomendavel)
	// variavel e não constante
	area := PI * m.Pow(raio, 2)

	// posso reatribuir valores desta forma
	// area = 1

	// não posso reatribuir valores desta forma, pois estaria declarando duas vezes
	// area :=2

	fmt.Println("A area da circusferencia é", area)

	// Grupo de constantes
	const (
		a = 1
		b = 2
	)

	// Grupo de variaveis
	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	// Declarando e atribuindo valores com tipagem unica
	var e, f bool = true, false

	fmt.Println(e, f)

	// Declarando e inicializando por tipando inferencia, a tipagem não pode ser alterada
	g, h, i := 2, false, "Epa!"

	fmt.Println(g, h, i)
}
