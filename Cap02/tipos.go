package main

import "fmt"

// Tipos em Go são estáticos, se uma variavel é por exmplo int, ela será int até o final
// Tipos de dados primitivos: disponíveis na linguagem nativamente como blocos básicos de construção int, string, bool
// Tipos de dados compostos: são tipos compostos de tipos primitivos, e criados pelo usuário slice, array, struct, map

var x int = 10

func main() {

	fmt.Printf("%v %T", x, x)

}
