package main

import "sort"

func findNonMinOrMax(nums []int) int {
	if len(nums) <= 2 {
		return -1
	}
	sort.Ints(nums[:3])
	return nums[1]
}
