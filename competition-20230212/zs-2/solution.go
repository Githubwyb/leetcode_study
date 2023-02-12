package main

import "sort"

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	l, r := len(nums), len(nums)	// l代表大于i的第一个下界，r代表大于l的第一个上界
	var result int64 = 0
	//
	for i := 0; i < len(nums) && r > i+1; i++ {
		// i增加后，l应该减小，那么就要查找 [i+1, l) 区间范围
		l = sort.Search(l-i-1, func(index int) bool {
			index += i + 1
			return nums[index]+nums[i] >= lower
		}) + i + 1
		// i增加后，r也应该减小，查找范围为，[l, r)
		r = sort.Search(r-l, func(index int) bool {
			index += l
			return nums[index]+nums[i] > upper
		}) + l
		result += int64(r - l)
	}
	return result
}
