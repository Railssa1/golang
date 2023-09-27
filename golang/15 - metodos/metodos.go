package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
}

func (usuario usuario) salvar(){
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", usuario.nome)
}

func (usuario usuario) isMaiorDeIdade() bool {
	return usuario.idade >= 18
}

func (usuario *usuario) fazerAniversario(){
	usuario.idade++
}

func main(){
	usuario := usuario{"Raissa", 20}
	usuario.salvar()
	fmt.Println(usuario.isMaiorDeIdade())
	usuario.fazerAniversario()
	fmt.Println(usuario.idade)
}