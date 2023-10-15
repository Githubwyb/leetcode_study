package main

import "sort"

func countWays(nums []int) int {
	sort.Ints(nums)
	res := 0
	n := len(nums)
	if nums[0] > 0 {
		res++
	}
	for i, x := range nums[:n-1] {
		// 前i+1个学生被选中
		if i+1 > x && i+1 < nums[i+1] {
			res++
		}
	}
	return res + 1
}
