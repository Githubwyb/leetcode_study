package main

func countSubmatrices(grid [][]int, k int) int {
	for _, v := range grid {
		s := 0
		for j, v1 := range v {
			s += v1
			v[j] = s
		}
	}
	s := make([]int, len(grid[0]))
	res := 0
	for _, v := range grid {
		for j, v1 := range v {
			s[j] += v1
			if s[j] > k {
				break
			}
			res++
		}
	}
	return res
}
