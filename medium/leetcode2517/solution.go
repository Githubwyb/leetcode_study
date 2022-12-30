package main

import "sort"

func maximumTastiness(price []int, k int) int {
	sort.Ints(price)
	return sort.Search((price[len(price)-1]-price[0])/(k-1), func(d int) bool {
		// 二分查找从0开始，而间隔最小是1，那么换算一下，真实间隔是d+1
		d++
		cnt, x := 1, price[0]
		for _, v := range price[1:] {
			if v >= x+d {
				cnt++
				x = v
			}
		}
		// 因为间隔越小越容易满足，那么二分的条件就是找到第一个不符合条件的d，那么减一就是要求的最大值
		// 正好d是间隔减一，直接返回即可
		return cnt < k
	})
}
