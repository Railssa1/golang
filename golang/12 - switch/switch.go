package main

import "fmt"

func getDiaDaSemana(dia int) string{
	switch dia {
		case 1:
			return "Domingo"
		case 2:
			return "Segunda-feira"
		case 3:
			return "Terça-feira"
		case 4:
			return "Quarta-feira"
		case 5:
			return "Quinta-feira"
		case 6:
			return "Sexta-feira"
		case 7:
			return "Sabado"
		default:
			return "Dia inválido"
	}
}


func main(){
	fmt.Println(getDiaDaSemana(3))
}