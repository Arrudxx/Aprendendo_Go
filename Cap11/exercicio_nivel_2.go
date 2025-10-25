package main

import "fmt"

type pessoa struct {
	nome            string
	sobrenome       string
	saboressorvetes []string
}

func main() {

	//var meumapa map[string]pessoa

	meumapa := make(map[string]pessoa)

	// meumapa2 := map[string]pessoa{
	// 	"Pimentão": pessoa{
	//		nome:      "Renata",
	//		sobrenome: "Pimentão",
	//		saboressorvetes:   []string{"pistache", "morango", "baunilha"}},
	//	"da Prússia": pessoa{"Frederico", "da Prússia",
	//		[]string{"sabão em pó", "pé de coelho", "feijão"}},
	// }

	meumapa["Pimentão"] = pessoa{
		nome:            "Renata",
		sobrenome:       "Pimentão",
		saboressorvetes: []string{"pistache", "morango", "baunilha"}}

	meumapa["da Prússia"] = pessoa{"Frederico", "da Prússia",
		[]string{"sabão em pó", "pé de coelho", "feijão"}}

	for _, valor := range meumapa {
		fmt.Println("Meu nome é", valor.nome, "e meus sorvetes favoritos são:")
		for _, valor := range valor.saboressorvetes {
			fmt.Println("–", valor)
		}
		fmt.Println("\n")
	}
}
