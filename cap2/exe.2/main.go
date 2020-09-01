package main

import "fmt"


func main(){
	a := (1 == 1)
	b := ("etho" != "etho")
	c := (15 > 14)
	d := (14 >= 15)
	e := (7 < 8)
	f := (8 <= 7)

	fmt.Printf(
		"a é %v, b é %v, c é %v, d é %v, e é %v e f é %v\n",
		a,b,c,d,e,f,
	)
}
