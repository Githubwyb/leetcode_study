package main

import "fmt"

func main() {
	a := 1
	fmt.Println(a) // 1
	a ^= 1
	fmt.Println(a) // 0
	a ^= 1
	fmt.Println(a) // 1
}
