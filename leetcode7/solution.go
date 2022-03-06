package main

import (
	"fmt"
	"math"
)

func reverse(x int) (result int) {
	tmp := fmt.Sprint(x)
	for i := range(tmp) {
		if tmp[len(tmp)-i-1] == '-' {
			result = -result
			break
		}
		if result > math.MaxInt32 / 10 || result < math.MinInt32 / 10 {
			result = 0
			return
		}
		result *= 10
		result += int(tmp[len(tmp)-1-i] - '0')
	}
	return
}
