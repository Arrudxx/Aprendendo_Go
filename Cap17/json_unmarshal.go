package main

import (
	"encoding/json"
	"fmt"
)

type info struct {
	Nome          string  `json:"Nome"`
	Sobrenome     string  `json:"Sobrenome"`
	Idade         int     `json:"Idade"`
	Profissao     string  `json:"Profissao"`
	Contabancaria float64 `json:"Contabancaria"`
}

func main() {

	sb := []byte(`{"Nome":"Daniel","Sobrenome":"Arruda","Idade":23,"Profissao":"Agente Secreto","Contabancaria":1000.2}`)

	var daniel info
	err := json.Unmarshal(sb, &daniel)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(daniel)
	fmt.Println(daniel.Profissao)

}
