package main

import (
	"fmt"
)

func findTheArrayConcVal(nums []int) int64 {
	var result int64 = 0
	for i, j := 0, len(nums)-1; i <= j; i, j = i+1, j-1 {
		if i == j {
			result += int64(nums[i])
			break
		}
		var tmp int64 = int64(nums[i])
		for _ = range fmt.Sprint(nums[j]) {
			tmp *= 10
		}
		tmp += int64(nums[j])
		result += tmp
	}
	return result
}
