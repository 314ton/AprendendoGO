package main

import (
	"fmt"
)

type pessoa struct {
	nome	string
	email	string
}
func (p pessoa) infos() {
	fmt.Printf("Nome: %s\n", p.nome)
	fmt.Printf("Email: %s\n", p.email)
}

func main() {
	etho := pessoa{
		nome:	"etho",
		email:	"etho@etho.etho",
	}
	etho.infos()
}