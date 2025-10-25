package main

import "fmt"

type Pessoa struct {
	nome            string
	sobrenome       string
	saboressorvetes []string
}

func main() {

	pessoa1 := Pessoa{
		nome:            "Daniel",
		sobrenome:       "Arruda",
		saboressorvetes: []string{"chocolate", "morango"},
	}

	pessoa2 := Pessoa{
		nome:            "Pedro",
		sobrenome:       "Teste",
		saboressorvetes: []string{"creme", "limão"},
	}

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)

	for i, v := range pessoa1.saboressorvetes {
		fmt.Println(i, " - ", v)
	}

}
