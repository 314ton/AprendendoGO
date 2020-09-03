package main

import (
	"fmt"
	"sort"
)

type seila struct {
	a1	string
	a2	int
	a3	float64
}
func ajeitaAi(s []seila){
	fmt.Println("\nSTRING \t\t| INT \t\t| FLOAT64\n")
	for i,v := range s {
		fmt.Println(i," a1:",v.a1,"\t| a2:",v.a2,"\t| a3:",v.a3)
	}
}
type byInt		[]seila
type byString	[]seila
type byFloat64	[]seila

// ByInt
func (b byInt) Len() int { return len(b) }
func (b byInt) Less(i, j int) bool { return b[i].a2 < b[j].a2 }
func (b byInt) Swap(i, j int) { b[i],b[j] = b[j], b[i] }

// ByString
func (b byString) Len() int { return len(b) }
func (b byString) Less(i,j int) bool { return b[i].a1 < b[j].a1 }
func (b byString) Swap(i,j int) { b[i],b[j] = b[j], b[i] }

// ByFloat64 
func (b byFloat64) Len() int { return len(b) }
func (b byFloat64) Less(i,j int) bool { return b[i].a3 < b[j].a3 }
func (b byFloat64) Swap(i,j int) { b[i], b[j] = b[j],b[i] }


func main() {
	var bacon = []seila{
		seila{"z",2,7.17},
		seila{"a",10,9.67},
		seila{"y",3,7.67},
		seila{"b",12,7},
		seila{"x",112,9.7},
		seila{"c",0,10.7},
		seila{"w",20,36.67},
		seila{"d",1,16.67},
	}
	println()
	println("--> Original")
	ajeitaAi(bacon)
	
	println()
	println("--> Float")
	sort.Sort(byFloat64(bacon))
	ajeitaAi(bacon)
	
	println()
	println("--> String")
	sort.Sort(byString(bacon))
	ajeitaAi(bacon)
	
	println()
	println("--> Integer")
	sort.Sort(byInt(bacon))
	ajeitaAi(bacon)
}