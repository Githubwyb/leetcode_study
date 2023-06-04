package main

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func differenceOfDistinctValues(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	tl := make([]int, 0, m)
	getAns := func(is, js int) {
		tlMap := make(map[int]int)
		tl = tl[:0]
		// 第一遍统计左上角的数量并记录在tl中
		for i, j := is, js; i < m && j < n; i, j = i+1, j+1 {
			tl = append(tl, len(tlMap))
			tlMap[grid[i][j]]++
		}
		// 第二遍计算右下角数量
		for i, j, k := is, js, 0; i < m && j < n; i, j, k = i+1, j+1, k+1 {
			v := grid[i][j]
			if tlMap[v] <= 1 {
				delete(tlMap, v)
			} else {
				tlMap[v]--
			}
			res[i][j] = abs(len(tlMap) - tl[k])
		}
	}
	for i := 0; i < n; i++ {
		getAns(0, i)
	}
	for i := 1; i < m; i++ {
		getAns(i, 0)
	}
	return res
}
