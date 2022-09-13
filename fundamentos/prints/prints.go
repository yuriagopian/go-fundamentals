package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516

	// Não funciona em go por conta da tipagem de x
	// fmt.Println("O valor de x é " + x)

	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é", x)

	fmt.Printf("O valor de x é %f", x)

	// Considerando duas casas decimais
	fmt.Printf("\nO valor de x é %.2f.", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	// Utilizar a letra certa de acordo com o tipo
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)

	// O v serve para praticamente todas e printa da forma correta
	fmt.Printf("\n%v %v %v %v", a, b, c, d)

}
