package main

import "fmt"

func main() {
	x := func() int {
		return 2
	}
	fmt.Println(x())
}
