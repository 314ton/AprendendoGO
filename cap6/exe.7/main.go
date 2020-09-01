package main

import (
	"fmt"
)

func main() {
	x := func()(int){
		return 10
	}
	fmt.Println(x())
	fmt.Printf("%T\n",x)
}