package main

import (
	"math"
	"sort"
)

func getLargestOutlier(nums []int) int {
	sum := 0
	numMap := make(map[int]int)
	for _, v := range nums {
		sum += v
		numMap[v]++
	}

	ans := math.MinInt
	for _, v := range nums {
		// 从大到小遍历异常值，找剩下的是否有和sum/2相等的
		tmp := sum - v
		if tmp%2 != 0 {
			continue
		}
		s := tmp / 2
		if (numMap[s] > 1 || (s != v && numMap[s] > 0)) && v > ans {
			ans = v
		}
	}
	return ans
}

func getLargestOutlier1(nums []int) int {
	sum := 0
	numMap := make(map[int]int)
	for _, v := range nums {
		sum += v
		numMap[v]++
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	for _, v := range nums {
		// 从大到小遍历异常值，找剩下的是否有和sum/2相等的
		tmp := sum - v
		if tmp%2 != 0 {
			continue
		}
		s := tmp / 2
		if numMap[s] > 1 || (s != v && numMap[s] > 0) {
			return v
		}
	}
	return -1
}
