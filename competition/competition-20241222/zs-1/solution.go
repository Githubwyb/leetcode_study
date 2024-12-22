package main

import "slices"

func minimumOperations(nums []int) int {
	m := map[int]bool{}
	for i, v := range slices.Backward(nums) {
		if m[v] {
			return i/3 + 1
		}
		m[v] = true
	}
	return 0
}
