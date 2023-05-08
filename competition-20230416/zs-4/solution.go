package main

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
	rel := make([][]int, n)
	for _, v := range edges {
		x, y := v[0], v[1]
		rel[x] = append(rel[x], y)
		rel[y] = append(rel[y], x)
	}
	// 因为只有一条路，所以深度优先
	// 起点、终点、上一步（因为不存在环路，所以判断不回来即可）
	var dfs func(s, e, last int) bool
	cnt := make(map[int]int)
	dfs = func(s, e, last int) bool {
		if s == e {
			cnt[e]++
			return true
		}
		for _, v := range rel[s] {
			if v != last && dfs(v, e, s) {
				cnt[s]++
				return true
			}
		}
		return false
	}
	for _, v := range trips {
		dfs(v[0], v[1], -1)
	}

	// 单向，使用last即可防止回路
	// 返回选和不选的两个大小
	var dfs1 func(s, last int) (nch, ch int)
	dfs1 = func(s, last int) (nch int, ch int) {
		nch = cnt[s] * price[s]
		ch = nch / 2
		for _, v := range rel[s] {
			if v == last {
				continue
			}
			vnch, vch := dfs1(v, s)
			ch += vnch // 当前选中了就只能选择没选中的
			nch += min(vnch, vch)
		}
		return
	}
	// 随便找一个要走的点开始遍历
	nch, ch := dfs1(trips[0][0], -1)
	return min(nch, ch)
}

func minimumTotalPrice1(n int, edges [][]int, price []int, trips [][]int) int {
	rel := make([][]int, n)
	for _, v := range edges {
		x, y := v[0], v[1]
		rel[x] = append(rel[x], y)
		rel[y] = append(rel[y], x)
	}

	// 树上差分计算cnt
	diff := make([]int, n)
	// 把查询都取出来，一次遍历全部查一遍
	qs := make(map[int][]int)
	for _, v := range trips {
		x, y := v[0], v[1]
		qs[x] = append(qs[x], y)
		if x != y {
			qs[y] = append(qs[y], x)
		}
	}
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
	// Tarjan算法计算公共祖先
	color := make([]bool, n)
	father := make([]int, n) // 每个节点的父节点
	var tarjan func(a, fa int)
	tarjan = func(a, fa int) {
		father[a] = fa
		for _, v := range rel[a] {
			if v == fa {
				continue
			}
			tarjan(v, a)
			// 进去出来后，将v为根节点的子树设置公共祖先为a
			merge(v, a)
		}

		// 查一下有没有要求的LCA
		for _, v := range qs[a] {
			if v != a && !color[v] {
				// 自己走到自己是可以计算的，要判断
				// v还没走到，继续
				continue
			}
			lca := find(v)
			diff[a]++
			diff[v]++
			diff[lca]--
			if lcaFa := father[lca]; lcaFa >= 0 {
				diff[lcaFa]--
			}
		}
		color[a] = true // a被灌了岩浆，也就是a的子树走完了，要向上走了
	}
	// 从0向下走
	tarjan(0, -1)

	// dfs，同时计算差分
	// 返回选和不选的两个大小
	var dfs1 func(s, fa int) (nch, ch, cnt int)
	dfs1 = func(s, fa int) (nch, ch, cnt int) {
		cnt = diff[s]
		for _, v := range rel[s] {
			if v == fa {
				continue
			}
			vnch, vch, ccnt := dfs1(v, s)
			ch += vnch // 当前选中了就只能选择没选中的
			nch += min(vnch, vch)
			cnt += ccnt // 当前节点cnt为自己的差分加上所有子节点cnt之和
		}
		nch += cnt * price[s]
		ch += cnt * price[s] / 2
		return
	}
	// 从根节点遍历
	nch, ch, _ := dfs1(0, -1)
	return min(nch, ch)
}
