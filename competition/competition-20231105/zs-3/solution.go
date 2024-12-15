package main

func maximumScoreAfterOperations(edges [][]int, values []int) int64 {
	treeMap := map[int][]int{}
	for _, v := range edges {
		x, y := v[0], v[1]
		treeMap[x] = append(treeMap[x], y)
		treeMap[y] = append(treeMap[y], x)
	}
	// 子树健康，返回大于等于0，子树不健康返回-1
	var getMax func(root int, par int) (maxValue int64, sum int64)
	getMax = func(root int, par int) (maxValue int64, sum int64) {
		maxValue = -1
		health := int64(0)
		if len(treeMap[root]) == 1 {
			if values[root] > 0 {
				return 0, int64(values[root])
			}
			return -1, int64(values[root])
		}
		for _, v := range treeMap[root] {
			if v == par {
				continue
			}
			m, s := getMax(v, root)
			sum += s
			if m < 0 {
				// 子树无法自己健康
				health = -1
				continue
			}
			if health >= 0 {
				health += m
			}
		}
		if health == -1 {
			// 子树有不可以自己保持健康的，那么靠根节点保持
			if values[root] > 0 {
				maxValue = sum
				sum += int64(values[root])
				return
			}
			sum += int64(values[root])
			// 根节点也不能保持健康
			maxValue = -1
			return
		}
		// 子树可以自己保持健康
		// 根节点无法保持健康
		if values[root] == 0 {
			maxValue = health
			return
		}
		// 根节点可以保持，如果大于子树和，返回根节点，小于子树和，直接返回子树和
		if health+int64(values[root]) > sum {
			maxValue = health + int64(values[root])
			sum += int64(values[root])
			return
		}
		maxValue = sum
		sum += int64(values[root])
		return
	}

	res := int64(0)
	sum := int64(0)
	health := true
	for _, v := range treeMap[0] {
		m, s := getMax(v, 0)
		sum += s
		if !health {
			continue
		}
		if m < 0 {
			// 子树无法自己健康
			health = false
			continue
		}
		res += m
	}
	if !health {
		return sum
	}
	if values[0] == 0 {
		return res
	}
	if int64(values[0])+res <= sum {
		return sum
	}
	return res + int64(values[0])
}
