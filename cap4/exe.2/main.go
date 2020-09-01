package main


import (
	"fmt"
)

func main(){
	
	slice := []int{
		10,20,30,40,50,60,70,80,90,
	}
	for _, val := range slice{
		fmt.Printf("type:-> %T val:-> %v\n",val,val)
	}
	fmt.Printf("slice type:-> %T\n",slice)
	
}