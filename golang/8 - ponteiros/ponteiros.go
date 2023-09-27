package main

import "fmt"

func main(){
	var variavel1 int
	var ponteiro *int

	variavel1 = 10
	ponteiro = &variavel1

	fmt.Println(variavel1, *ponteiro) //desreferenciação
}