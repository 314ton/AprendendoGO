package main

import "fmt"

type Zint int
//type Zint2 Zint

var (
	x Zint
	// y Zint2
	y int
)


func main(){
	fmt.Printf("\nx_type -> %T \nx_val -> %v\n",x,x)
	x = 42
	fmt.Printf("x_val -> %v\n\n",x)
	y = int(x)
	fmt.Printf("y_type -> %T\ny_val -> %v\n",y,y)
}
