package main

func findChampion(grid [][]int) int {
	chp := 0
	n := len(grid)
	for {
		find := false
		for i := 0; i < n; i++ {
			if i != chp && grid[chp][i] != 1 {
				chp = i
				find = true
				break
			}
		}
		if !find {
			break
		}
	}
	return chp
}
