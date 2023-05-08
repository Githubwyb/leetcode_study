package main

import "math"

func findSmallestInteger(nums []int, value int) int {
	grp := make([]int, value)
	for _, v := range nums {
		i := 0
		if v >= 0 {
			i = v % value
		} else {
			i = (value - (-v)%value) % value
		}
		grp[i]++
	}

	min, index := math.MaxInt, -1
	for i, v := range grp {
		if v < min {
			// 找到数量最小的某个位置和数量
			min = v
			index = i
		}
	}
	return index + min*value
}
