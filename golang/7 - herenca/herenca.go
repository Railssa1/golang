package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
}

type estudante struct {
	pessoa
	ra uint8
	curso string
}

func main() {
	p1 := pessoa{"Raissa", "Moraes", 20}
	e1 := estudante{p1, 123, "Administração"}

	fmt.Println(e1)
}
