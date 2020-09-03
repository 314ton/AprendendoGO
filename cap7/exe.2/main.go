package main


type pessoa struct {
	nome	string
	email	string
}
func mudeMe(p pessoa) {
	p.nome = "juão"
	p.email = "juão@juão"
}
func mudeMe2(p *pessoa) {
	// ambos estão 'certos'
	(*p).nome = "predo"
	p.email = "predo@predo"
}


func main() {
	p1 := pessoa{"etho","etho@etho"}
	println(p1.nome)
	println(p1.email)
	
	mudeMe(p1)
	println(p1.nome)
	println(p1.email)
	
	mudeMe2(&p1)
	println(p1.nome)
	println(p1.email)
	
}