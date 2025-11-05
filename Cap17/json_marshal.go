package main

import (
	"encoding/json"
	"fmt"
)

type pessoa struct {
	Nome          string
	Sobrenome     string
	Idade         int
	Profissao     string
	Contabancaria float64
}

func main() {
	daniel := pessoa{
		Nome:          "Daniel",
		Sobrenome:     "Arruda",
		Idade:         23,
		Profissao:     "Agente Secreto",
		Contabancaria: 1000.20,
	}

	bruno := pessoa{
		Nome:          "Bruno",
		Sobrenome:     "Arruda",
		Idade:         28,
		Profissao:     "radiologista",
		Contabancaria: 500.50,
	}

	d, err := json.Marshal(daniel)
	if err != nil {
		fmt.Println(err)
	}

	b, err := json.Marshal(bruno)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(d))
	fmt.Println(string(b))

}
