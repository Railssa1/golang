package main

import (
	"fmt"
	"time"
)

func main(){
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementado J", j)
		time.Sleep(time.Second)		
	}

	nomes := [3]string{"Jonas", "Lucas", "Paulo"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA"{
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome": "Raissa",
		"Sobrenome": "Alves",
	}

	for chave, valor := range usuario{
		fmt.Println(chave, valor)
	}
}