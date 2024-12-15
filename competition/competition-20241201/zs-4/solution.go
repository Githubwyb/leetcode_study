package main

func maxTargetNodes(edges1 [][]int, edges2 [][]int) []int {
	tagTree := func(edges [][]int) (rMap map[int]bool, gMap map[int]bool) {
		n := len(edges) + 1
		tree := make([][]int, n)
		for _, v := range edges {
			x, y := v[0], v[1]
			tree[x] = append(tree[x], y)
			tree[y] = append(tree[y], x)
		}
		rMap = make(map[int]bool)
		gMap = make(map[int]bool)
		var dfs func(tree [][]int, i, last, deep int)
		dfs = func(tree [][]int, i, last, deep int) {
			if deep%2 == 0 {
				rMap[i] = true
			} else {
				gMap[i] = true
			}
			for _, j := range tree[i] {
				if j == last {
					continue
				}
				dfs(tree, j, i, deep+1)
			}
		}
		dfs(tree, 0, -1, 0)
		return
	}
	r, g := tagTree(edges2)
	maxCnt := max(len(r), len(g))
	r, g = tagTree(edges1)
	ans := make([]int, len(edges1)+1)
	for i := range ans {
		if r[i] {
			ans[i] = len(r) + maxCnt
		} else {
			ans[i] = len(g) + maxCnt
		}
	}
	return ans
}
