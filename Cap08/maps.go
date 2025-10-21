package main

import "fmt"

func main() {

	amigos := map[string]int{
		"alfredo": 13123,
		"joana":   3123914,
	}

	fmt.Println(amigos["joana"])

	amigos["daniel"] = 3123

	fmt.Println(amigos, "\n")

	if sera, ok := amigos["bruno"]; !ok {
		fmt.Println("Não tem")
	} else {
		fmt.Println(sera)
	}

}
