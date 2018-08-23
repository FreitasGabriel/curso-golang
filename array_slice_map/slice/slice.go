package main

import (
	"fmt"
)

func main() {
	a1 := [3]int{1, 2, 3} //ARRAY
	s1 := []int{4, 5, 6}  //SLICE

	fmt.Println(a1, s1)

	a2 := [5]int{1, 2, 3, 4, 5}

	//slice não é um array. Slice define um pedaço de um array

	s2 := a2[1:3]

	fmt.Println(a2, s2)

	s3 := a2[1:4]

	fmt.Println(a2, s3)

}
