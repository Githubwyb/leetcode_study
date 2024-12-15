package main

func resultArray(nums []int) []int {
	if len(nums) <=2 {
		return nums
	}
	var arr1, arr2 []int
	arr1 = append(arr1, nums[0])
	last1, last2 := nums[0], nums[1]
	arr2 = append(arr2, nums[1])
	for _, v := range nums[2:] {
		if last1 > last2 {
			arr1 = append(arr1, v)
			last1 = v
		} else {
			arr2 = append(arr2, v)
			last2 = v
		}
	}
	return append(arr1, arr2...)
}
