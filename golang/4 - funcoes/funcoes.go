package main

import "fmt"

func somar(numero1 int8, numero2 int8) int8 {
	return numero1 + numero2
}

func calculos(numero1, numero2 int8) (int8, int8) {
	soma := numero1 + numero2
	subtracao := numero1 - numero2
	return soma, subtracao
}

func main() {
	soma := somar(23, 45)
	fmt.Println(soma)

	var f = func(texto string) string {
		fmt.Println(texto)
		return texto
	}

	resultadoDaFuncao := f("Chamando função armazenada na variável")
	fmt.Println(resultadoDaFuncao)

	//Utiliza o "_" quando só deseja um resultado do return
	resultadoSoma, _ := calculos(32, 7)
	fmt.Println(resultadoSoma)
}
