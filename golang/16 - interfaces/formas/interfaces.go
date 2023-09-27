package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

func escreverArea(forma forma){
	fmt.Printf("A área da forma é: %0.2f\n", forma.area())
}

type retangulo struct {
	altura float64
	largura float64
}

func (retangulo retangulo) area() float64{
	return retangulo.altura * retangulo.largura
}

type circulo struct {
	raio float64
}

func (circulo circulo) area() float64{
	return math.Pi * math.Pow(circulo.raio, 2)
}

func main(){
	retangulo := retangulo{12, 4}
	escreverArea(retangulo)

	circulo := circulo{10}
	escreverArea(circulo)
}