package main

import "sort"

func deleteGreatestValue(grid [][]int) int {
	for i := range grid {
		sort.Ints(grid[i])
	}

	result := 0
	for i := range grid[0] {
		max := grid[0][i]
		for j := range grid {
			if max < grid[j][i] {
				max = grid[j][i]
			}
		}
		result += max
	}

	return result
}
