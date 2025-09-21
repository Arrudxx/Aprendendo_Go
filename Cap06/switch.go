package main

import "fmt"

func main() {

	// x := 10

	// switch {
	// case x < 5:
	// 	fmt.Println("x é menor que cinco")
	// case x == 5:
	// 	fmt.Println("x é igual a cinco")
	// case x > 5:
	// 	fmt.Println("x é maior que cinco")
	// }

	nome := "maria"

	// fallthrough: roda esse e o proximo caso

	switch nome {
	case "zezinho", "maria":
		fmt.Println("O nome é o zezinho e maria")
		// fallthrough
	case "joao", "daniel":
		fmt.Println("O nome é o joao e daniel")
	default:
		fmt.Println("entrou no default")
	}

}
