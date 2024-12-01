package main

func maxInts(nums ...int) int {
	ans := nums[0]
	for _, v := range nums[1:] {
		ans = max(ans, v)
	}
	return ans
}

func maxTargetNodes(edges1 [][]int, edges2 [][]int, k int) []int {
	// 获取树上每个节点的其他节点到其距离小于等于k的数量数组
	getKCnt := func(edges [][]int, k int) []int {
		n := len(edges) + 1
		tree := make([][]int, n)
		for _, v := range edges {
			x, y := v[0], v[1]
			tree[x] = append(tree[x], y)
			tree[y] = append(tree[y], x)
		}

		// 到i节点距离为k的有几个点
		var dfs func(tree [][]int, i int, last int, k int) int
		dfs = func(tree [][]int, i int, last int, k int) int {
			if k < 0 {
				return 0
			}
			// 自己到自己计算为1
			ans := 1
			for _, j := range tree[i] {
				if j == last {
					// 防止回退
					continue
				}
				ans += dfs(tree, j, i, k-1)
			}
			return ans
		}

		ans := make([]int, n)
		for i := range tree {
			ans[i] = dfs(tree, i, -1, k)
		}
		return ans
	}
	tmp := getKCnt(edges2, k-1)
	maxCnt := maxInts(tmp...)
	ans := getKCnt(edges1, k)
	for i := range ans {
		ans[i] += maxCnt
	}

	return ans
}
