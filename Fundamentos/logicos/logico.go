package main

import (
	"fmt"
)

func comprar(trab1, trab2 bool) (bool, bool, bool) {

	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2 // ou esxclusivo
	comprarSoverte := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSoverte

}

func main() {
	tv50, tv32, sorvete := comprar(true, true)
	fmt.Printf("TV50: %t, Tv32: %t, Sorvete: %t, Saudável:  %t \n", tv50, tv32, sorvete, !sorvete)
}
