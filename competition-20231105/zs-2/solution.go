package main

func findChampion(n int, edges [][]int) int {
	state := make([][]int, n)
	for i := range state {
		state[i] = make([]int, n)
	}
	for _, v := range edges {
		x, y := v[0], v[1]
		state[x][y] = 2
		state[y][x] = 1
	}
	chp := -1
	for i := 0; i < n; i++ {
		isDefeat := false
		for _, v := range state[i] {
			if v == 1 {
				isDefeat = true
				break
			}
		}
		if !isDefeat {
			if chp != -1 {
				return -1
			}
			chp = i
		}
	}
	return chp
}
