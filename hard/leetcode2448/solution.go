package main

import "sort"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minCost(nums []int, cost []int) int64 {
	type item struct {
		val, count int
	}
	countMap := make([]item, len(nums))
	sum := 0
	for i, v := range nums {
		countMap[i] = item{
			val:   v,
			count: cost[i],
		}
		sum += cost[i]
	}
	sort.Slice(countMap, func(i, j int) bool { return countMap[i].val < countMap[j].val })
	midI := sum / 2
	median := 0
	for _, v := range countMap {
		if midI < v.count {
			median = v.val
			break
		}
		midI -= v.count
	}
	var res int64 = 0
	// 所有数变成中位数
	for _, v := range countMap {
		res += int64(abs(median-v.val) * v.count)
	}
	return res
}
