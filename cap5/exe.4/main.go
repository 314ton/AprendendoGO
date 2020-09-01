package main


import "fmt"


func main(){
	x := struct{
		meus_números map[string][]int
		
	}{
		meus_números: map[string][]int{
			"vivo": []int{
				998319634,
			},
			"oi": []int{
				988011213,
			},
		},
	}
	fmt.Println(x.meus_números)
	fmt.Println(x.meus_números["vivo"])
	fmt.Println(x.meus_números["oi"])
}