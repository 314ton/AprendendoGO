package main

import (
	"os"
	"fmt"
	"encoding/json"
)

type qualquer struct {
	A1	string
	A2	int
}

func (q *qualquer) SetA1(x string) *qualquer {
	q.A1 = x
	return q
}
func (q *qualquer) SetA2(x int) *qualquer {
	q.A2 = x
	return q
}

func main() {
	b1 := qualquer{
		A1: "batata",
		A2: 12,
		
	}
	b2 := qualquer{
		A1: "pepsi",
		A2: 1,
	}
	
	b3 := new(qualquer)
	b3.SetA1("Etho").SetA2(13)
	
	b4 := new(qualquer)
	fmt.Println(
		"aaa",b4.
		SetA1("Jão").
		SetA2(31),
	)
	
	fmt.Println("b3->",*b3)
	fmt.Println("b4->",*b4)
	c := []qualquer{b1,b2,b1,b2}
	
	
	json.NewEncoder(os.Stdout).Encode(c)
}