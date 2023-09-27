package main

import "fmt"

func main(){

	usuario := map[string]string {
		"nome": "Pedro",
		"sobrenome": "Alves",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome":{
			"primeiro": "João",
			"ultimo": "Pedro",
		},
		"curso": {
			"nome": "Engenharia",
			"campus": "Campus 1",
		},
	}

	delete(usuario2, "curso")

	usuario2["signo"] = map[string]string{
		"nome": "Gemeos",
	}

	fmt.Println(usuario2)
}