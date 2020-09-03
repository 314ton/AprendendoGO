package main

import "fmt"

func main() {
	var x = 10
	fmt.Println("Original")
	fmt.Printf("O endereço da variável X é %v\n",&x)
	fmt.Println("Decimal")
	fmt.Printf("O endereço da variável X é %d\n",&x)
	fmt.Println("Octal")
	fmt.Printf("O endereço da variável X é %o\n",&x)
	fmt.Println("Binario")
	fmt.Printf("O endereço da variável X é %b\n",&x)
}