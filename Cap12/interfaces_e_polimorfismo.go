package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type dentista struct {
	pessoa
	dentesarrancados int
	salario          int
}

type arquiteto struct {
	pessoa
	tipodeconstrucao string
	tamanhodaloucura string
}

func (x dentista) oibomdia() {
	fmt.Println("Meu nome é ", x.nome, "eu tenho ", x.dentesarrancados, " e Bom dia!")
}

func (x arquiteto) oibomdia() {
	fmt.Println("Meu nome é", x.nome, "e Bom dia!")
}

type gente interface {
	oibomdia()
}

func serhumano(g gente) {
	g.oibomdia()
	switch g.(type) {
	case dentista:
		fmt.Println("Eu ganho:", g.(dentista).salario)

	case arquiteto:
		fmt.Println("Eu construo:", g.(arquiteto).tipodeconstrucao)
	}
}

func main() {
	mrdente := dentista{
		pessoa: pessoa{
			nome:      "Dente",
			sobrenome: "Da Silva",
			idade:     50,
		},
		dentesarrancados: 8973,
		salario:          31234,
	}
	mrpredio := arquiteto{
		pessoa: pessoa{
			nome:      "Predio",
			sobrenome: "Dos Santos",
			idade:     43,
		},
		tipodeconstrucao: "hospital",
		tamanhodaloucura: "Muito",
	}

	serhumano(mrdente)
	serhumano(mrpredio)

}
