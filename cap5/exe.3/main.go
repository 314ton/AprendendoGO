package main

import (
	"fmt"
)

type veículo struct{
	cor 	string
	portas	int
	modelo	string
}
type sedan struct{
	veículo
	modeloLuxo	bool
}
type caminhonete struct{
	veículo
	traçãoNasQuatro bool
}

func main(){
	sedan := sedan{
		veículo: veículo{
			cor: "azul",
			portas: 2,
			modelo: "camaro",
		},
		modeloLuxo: true,
	}
	caminhonete := caminhonete{
		veículo: veículo{
			cor: "prata",
			portas: 4,
			modelo: "hilux",
		},
		traçãoNasQuatro: true,
	}
	fmt.Println("caminhonete: ->",caminhonete.modelo)
	fmt.Println(
		"caminhonete: ->",caminhonete.modelo,
		"Cor:->",caminhonete.cor,
	)
	
	fmt.Println("sedan: ->",sedan.modelo)
	fmt.Println(
		"sedan: ->",sedan.modelo,
		"Cor:->",sedan.cor,
	)
}