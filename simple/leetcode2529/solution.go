package main

import "sort"

func maximumCount(nums []int) int {
	// 二分查找第一个正数的位置
	pos := sort.Search(len(nums), func(i int) bool {
		return nums[i] > 0
	})
	// 对剩余的二分查找第一个0
	neg := sort.Search(pos, func(i int) bool {
		return nums[i] >= 0
	})

	if neg < len(nums)-pos {
		return len(nums) - pos
	}
	return neg
}
