package main

import "sort"

func maxNumOfMarkedIndices(nums []int) int {
	sort.Ints(nums)
	n := len(nums)

	// k的范围是0到n/2
	// 二分查找第一个不满足条件的，所以二分的范围应该是1到n/2 + 1
	// 转化一下初始值为1，二分进去的数要加一，出来的数是第一个不满足条件的，需要减一乘2
	// 可是初始值加了1就不用减了，直接乘2返回
	return sort.Search(n/2, func(k int) bool {
		k++
		for i := 0; i < k; i++ {
			if nums[i]*2 > nums[n-k+i] {
				return true
			}
		}
		return false
	}) * 2
}
