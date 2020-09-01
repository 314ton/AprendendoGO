package main

import (
	"fmt"
)


func soma_slice( parâmetro []int ) int {
	var total int
	for _, valor := range parâmetro {
		total += valor
	}
	return total
}


func soma_int( parâmetro ...int ) int {
	var total int
	for _, valor := range parâmetro {
		total += valor
	}
	return total
}

func main() {
	argumento := []int{0,1,2,3,4,5,6,7,8,9}
	
	tot_soma1 := soma_slice(argumento)
	tot_soma2 := soma_int(argumento...)
	
	fmt.Printf(
		"O total da primeira soma é: %d\n",
		tot_soma1,
	)
	fmt.Printf(
		"O total da segunda  soma é: %d\n",
		tot_soma2,
	)
}