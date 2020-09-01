package main

import "fmt"


func main(){
	var x int
	if fmt.Scan(&x); x > 10{
		fmt.Println("X não é maior a 10!",x)
	} else if x < 10{
		fmt.Println("X não é menor que 10!",x)
	} else {
		fmt.Println("X é igual a 10!",x)
	}
}
