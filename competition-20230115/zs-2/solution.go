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
