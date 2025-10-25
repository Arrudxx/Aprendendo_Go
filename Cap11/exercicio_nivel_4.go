package main

import "fmt"

func main() {

	pessoa := struct {
		nome           string
		idade          int
		altura         int
		comidafavorita []string
		pratica        map[string]bool
	}{
		nome:           "Daniel",
		idade:          23,
		altura:         183,
		comidafavorita: []string{"pizza", "hamburguer"},
		pratica: map[string]bool{
			"natacao": true,
			"luta":    false,
		},
	}

	fmt.Println(pessoa)

}
