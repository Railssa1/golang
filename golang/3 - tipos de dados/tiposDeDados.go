package main

import (
	"errors"
	"fmt"
)

func main(){
	// 4 tipos de inteiro int8, int16, int32 e int64
	var numero int64 = 14500234320
	fmt.Println(numero)

	// Quando infere o valor do int ele assume a configuração do pc
	numero2 := -45656
	fmt.Println(numero2)

	// alias
	// int32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 uint8 = 10
	fmt.Println(numero4)

	// Numeros reias
	var numeroReal1 float32 = 124.67
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1333344.90
	fmt.Println(numeroReal2)

	numeroReal3 := 23345.04
	fmt.Println(numeroReal3)

	var erro error = errors.New("Erro interno.")
	fmt.Println(erro)
}