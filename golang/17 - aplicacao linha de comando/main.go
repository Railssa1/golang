package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main(){
	fmt.Println("Ponto de partida")

	app := app.Gerar()
	if erro := app.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}