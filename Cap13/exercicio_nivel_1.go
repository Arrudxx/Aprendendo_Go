package main

import "fmt"

func main() {
	fmt.Println(retornaint())
	fmt.Println(retornastring())
}

func retornaint() int {
	return 8
}

func retornastring() (int, string) {
	return 20, "vint"
}
