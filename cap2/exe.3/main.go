package main

import "fmt"

const (
	nome = "Etho"
	idade = 19
)
const (
	casado bool = false
	altura float64 = 1.75
)
func main(){
	fmt.Printf(
		"\n\tNome: %v,\n\tIdade: %v,\n\tcasado: %v,\n\taltura: %v\n\n",
		nome, idade,
		casado, altura,
	)
}
