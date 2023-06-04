package main

import "sort"

func maxIncreasingCells(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	type pair struct{ i, j int }
	grp := make(map[int][]pair)
	rowMax := make([]int, m)
	colMax := make([]int, n)
	for i := range mat {
		for j, v := range mat[i] {
			grp[v] = append(grp[v], pair{i, j})
		}
	}
	keys := make([]int, 0, len(grp))
	for i := range grp {
		keys = append(keys, i)
	}
	sort.Ints(keys)

	res := 0
	for _, v := range keys {
		// 找这一行的最大值和这一列的最大值，因为从小到大遍历的，所以直接为其加一即可
		maxNums := make([]int, len(grp[v]))
		for i, p := range grp[v] {
			maxNums[i] = max(rowMax[p.i], colMax[p.j]) + 1
			res = max(res, maxNums[i])
		}
		for i, p := range grp[v] {
			rowMax[p.i] = max(rowMax[p.i], maxNums[i])
			colMax[p.j] = max(colMax[p.j], maxNums[i])
		}
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
