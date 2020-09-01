package main

import (
	"fmt"
)

func main(){
	
	slice_x := []int{
		42, 43, 44, 45, 46,
		47, 48, 48, 49, 50, 51}
	slice_y := append(
		slice_x[:3],
		slice_x[7:]...,
	)
	fmt.Println(slice_y)
}

