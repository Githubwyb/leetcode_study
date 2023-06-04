package main

func maxMoves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	pre := make([]bool, m) // 前一列
	cur := make([]bool, m) // 当前列
	// 从第一列出发，都能到
	for i := 0; i < m; i++ {
		cur[i] = true
	}
	for i := 1; i < n; i++ {
		pre, cur = cur, pre
		cnt := 0
		for j := 0; j < m; j++ {
			v := grid[j][i]
			check := (j > 0 && pre[j-1] && v > grid[j-1][i-1]) ||
				(pre[j] && v > grid[j][i-1]) ||
				(j < m-1 && pre[j+1] && v > grid[j+1][i-1])
			if check {
				cnt++
			}
			cur[j] = check
		}
		if cnt == 0 {
			return i - 1
		}
	}
	return n - 1
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func maxMoves1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	last, cur := make([]int, m), make([]int, m)
	// 从最后一列向前递推，最后一列都是0步，不能再走了
	for i := n - 2; i >= 0; i-- {
		cur, last = last, cur
		for j := 0; j < m; j++ {
			v := grid[j][i]
			vc := 0
			// 取下面三个点可以到的点，取答案较大的那个
			if j > 0 && v < grid[j-1][i+1] {
				vc = max(vc, last[j-1]+1)
			}
			if v < grid[j][i+1] {
				vc = max(vc, last[j]+1)
			}
			if j < m-1 && v < grid[j+1][i+1] {
				vc = max(vc, last[j+1]+1)
			}
			cur[j] = vc
		}
	}
	res := 0
	for _, v := range cur {
		res = max(res, v)
	}
	return res
}
