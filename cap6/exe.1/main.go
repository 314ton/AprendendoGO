package main


import (
	"fmt"
)

func dobro(x int) int {
	return x*2
}

func triplo( x int ) ( int, string ) {
	return x*3, fmt.Sprintf("o triplo de %d é %d",x,x*3)
}

func main() {
	fmt.Println(dobro(8))
	x,y := triplo(8)
	fmt.Println(x,y)
}