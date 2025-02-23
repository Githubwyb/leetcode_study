package main

import (
	"sort"
)

type sortI struct {
	grid [][]int
	i    int
	j    int
	n    int
}

func (x *sortI) Len() int { return x.n }

// 为true，i向前；false，j向前。要满足相等时返回false
func (x *sortI) Less(i, j int) bool {
	if x.j == 0 {
		return x.grid[x.i+i][x.j+i] > x.grid[x.i+j][x.j+j]
	}
	return x.grid[x.i+i][x.j+i] < x.grid[x.i+j][x.j+j]
}
func (x *sortI) Swap(i, j int) {
	x.grid[x.i+i][x.j+i], x.grid[x.i+j][x.j+j] = x.grid[x.i+j][x.j+j], x.grid[x.i+i][x.j+i]
}

func sortMatrix(grid [][]int) [][]int {
	n := len(grid)

	// 对角线公式 i-j+n-1=k
	for k := 0; k < 2*n-1; k++ {
		i := max(0, k-n+1)
		j := i + n - 1 - k
		l := min(n-1, 2*n-2-k) - j + 1
		sort.Sort(&sortI{grid, i, j, l})
	}

	return grid
}
