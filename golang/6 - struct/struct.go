package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero uint8
}

func main() {
	var u usuario
	u.nome = "Raissa"
	u.idade = 20

	enderecoUsuario := endereco{"Rua ABC", 123}

	usuario2 := usuario{"Raissa", 23, enderecoUsuario}
	usuario3 := usuario{idade: 4}

	fmt.Println(u, usuario2, usuario3)
}
