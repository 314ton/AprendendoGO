package main

import "fmt"

var (
	x, y, z = "James Bond", 42, false
)


func main(){
	s := fmt.Sprintf("Nome: %v, Idade: %v , Casado: %v",x,y,z)
	fmt.Println(s)
}
