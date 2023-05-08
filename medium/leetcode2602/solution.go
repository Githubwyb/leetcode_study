package main

import "sort"

func minOperations(nums []int, queries []int) []int64 {
	sort.Ints(nums)
	n := len(nums)
	tmp := make([]int64, n)
	for _, v := range nums {
		tmp[0] += int64(v - nums[0])
	}
	for i, v := range nums {
		if i == 0 {
			continue
		}
		tmp[i] = tmp[i-1] - int64((n-i-i)*(v-nums[i-1]))
	}
	res := make([]int64, len(queries))
	for j, v := range queries {
		index := sort.Search(n, func(i int) bool {
			return nums[i] >= v
		})
		if index == n {
			res[j] = tmp[index-1] + int64(n*(v-nums[index-1]))
		} else {
			res[j] = tmp[index] + int64((nums[index]-v)*(n-index-index))
		}
	}
	return res
}

func minOperations1(nums []int, queries []int) []int64 {
	sort.Ints(nums)
	n := len(nums)
	prefixSum := make([]int64, n+1) // 多一个防止越界
	var sum int64 = 0
	for i, v := range nums {
		sum += int64(v)
		prefixSum[i+1] = sum
	}

	// 开始二分找queries的位置计算
	res := make([]int64, len(queries))
	for i, v := range queries {
		index := sort.SearchInts(nums, v)
		res[i] = int64((index*2-n)*v) + sum - prefixSum[index]*2
	}
	return res
}
