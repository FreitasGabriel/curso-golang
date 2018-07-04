package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//numeros inteiros
	fmt.Println(1, 2, 1000)

	fmt.Println("Literal inteiro é: ", reflect.TypeOf(32000))

	//sem sinal (só positivos)... unit8 unit16 unit32 unit64
	var b byte = 255
	fmt.Println("0 byte é: ", reflect.TypeOf(b))

	fmt.Println("============INTEIROS==============")

	//com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valro máximo do int é: ", i1)
	fmt.Println("O tipo de i1 é: ", reflect.TypeOf(i1))

	var i2 rune = 'a' //representa um mapeamento da tabela Unicode (int32)
	fmt.Println("o rune é: ", reflect.TypeOf(i2))
	fmt.Println(i2)

	fmt.Println("============REAIS==============")

	//numeros reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é: ", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é: ", reflect.TypeOf(49.99))

	fmt.Println("============BOOLEAN==============")
	//boolean
	bo := true
	fmt.Println("O tipo de bo é: ", reflect.TypeOf(bo))
	fmt.Println(!bo)

	fmt.Println("============STRING==============")

	//string
	s1 := "Olá meu nome é Gabriel"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é: ", len(s1))

	//string com multiplas linhas

	s2 := `Olá
	meu 
	nome 
	é 
	gabriel`

	fmt.Println("O tamanho da string é: ", len(s2))

	fmt.Println("============CHAR==============")
	//char??
	//NÃO EXISTE O TIPO CHAR
	char := 'a'
	fmt.Println("O tipo de char é: ", reflect.TypeOf(char))
}
