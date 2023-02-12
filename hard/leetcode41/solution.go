package main

func firstMissingPositive(nums []int) int {
	for i := range nums {
		// 不在范围内或相等（交换也相等或位置一样）就退出
		for nums[i] > 0 && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := range nums {
		if i+1 != nums[i] {
			return i + 1
		}
	}
	return len(nums) + 1
}
