package main

import (
	"math"
)

func myAtoi(s string) int {
	var result int64
	sign := 1
	isParse := false
	for _, v := range s {
		if v <= '9' && v >= '0' {
			if !isParse {
				result = int64(v - '0')
				isParse = true
				continue
			}
			result *= 10
			result += int64(v - '0')
			if sign == 1 && result > math.MaxInt32 {
				return math.MaxInt32
			} else if result-1 > math.MaxInt32 {
				return math.MinInt32
			}
			continue
		} else if isParse {
			break
		} else if v == '-' {
			sign = -1
			isParse = true
		} else if v == '+' {
			isParse = true
		} else if v != ' ' {
			return 0
		}
	}
	return int(result) * sign
}
