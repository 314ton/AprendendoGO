package main

import "fmt"

type Zint int
var x Zint


func main(){
	fmt.Printf("\nx_type -> %T \nx_val -> %v\n",x,x)
	x = 42
	fmt.Printf("x_val -> %v\n\n",x)
}
