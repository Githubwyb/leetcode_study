package main

import "sort"

func unequalTriplets(nums []int) int {
	result := 0
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			if nums[j] == nums[i] {
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				if nums[k] == nums[j] || nums[k] == nums[i] {
					continue
				}
				result++
			}
		}
	}
	return result
}

func unequalTriplets1(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	// x作为中间的数，那么小于x的有a个，x有b个，大于x的有c个，x在中间的数对有abc个
	start, res := 0, 0 // start为x的起始位置
	for i, x := range nums[:n-1] {
		if x != nums[i+1] {
			res += start * (i + 1 - start) * (n - i - 1)
			start = i + 1
		}
	}
	return res
}

func unequalTriplets2(nums []int) int {
	h := make(map[int]int)
	for _, v := range nums {
		h[v]++
	}
	a, c, res := 0, len(nums), 0
	for _, b := range h {
		c -= b
		res += a * b * c
		a += b
	}
	return res
}
