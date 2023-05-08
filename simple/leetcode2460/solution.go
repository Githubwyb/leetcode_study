package main

func applyOperations(nums []int) []int {
	n := len(nums)
	j := 0 // j是非0的索引
	for i := 0; i < n; i++ {
		if i == n-1 {
			nums[j] = nums[i]
			j++
			continue
		}
		if nums[i] > 0 {
			if nums[i] == nums[i+1] {
				nums[j] = nums[i] * 2
				nums[i+1] = 0
				i++
			} else {
				nums[j] = nums[i]
			}
			j++
		}
	}
	for ; j < n; j++ {
		nums[j] = 0
	}
	return nums
}
