package main

func countCompleteComponents(n int, edges [][]int) int {
	// 使用并查集合并所有的点
	// 并查集模板
	uf := make([]int, n)
	for i := range uf {
		uf[i] = i
	}
	find := func(x int) int {
		ap := uf[x]
		// 找到最终节点
		for ap != uf[ap] {
			ap = uf[ap]
		}
		// 沿途都赋值最终节点
		for x != ap {
			uf[x], x = ap, uf[x]
		}
		return ap
	}
	// 把a的子集合并到b上，如果b是树根节点，a的所有子节点查找都会查找到b
	merge := func(a, b int) {
		uf[find(a)] = find(b)
	}

	// 并查集整理好
	for _, v := range edges {
		merge(v[0], v[1])
	}

	// 并查集的每个集合统计点数和边数
	pCnt := make([]int, n)
	eCnt := make([]int, n)
	for i := 0; i < n; i++ {
		pCnt[find(i)]++
	}
	for _, v := range edges {
		eCnt[find(v[0])]++
	}

	// 判断数量
	res := 0
	for i, v := range pCnt {
		if v == 0 {
			continue
		}
		if v*(v-1)/2 == eCnt[i] {
			res++
		}
	}
	return res
}

func countCompleteComponents1(n int, edges [][]int) int {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[y] = append(g[y], x)
		g[x] = append(g[x], y)
	}

	vis := make([]bool, n)
	pCnt := 0
	eCnt := 0
	var dfs func(x int)
	dfs = func(x int) {
		vis[x] = true
		pCnt++
		eCnt += len(g[x])
		for _, v := range g[x] {
			if !vis[v] {
				dfs(v)
			}
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		if !vis[i] {
			eCnt, pCnt = 0, 0
			dfs(i)
			if pCnt*(pCnt-1) == eCnt {
				res++
			}
		}
	}
	return res
}
