package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	altura	float64
	largura	float64
}
func ( q quadrado ) area() float64 {
	return (q.altura * q.largura)
}


type circulo struct {
	raio float64
}
func ( c circulo ) area() float64 {
	return (2*math.Pi*c.raio)
}


type figura interface {
	area() float64
}
func info(f figura) float64 {
	return f.area()
}


func main() {
	quadro	:= quadrado{4,4}
	circo	:= circulo{3}
	fmt.Println(info(quadro))
	fmt.Println(info(circo))
}