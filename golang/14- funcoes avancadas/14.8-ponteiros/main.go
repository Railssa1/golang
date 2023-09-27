package main

import (
	"fmt"
)

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int){
	*numero = *numero * -1
}

func main(){
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	numero2 := 40
	fmt.Println(numero2)
	inverterSinalComPonteiro(&numero2)
	fmt.Println(numero2)

}