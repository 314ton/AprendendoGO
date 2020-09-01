package main

import (
	"fmt"
)


func main(){
	slice := []int{
		41, 42, 43, 44, 45,
		46, 47, 48, 49, 51,
	}
	
	slice = append(slice, 52)
	fmt.Printf("Slice: -> %v\n", slice)
	
	slice = append(slice, 53, 54, 55)
	fmt.Printf("Slice: -> %v\n", slice)
	
	slice2 := []int{
		56, 57, 58, 59, 60,
	}
	slice = append(slice, slice2...)
	fmt.Printf("Slice: -> %v\n", slice)
}