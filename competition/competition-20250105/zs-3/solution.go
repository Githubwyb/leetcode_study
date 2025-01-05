package main

import (
	"sort"
)

func maximumCoins(coins [][]int, k int) int64 {
	sort.Slice(coins, func(i, j int) bool {
		return coins[i][0] < coins[j][0]
	})
	n := len(coins)
	res := int64(0)
	left, right := 0, 0 // right不可达
	coverLeft, coverRight := int64(0), int64(0)
	for _, v := range coins {
		// 对每个区间作为左端点和右端点进行计算
		s, e := v[0], v[1]
		le, rs := s+k-1, e-k+1 // le为作为左端点的终点，rs为作为右端点的起点
		// 作为左端点，找可以完全覆盖到的所有段
		for ; right < n && coins[right][1] <= le; right++ {
			coverLeft += int64((coins[right][1] - coins[right][0] + 1) * coins[right][2])
		}
		tmp := coverLeft
		if right < n && coins[right][0] <= le {
			tmp += int64(coins[right][2] * (le - coins[right][0] + 1))
		}
		res = max(tmp, res)
		coverLeft -= int64(v[2] * (v[1] - v[0] + 1))

		// 作为右端点，找可以完全覆盖到的所有段
		coverRight += int64(v[2] * (v[1] - v[0] + 1))
		for ; left < n && coins[left][0] < rs; left++ {
			coverRight -= int64((coins[left][1] - coins[left][0] + 1) * coins[left][2])
		}
		tmp = coverRight
		if left > 0 && coins[left-1][1] >= rs {
			tmp += int64(coins[left-1][2] * (coins[left-1][1] - rs + 1))
		}
		res = max(tmp, res)

	}
	return res
}
