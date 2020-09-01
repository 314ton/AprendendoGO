package main 

import (
	"fmt"
)

type pessoa struct{
	nome		string
	sobrenome	string
	idade		int 
	sorvete		[]string
}
func main(){
	
	etho := pessoa{
		nome: "etho",
		sobrenome: "martins",
		idade: 19,
		sorvete: []string{
			"napolitano",
			"passas ao rum",
			"abacaxi com vinho",
			"ninho",
		},
	}
	etho2 := pessoa{
		nome: "etho2",
		sobrenome: "martins2",
		idade: 21,
		sorvete: []string{
			"napolithanos",
			"passas ao ruim",
			"abacaxi calvinho",
		},
	}
	etho3 := map[string]pessoa{
		"martins3":pessoa{
			"etho",
			"martins3",
			19,
			[]string{
				"jaca",
				"tomate",
				"saguadin",
			},
		},
	}
	fmt.Println("-------------------")
	fmt.Printf("Nome: %s\n",etho.nome)
	fmt.Printf("Sobrenome: %s\n",etho.sobrenome)
	fmt.Printf("Idade: %d\n",etho.idade)
	fmt.Println("Sabores de Sorvete")
	for _, v := range (etho.sorvete){
		fmt.Printf("->\t%s\n",v)
	}
	fmt.Println("-------------------")
	fmt.Printf("Nome: %s\n",etho2.nome)
	fmt.Printf("Sobrenome: %s\n",etho2.sobrenome)
	fmt.Printf("Idade: %d\n",etho2.idade)
	fmt.Println("Sabores de Sorvete")
	for _, v := range (etho2.sorvete){
		fmt.Printf("->\t%s\n",v)
	}
	for _,val := range etho3{
		fmt.Println(val)
		for _,v := range val.sorvete{
			println(v)
		}
	}
}