package main

import "math"

func diagonalPrime(nums [][]int) int {
	res := 0
	n := len(nums)
	isprime := func(x int) bool {
		mx := int(math.Sqrt(float64(x))) + 1
		for i := 2; i < mx; i++ {
			if x%i == 0 {
				return false
			}
		}
		return x >= 2
	}
	for i := range nums {
		v := nums[i][i]
		if v > res && isprime(v) {
			res = v
		}
		v = nums[i][n-i-1]
		if v > res && isprime(v) {
			res = v
		}
	}
	return res
}
