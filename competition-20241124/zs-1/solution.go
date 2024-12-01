package main

import "math"

func minimumSumSubarray(nums []int, l int, r int) int {
	ans := math.MaxInt
	arr := make([]int, r-l+1)
	for i, v := range nums {
		for j := range arr {
			arr[j] += v
			if l+j <= i {
				arr[j] -= nums[i-l-j]
			}
			if l+j <= i+1 && arr[j] > 0 && arr[j] < ans {
				ans = arr[j]
			}
		}
	}
	if ans == math.MaxInt {
		ans = -1
	}
	return ans
}
