package main

import "math"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func paintWalls(cost []int, time []int) int {
	n := len(cost)

	mMap := make([][]int, n)
	for i := range mMap {
		// 最大的时间不会超过n，因为超过n直接返回0不会进map进行判断
		// 所有时间加上n作为偏移
		mMap[i] = make([]int, 2*n+1)
		for j := range mMap[i] {
			mMap[i][j] = -1
		}
	}
	// 选和不选代价和时间花费不一样，可以认为付费的时间和花费是正的，免费的时间是负的，花费为0
	// 最终求得是时间不为负数的花费最小值
	var dfs func(i, t int) int
	dfs = func(i, t int) int {
		if t-n > i {
			return 0
		}
		if i < 0 {
			return math.MaxInt / 2
		}
		if mMap[i][t] >= 0 {
			return mMap[i][t]
		}
		// 为当前免费和当前不免费的最小值
		res := min(dfs(i-1, t-1), dfs(i-1, t+time[i])+cost[i])
		mMap[i][t] = res
		return res
	}
	res := dfs(n-1, n)
	return res
}

func paintWalls1(cost []int, time []int) int {
	n := len(cost)
	// 偏移n
	f := make([]int, 2*n+1)
	// 偏移前可以认为是[-n,0)都是超级大
	for i := 0; i < n; i++ {
		f[i] = math.MaxInt / 2
	}

	f1 := make([]int, 2*n+1)
	for i := 0; i < n; i++ {
		f1, f = f, f1
		for t := i + 1; t < n+i+1; t++ {
			f[t] = min(f1[t-1], f1[min(t+time[i], 2*n)]+cost[i])
		}
	}
	return f[n]
}

func paintWalls2(cost []int, time []int) int {
	n := len(cost)
	// 偏移n
	f := make([]int, 2*n+1)
	// 偏移前可以认为是[-n,0)都是超级大
	for i := 0; i < n; i++ {
		f[i] = math.MaxInt / 2
	}

	f1 := make([]int, 2*n+1)
	for i := 0; i < n; i++ {
		f1, f = f, f1
		fLen := 2*n + 1 - i - 1
		f = f[:fLen]
		for j := range f[:n] {
			f[j] = min(f1[j], f1[min(j+1+time[i], fLen-1)]+cost[i])
		}
	}
	return f[0]
}
