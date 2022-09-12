// Programas executáveis iniciam pelo pacote main
package main

import "fmt"

/*
Os codigos em Go são organizados em pacotes
e para usá-los é necessarios declarar um ou vários imports
*/
func main() {
	// inline
	fmt.Print("Primeiro ")
	fmt.Print("Programa!")
	// break-line
	fmt.Println("")
	fmt.Println("Primeiro ")
	fmt.Println("Programa!")
	/*
		Sobre comentários
		1) Priorize código legivel e faça comentarios que são necessarios para o entendimento
		ou para fins de alertar algo
		2) Evite comentários obvios
	*/
}
