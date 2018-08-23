package main

import (
	"fmt"
)

func main() {
	//homogêneo (mersmo tipo) e estática (fixo)

	var notas [3]float64

	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 8.8, 4.5, 7

	total := 0.0

	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))

	fmt.Printf("Média %.2f\n", media)
}
