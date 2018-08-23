package main

import (
	"fmt"
)

func main() {
	numeros := [...]int{1, 2, 3, 4, 5}

	for i, numero := range numeros { //o primeiro elemento do for (i) é o indice
		fmt.Printf("%d) %d\n", i+1, numero)

	}

	for _, num := range numeros { //o segundo elemento do for (num) é o valor da posição no array
		fmt.Println(num)
	}
}
