package main

import (
	"fmt"
)

type Pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (pessoa Pessoa) demonstrar() {
	fmt.Println("Nome completo dessa pessoa é:", pessoa.nome, pessoa.sobrenome, "e a sua idade é:", pessoa.idade, "anos")
}

func main() {

	pessoa1 := Pessoa{
		nome:      "Daniel",
		sobrenome: "Arruda",
		idade:     13,
	}

	pessoa1.demonstrar()

}
