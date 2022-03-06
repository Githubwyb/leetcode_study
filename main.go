package main

import (
	"fmt"
)

const(
    a = iota		// 0	itoa = 0
    b = 5			// 5	itoa = 1
    c = 3			// 3	itoa = 2
    d = iota		// 3	itoa = 3
    e = 6			// 6
    f				// 6
    g				// 6
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}
