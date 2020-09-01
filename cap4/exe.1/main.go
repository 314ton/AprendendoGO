package main


import "fmt"

func main(){
	array := [5]int{9, 8, 7, 6, 5 }
	for index, val := range array{
		fmt.Printf(
			"Val: %v \tPos: %v \tTipo: %T \n",
			val, index, val,
			)
	}
}
