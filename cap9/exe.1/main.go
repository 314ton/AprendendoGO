package main

import (
	"fmt"
//	"sync"
)

//var wt sync.WaitGroup


func main() {
	//wt.Add(2)
	
	a := fib(5)
	fmt.Println(a)
	
	//wt.Wait()
}

func fib(x int) int {
	if x == 1{
	//	wt.Done()
		return 1
	} else {
		return x * fib(x-1)
	}
}
