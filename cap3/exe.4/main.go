package main

import "fmt"

var ano = 2001


func main(){
	for{
		fmt.Println(ano)
		if ano >= 2020{ break }
		ano++
	}
}
