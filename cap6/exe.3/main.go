package main

import (
	"fmt"
)

func main() {
	for i := 0;i < 5; i++ {
		defer fmt.Println("primeiro valor",i)
		fmt.Println("segundo valor",i)
	}
	println()
}