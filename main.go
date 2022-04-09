package main

import (
	"fmt"
)

func main() {
	testArr := []int{2, 7, 4, 9, 1, 4, 8}
	fmt.Println(cap(testArr))
	testArr = testArr[:0]
	fmt.Println(cap(testArr))
	fmt.Println(testArr)
	// testArr = testArr[1:len(testArr)-1]
	// fmt.Println(cap(testArr))
}
