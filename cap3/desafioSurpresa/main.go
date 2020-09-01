package main

import "fmt"


func main(){
	for i := 33; i<= 122; i++{
		if i%2 == 0{
			fmt.Printf("\033[92mi-> %v == \033[m%c\n",i,i)
		}else{
			fmt.Printf("\033[94mi-> %v == \033[m%c\n",i,i)
		}
	}
	//fmt.Println("\033[m")
}
