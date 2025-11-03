package main

import "fmt"

type Pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func mudeMe(p *Pessoa) {
	(*p).nome = "Daniel"
	p.sobrenome = "Arruda"
}

func main() {

	Thiago := Pessoa{"Thiago", "Alves", 22}
	fmt.Println(Thiago)
	mudeMe(&Thiago)
	fmt.Println(Thiago)

}
