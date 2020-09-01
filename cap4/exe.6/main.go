package main

import (
	"fmt"
)


func main(){
	slice := make([]string ,0,28)
	
	slice = []string{
		"Acre (AC)",
		"Alagoas (AL)",
		"Amapá (AP)",
		"Amazonas (AM)",
		"Bahia (BA)",
		"Ceará (CE)",
		"Distrito Federal (DF)",
		"Espírito Santo (ES)",
		"Goiás (GO)",
		"Maranhão (MA)",
		"Mato Grosso (MT)",
		"Mato Grosso do Sul (MS)",
		"Minas Gerais (MG)",
		"Pará (PA)",
		"Paraíba (PB)",
		"Paraná (PR)",
		"Pernambuco (PE)",
		"Piauí (PI)",
		"Rio de Janeiro (RJ)",
		"Rio Grande do Norte (RN)",
		"Rio Grande do Sul (RS)",
		"Rondônia (RO)",
		"Roraima (RR)",
		"Santa Catarina (SC)",
		"São Paulo (SP)",
		"Sergipe (SE)",
		"Tocantins (TO",
	}
	fmt.Printf(
		"\033[34mCapacit.-> %v, Len.-> %v\n\n\033[m",
		cap(slice),
		len(slice),
	)
	for i,v := range slice {
		fmt.Printf("Index[ %.0f ] -> %v \n",float64(i),v)
	}
	println("")
}