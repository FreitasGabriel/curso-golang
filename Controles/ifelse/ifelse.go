package main

import "fmt"

func imprimirresultado(nota float64){
	if nota >= 7 {
		fmt.Println("Aprovado cojm nota ", nota)
	}else{
		fmt.Println("Reprovado com nota ", nota)
	}
}

func main(){
	imprimirresultado(8)
	imprimirresultado(2)
}