package main

import "fmt"

func recuperarExecucao(){
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso")
	}
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2
	if media > 6{
		fmt.Println("Aluno esta aprovado")
		return true
	}else if media < 6 {
		fmt.Println("Aluno esta reprovado")
		return false
	}	

	panic("O sistema falhou!")
}

func main(){
	alunoEstaAprovado(6, 6)
}