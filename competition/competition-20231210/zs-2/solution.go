package main

import "math"

func getGoodIndices(variables [][]int, target int) []int {
	res := make([]int, 0, len(variables))
	for i, v := range variables {
		a, b, c, m := v[0], v[1], v[2], v[3]
		if target >= m {
			continue
		}
		a1 := int(math.Pow(float64(a%10), float64(b)))
		a2 := a1 % 10
		a3 := int(math.Pow(float64(a2%m), float64(c)))
		a4 := a3 % m
		if a4 == target {
			res = append(res, i)
		}
	}
	return res
}
