package main

var psMap = map[int]bool{}

func init() {
	var primes = make([]int, 0, 9593)
	oulerPrimes(1e5+1, &primes)
}

func oulerPrimes(mx int, primes *[]int) {
	flag := make([]bool, mx+1) // 标记数有没有被筛掉，false就是没有
	for i := 2; i < mx+1; i++ {
		if !flag[i] {
			// 数没有被比自己小的数筛掉，就代表是质数
			*primes = append(*primes, i)
			psMap[i] = true
		}
		for _, v := range *primes {
			if i*v > mx {
				break
			}
			// 每一个数都作为因子乘以比自己小的素数筛掉后面的数
			flag[i*v] = true
			if i%v == 0 {
				// 减少时间复杂度的关键算法
				// 12 = 2 * 3 * 2，i = 4时，只排了8就退出了，因为6会将12排除
				// 也就是，假设v可以整除i即i = kv，有某个数为x = mi = kmv
				//        那么存在一个数 i < km < x可以把x排掉，用i乘以所有的质数去排除就没什么意义了，提前退出减少时间复杂度
				break
			}
		}
	}
}

func countPaths(n int, edges [][]int) int64 {
	nodeMap := make([][]int, n+1)

	for _, v := range edges {
		x, y := v[0], v[1]
		nodeMap[x] = append(nodeMap[x], y)
		nodeMap[y] = append(nodeMap[y], x)
	}

	// 遍历每个节点，找每个节点各个路径有多少个点
	// 并查集模板
	uf := make([]int, n+1)
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

	// 找当前节点到质数之间有几个点，返回点数量和遇到的质数
	var countPoint func(cur, last int) int
	countPoint = func(cur int, last int) int {
		if psMap[cur] {
			return 0
		}
		merge(cur, last)

		cnt := 1 // 当前点
		for _, v := range nodeMap[cur] {
			if v == last {
				continue
			}
			cnt += countPoint(v, cur)
		}
		return cnt
	}

	var res int64
	cntMaps := make([]int64, n+1)
	for i := 1; i < n+1; i++ {
		// 遍历所有质数
		if !psMap[i] {
			continue
		}

		// 统计质数周围有几个连通块
		var sum int64 = 0
		for _, v1 := range nodeMap[i] {
			p := find(v1)
			c := cntMaps[p]
			if c == 0 {
				c = int64(countPoint(v1, v1))
				p = find(v1)
				cntMaps[p] = c
			}
			res += sum * c
			sum += c
		}
		res += sum
	}
	return res
}
