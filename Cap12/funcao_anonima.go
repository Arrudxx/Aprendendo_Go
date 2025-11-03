package main

import "fmt"

func main() {
	x := 32

	func(x int) {
		fmt.Println(x, "vezes", x, "é:")
		fmt.Println(x * x)
	}(x)
}
