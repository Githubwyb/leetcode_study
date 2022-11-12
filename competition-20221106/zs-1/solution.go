package main

func applyOperations(nums []int) []int {
	result := make([]int, len(nums))
	j := 0
	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 {
			result[j] = nums[i]
			continue
		}
		if nums[i] == 0 {
			continue
		}
		result[j] = nums[i]
		if nums[i] == nums[i+1] {
			result[j] += nums[i]
			i++
		}
		j++
	}
	return result
}
