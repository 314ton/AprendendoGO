package main

import "fmt"

func iotario() (func() int) {
	x := 0
	return func() int {
		x++ 
		return x
	}
}


func main() {
	a := iotario()
	b := iotario()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	
}