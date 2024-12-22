package main

import (
	"math"
	"sort"
)

func maxDistinctElements(nums []int, k int) int {
	sort.Ints(nums)
	res := 0
	last := math.MinInt
	for _, v := range nums {
		// 如果当前数字变成前一个+1
		x := min(max(last+1, v-k), v+k)
		if x > last {
			res++
			last = x
		}
	}
	return res
}
