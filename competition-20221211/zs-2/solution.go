package main

import (
	"sort"
)

func longestSquareStreak(nums []int) int {
	dp := make(map[int]int)
	sort.Ints(nums)
	ans := 1
	for _, v := range nums {
		c := 0
		var ok bool
		if c, ok = dp[v]; ok {
			c += 1
		} else {
			c = 1
		}
		dp[v*v] = c
		if c > ans {
			ans = c
		}
	}

	if ans == 1 {
		return -1
	}
	return ans
}

func longestSquareStreak1(nums []int) int {
	set := make(map[int]bool)
	for _, v := range nums {
		set[v] = true
	}

	ans := 1
	for _, v := range nums {
		cnt := 1
		for v *= v; set[v]; v *= v {
			cnt++
		}
		if cnt > ans {
			ans = cnt
		}
	}
	if ans == 1 {
		return -1
	}
	return ans
}
