package main

import (
	"fmt"
)

func differenceOfSum(nums []int) int {
	l, r := 0, 0
	for _, v := range nums {
		l += v
		for _, v1 := range fmt.Sprint(v) {
			r += int(v1 - '0')
		}
	}
	return l - r
}
