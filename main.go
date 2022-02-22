package main

import (
	"fmt"
)

func main() {
	s := make([]int, 5)
	s1 := make([]int, 0, 10)
	fmt.Println("s", len(s), cap(s))
	fmt.Println("s1", len(s1), cap(s1))
}
