package main 


import (
	"fmt"
)


func main(){
	dict := map[string][]string{
		"etho": []string{
			"estudar programação",
			"estudar GO",
			"estudar flask",
		},
		"amarelo": []string{
			"ir a igreja",
			"trabalhar",
			"zuar",
		},
		"marcus": []string{
			"zuar também",
			"nada",
		},
	}
	dict["juoquim"] = []string{
			"batata",
			"sabonete",
			"beterraba",
			"Alcachofra",
	}
	delete(dict,"marcus")
	println()
	for k,v := range dict{
		fmt.Printf("%s gosta de:\n",k)
		for _,item := range v{
			fmt.Printf("->\t%s\n",item)
		}
		println()
	}
	println()
}