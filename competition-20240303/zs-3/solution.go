package main

import "math"

func minimumOperationsToWriteY(grid [][]int) int {
	beY := []int{0, 0, 0}
	noY := []int{0, 0, 0}
	n := len(grid)
	for i, v := range grid {
		for j, v1 := range v {
			if (i == j && i <= n/2) ||
				(i == n-j-1 && i < n/2) ||
				(j == n/2 && i > n/2) {
				beY[v1]++
			} else {
				noY[v1]++
			}
		}
	}

	res := math.MaxInt
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == i {
				continue
			}
			op := 0
			for k, v := range beY {
				if k == i {
					continue
				}
				op += v
			}
			for k, v := range noY {
				if k == j {
					continue
				}
				op += v
			}
			if op < res {
				res = op
			}
		}
	}
	return res
}
