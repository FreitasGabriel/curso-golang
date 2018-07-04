package main

import (
	"fmt"
)

func main() {
	fmt.Print("Mesma")
	fmt.Print("Linha")

	x := 3.1415

	fmt.Println()

	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é: " + xs)

	a := 1
	b := 1.9999
	c := false
	d := "teste"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %.1v %v %v", a, b, b, c, d)
}
