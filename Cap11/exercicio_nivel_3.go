package main

import "fmt"

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	quatroRodas bool
}

type sedam struct {
	veiculo
	modeloluxo bool
}

func main() {

	sedam := sedam{
		veiculo:    veiculo{2, "roxo"},
		modeloluxo: true,
	}

	caminhonete := caminhonete{
		veiculo:     veiculo{4, "Azul"},
		quatroRodas: true,
	}

	fmt.Println(caminhonete)
	fmt.Println(sedam)
	fmt.Println(sedam.cor)
	fmt.Println(caminhonete.quatroRodas)

}
