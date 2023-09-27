package main

import "fmt"

func main(){
	// OPERADORES ARITMETICOS
	soma := 34 + 3
	subtracao := 14 - 2
	divisao := 12 / 4
	multiplicacao := 3 * 2
	restoDaDivisao := 9 / 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	// OPERADORES LOGICOS 
	fmt.Println(true && true)
}