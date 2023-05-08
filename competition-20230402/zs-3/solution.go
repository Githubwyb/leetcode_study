package main

import "sort"

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	res := 0
	n := len(reward1)
	diff := make([]int, n)
	// 全部让老鼠2吃
	for i, v := range reward2 {
		res += v
		diff[i] = reward1[i] - reward2[i]
	}
	sort.Ints(diff)
	for i := 0; i < k; i++ {
		res += diff[n-1-i]
	}
	return res
}
