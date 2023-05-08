package common

func test() {
	n := 10

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

	merge(0, 0)
}

// 函数外模板
type unionFind []int

func initUnionFind(n int) unionFind {
	u := make(unionFind, n)
	for i := range u {
		u[i] = i
	}
	return u
}

func (u unionFind) find(a int) int {
	ap := u[a]
	// 找到最终节点
	for ap != u[ap] {
		ap = u[ap]
	}
	// 沿途都赋值最终节点
	for a != ap {
		u[a], a = ap, u[a]
	}
	return ap
}

// 把a的子集合并到b上，如果b是树根节点，a的所有子节点查找都会查找到b
func (u unionFind) merge(a, b int) {
	u[u.find(b)] = u.find(a)
}
