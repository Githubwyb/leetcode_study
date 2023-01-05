package main

import "fmt"

func countDigits(num int) int {
	result := 0
	for _, v := range fmt.Sprint(num) {
		if v != '0' && num%int(v-'0') == 0 {
			result++
		}
	}
	return result
}
