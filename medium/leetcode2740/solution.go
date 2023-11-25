package main

import (
	"math"
	"sort"
)

func findValueOfPartition(nums []int) int {
	sort.Ints(nums)
	res := math.MaxInt
	for i := range nums[1:] {
		diff := nums[i+1] - nums[i]
		if diff < res {
			res = diff
		}
	}
	return res
}
