package main

import "sort"

func minimumAddedCoins(coins []int, target int) int {
	// 从小到大取硬币
	// 假设存在一个可以满足的区间 [a, b)，新增一个硬币c，那么一定可以满足 [a+c, b+c)
	// 可以满足的区间就是两个区间取并集，假设b不在新区间内，就是要新增硬币来满足
	// 这个时候直接新增b比增加1带来的收益大
	addCount := 0
	sort.Ints(coins)
	r := 1 // 区间右端点，不包含
	i := 0 // 索引
	n := len(coins)
	for r <= target {
		if i < n && coins[i] <= r {
			r += coins[i]
			i++
		} else {
			addCount++
			r *= 2
		}
	}
	return addCount
}
