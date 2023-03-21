package main

import "fmt"

func evenOddBit(n int) []int {
	ev, od := 0, 0
	str := fmt.Sprintf("%b", n)
	length := len(str)
	for i := 0; i < length; i++ {
		if str[length-i-1] == '1' {
			if i%2 == 0 {
				ev++
			} else {
				od++
			}
		}
	}
	return []int{ev, od}
}
