package main

import "fmt"

func main() {
	numero := 20
	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor do que 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero maior que 0")
	} else {
		fmt.Println("NUmero menor que 0")
	}
}
