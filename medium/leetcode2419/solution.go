package main

func longestSubarray(nums []int) int {
	max := nums[0]
	result := 1
	length := 1
	for i, v := range nums {
		if i == 0 {
			continue
		}
		if v > max {
			max = v
			length = 1
			result = 1
		} else if v == max {
			length++
		} else {
			if length > result {
				result = length
			}
			length = 0
		}
	}
	if length > result {
		result = length
	}
	return result
}
