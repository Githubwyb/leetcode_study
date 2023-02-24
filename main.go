package main

import "fmt"

func getFirst(i int) int {
	v := (^i) + 1
	return i & v
}

func main() {
	fmt.Printf("0b%b\r\n", getFirst(0b00110110))
}
