package main

import "sort"

func maximumCount(nums []int) int {
	neg := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= 0
	})
	pos := sort.Search(len(nums)-neg, func(i int) bool {
		return nums[i+neg] > 0
	}) + neg
	if neg < len(nums)-pos {
		return len(nums) - pos
	}
	return neg
}
