package main

import (
	"fmt"
)

func getKey(a, b, g int) string {
	return fmt.Sprintf("%d#%d#%d", a, b, g)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMaxScore(a, b, g int, in []int, dp *map[string]int) int {
	if g <= 0 || a > b {
		return 0
	}
	if a == b && g >= 1 {
		return in[a]
	}
	if a+1 == b && g >= 1 {
		return max(in[a], in[b])
	}
	if v, ok := (*dp)[getKey(a, b, g)]; ok {
		return v
	}

	result := max(in[a]+getMaxScore(a+2, b, g-1, in, dp), getMaxScore(a+1, b, g, in, dp))
	(*dp)[getKey(a, b, g)] = result
	return result
}

func applesSum(k int, in []int) int {
	// 记忆化搜索，取f(a，b, g)为[a, b]之间（可以吃a和b）吃g个的最大值
	// f(a, b, g) = max(in[a] + f(a+2, b, g-1), f(a+1, b, g))
	// a == b && g >= 1 return a
	// a == b-1 && g >= 1 return max(a, b)
	dp := make(map[string]int)
	if len(in) == 1 && k >= 1 {
		return in[0]
	}
	return max(getMaxScore(0, len(in)-2, k, in, &dp), getMaxScore(1, len(in)-1, k, in, &dp))
}

// 动态规划
func getMaxScore1(a, b, g int, in []int) int {
	dp := make([][]int, len(in)) // 0-n, k    (i, n)之间吃k个的最大分数

	for i := range dp {
		dp[i] = make([]int, g)
	}
	// 从b向前计算，将dp填满，然后返回结果
	return 0
}

func applesSum1(k int, in []int) int {
	// 动态规划，取f(a，b, g)为[a, b]之间（可以吃a和b）吃g个的最大值
	// f(a, b, g) = max(in[a] + f(a+2, b, g-1), f(a+1, b, g))
	// a == b && g >= 1 return a
	// a == b-1 && g >= 1 return max(a, b)
	return max(getMaxScore1(0, len(in)-2, k, in), getMaxScore1(1, len(in)-1, k, in))
}
