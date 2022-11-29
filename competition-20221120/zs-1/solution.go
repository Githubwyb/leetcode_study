package main

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
