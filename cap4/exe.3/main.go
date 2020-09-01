package main

import (
	"fmt"
)

func main(){
	
	slice := []int{11,13,15,17,19,21,23,25,27,29}
	fmt.Println(slice[:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:7])
	fmt.Println(slice[3:9])
	fmt.Println(slice[3:len(slice)-1])
	
}