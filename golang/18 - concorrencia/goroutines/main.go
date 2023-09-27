package main

import (
	"fmt"
	"time"
)

func main(){
	// Concorrência != paralelismo

	go escrever("Bem-vindo(a)")
	escrever("Programando em go")
}


func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}