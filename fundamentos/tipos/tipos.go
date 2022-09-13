package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	// Numeros inteiros
	fmt.Println(1, 2, 2000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(3200))

	// Sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo do i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' /// representa mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println("i2", i2)

	// numeros reais (float32, float64)
	var x float32 = 49.99

	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))

	// Conversão já que por padrão o literal é floar64
	var xy = float32(49.99)
	fmt.Println("O tipo de xy é", reflect.TypeOf(xy))

	// boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "olá meu nome é Yuri"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	// string com multiplas linha com backtick

	s2 := `Olá
	meu 
	nome 
	é
	Yuri`

	fmt.Println("O tamanho da string é", len(s2))

	// char??
	// var x char = 'b' -> Não existe em go

	char := 'a'  // pos na tabela unicode
	char2 := "b" // string
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println("O tipo de char2 é", reflect.TypeOf(char2))
	fmt.Println(char)
	fmt.Println(char2)

}
