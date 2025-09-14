package main

import "fmt"

type hotdog int //Criando tipo

var b hotdog = 101 // Atribuindo tipo a uma variavel

func main() {
	x := 10
	fmt.Printf("%v\n", x)

	x = int(b)
	fmt.Printf("%v\n", x)
	fmt.Printf("%v", b)

}

// Conversion = Termo geral para qualquer alteração de tipo.
// Casting = Conversão manual, solicitada pelo programador.
// coercion = Conversão automática, feita pelo compilador ou interpretador.
