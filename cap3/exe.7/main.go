package main

import "fmt"

func main(){
	a := true
	switch {
	case a == true:
		fmt.Println("a é igual a True")
	default:
		fmt.Println("a é igual a False")
	}
	
}