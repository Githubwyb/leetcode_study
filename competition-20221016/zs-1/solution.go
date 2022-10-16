package main

import "sort"

func findMaxK(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums) && nums[i] < 0; i++ {
		for j := len(nums) - 1; j >= 0 && nums[j] > 0; j-- {
			if -nums[i] == nums[j] {
				return nums[j]
			}
			if -nums[i] > nums[j] {
				break
			}
		}
	}
	return -1
}
