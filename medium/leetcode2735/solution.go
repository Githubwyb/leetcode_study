package main

import "math"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minCost(nums []int, x int) int64 {
	n := len(nums)
	minNums := make([]int, n)
	copy(minNums, nums)
	var res int64 = math.MaxInt64
	for i := 0; i < n; i++ {
		var sum int64 = int64(i * x)
		for j, v := range minNums {
			k := (i + j) % n
			minNums[j] = min(v, nums[k])
			sum += int64(minNums[j])
		}
		if sum < res {
			res = sum
		}
	}
	return res
}
