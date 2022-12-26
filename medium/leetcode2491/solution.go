package main

import "sort"

func dividePlayers(skill []int) int64 {
	n := len(skill)
	sort.Ints(skill)
	tmp := skill[0] + skill[n-1]
	var result int64 = 0
	for i := 0; i < n/2; i++ {
		if tmp != skill[i] + skill[n-i-1] {
			return -1
		}
		result += int64(skill[i]*skill[n-i-1])
	}
	return result
}
