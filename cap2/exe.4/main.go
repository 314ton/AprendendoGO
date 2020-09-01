package main

import "fmt"

func main(){
	x := 123
	fmt.Printf(
		"x ->\nBin %#b \nDec %#d\nHexa %#x \n",
		x,x,x,
	)
	y := x << 1
	fmt.Printf(
		"y ->\nBin %#b \nDec %#d \nHexa %#x\n",
		y,y,y,
	)

}
