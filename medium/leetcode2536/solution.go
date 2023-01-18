package main

func rangeAddQueries(n int, queries [][]int) [][]int {
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	for _, v := range queries {
		for i := v[0]; i <= v[2]; i++ {
			for j := v[1]; j <= v[3]; j++ {
				result[i][j]++
			}
		}
	}
	return result
}

func rangeAddQueries1(n int, queries [][]int) [][]int {
	// 防止越界
	result := make([][]int, n+1)
	for i := range result {
		result[i] = make([]int, n+1)
	}

	for _, v := range queries {
		r1, c1, r2, c2 := v[0], v[1], v[2]+1, v[3]+1
		result[r1][c1] += 1
		result[r1][c2] -= 1
		result[r2][c1] -= 1
		result[r2][c2] += 1
	}

	for i := range result {
		result[i] = result[i][:n]
		for j := range result[i] {
			// 每个点都是其自身和左上所有元素的和
			// 那么每个点都是自己+左邻居+上邻居-坐上邻居
			add := 0
			if i > 0 {
				add += result[i-1][j]
			}
			if j > 0 {
				add += result[i][j-1]
			}
			if i > 0 && j > 0 {
				add -= result[i-1][j-1]
			}
			result[i][j] += add
		}
	}
	return result[:n]
}
