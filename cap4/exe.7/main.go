package main

import (
	"fmt"
)

func main(){
	slice := [][]string{
		[]string{
			"nome",
			"sobrenome",
			"hobby",
		},
		[]string{
			"etho",
			"martins",
			"programar",
		},
		[]string{
			"clovis",
			"martins",
			"mexer no carro",
		},
		[]string{
			"tião",
			"domingues",
			"tatuagem",
		},
	}
	for _,v := range slice{
		for l, val := range v{
			if l != (len(slice[1])-1){
				fmt.Printf("%s\t |\t ",val)
			} else {
				fmt.Printf("%s",val)
			}
		}
		println("")
	}
}

