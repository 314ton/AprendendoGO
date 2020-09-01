package main

import "fmt"

func main() {
	fmt.Println(prim(mais1,5))
}

func prim( f func(x int) int, x int) int {
	x++
	println(x)
	return f(x)
}
func mais1(x int) int {
	return x+1
}