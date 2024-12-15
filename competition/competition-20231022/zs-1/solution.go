package main

import "math"

func minimumSum(nums []int) int {
	n := len(nums)
	// 先从左向右遍历一边，找每个下标小于其的最小值
	leftNums := make([]int, n)
	min := math.MaxInt
	for i, v := range nums {
		if i == 0 {
			min = v
			continue
		}
		if min < v {
			leftNums[i] = min
		}
		if min > v {
			min = v
		}
	}
	// 再从右侧遍历一遍
	rightNums := make([]int, n)
	min = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		v := nums[i]
		if min < v {
			rightNums[i] = min
		}
		if min > v {
			min = v
		}
	}

	res := -1
	for i := 1; i < n-1; i++ {
		if leftNums[i] == 0 || rightNums[i] == 0 {
			continue
		}
		sum := leftNums[i] + rightNums[i] + nums[i]
		if res == -1 || res > sum {
			res = sum
		}
	}
	return res
}
