package main

import "fmt"

func main(){
	for c:=65; c<=90; c++{
		fmt.Printf("%v:\n",c)
		for i:=0;i<3;i++{
			fmt.Printf("\t%#U -> %c\n",c,c)
		}
		fmt.Println("")
	}
}
