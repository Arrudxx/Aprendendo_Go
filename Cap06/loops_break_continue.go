package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			// quebra somente a interação
			continue
			// quebra todo o loop
			// break
		}
		fmt.Println(i)
	}
}
