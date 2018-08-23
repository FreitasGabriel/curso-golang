package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Println(i)

	}

	for {
		fmt.Println("Para sempre...")
		time.Sleep(time.Second * 2)
	}
}
