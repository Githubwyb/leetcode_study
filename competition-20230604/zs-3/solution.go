package main

func matrixSumQueries(n int, queries [][]int) int64 {
	// 倒序处理queries
	gMap := [2]map[int]bool{{}, {}}
	var sum int64 = 0
	for i := len(queries) - 1; i >= 0; i-- {
		t, i, v := queries[i][0], queries[i][1], queries[i][2]
		if gMap[t][i] {
			// 已经操作过，跳过
			continue
		}
		gMap[t][i] = true
		sum += int64(v * (n - len(gMap[t^1])))
	}
	return sum
}
