package main

import "fmt"

func main() {

	i := 33
	for i <= 122 {
		fmt.Printf(", %d, %#x, %#U \n", i, i, i)
		i++
	}

}
