package main

import (
	"encoding/json"
	"os"
)

type pessoa struct {
	Nome          string  `json:"Nome"`
	Sobrenome     string  `json:"Sobrenome"`
	Idade         int     `json:"Idade"`
	Profissao     string  `json:"Profissao"`
	Contabancaria float64 `json:"Contabancaria"`
}

func main() {

	daniel := pessoa{
		Nome:          "Daniel",
		Sobrenome:     "Arruda",
		Idade:         23,
		Profissao:     "Agente Secreto",
		Contabancaria: 1000.20,
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(daniel)
}
