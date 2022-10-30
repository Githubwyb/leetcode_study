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

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func findMaxK1(nums []int) int {
	numMap := make(map[int]bool)
	result := -1
	for _, v := range nums {
		if _, ok := numMap[-v]; ok {
			if result < abs(v) {
				result = abs(v)
			}
		}
		numMap[v] = true
	}
	return result
}

func findMaxK2(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	for l < r {
		addResult := nums[l] + nums[r]
		if addResult == 0 {
			return nums[r]
		} else if addResult > 0 {
			r--
		} else {
			l++
		}
	}
	return -1
}
