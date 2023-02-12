package main

import (
	"fmt"
	"strings"
)

func substringXorQueries(s string, queries [][]int) [][]int {
	result := make([][]int, len(queries))
	for i, v := range queries {
		findStr := fmt.Sprintf("%b", v[1]^v[0])
		l := strings.Index(s, findStr)
		if l == -1 {
			result[i] = []int{l, l}
		} else {
			result[i] = []int{l, l + len(findStr) - 1}
		}
	}
	return result
}

func substringXorQueries1(s string, queries [][]int) [][]int {
	type pair struct{ l, r int }
	checkMap := make(map[int]pair)
	for l := range s {
		for r, x := l, 0; r < l+30 && r < len(s); r++ {
			x = (x << 1) | int(s[r]&1)
			if p, ok := checkMap[x]; !ok || r-l < p.r-p.l {
				checkMap[x] = pair{l, r}
			}
		}
	}
	result := make([][]int, len(queries))
	notFound := []int{-1, -1}
	for i, v := range queries {
		if p, ok := checkMap[v[1]^v[0]]; ok {
			result[i] = []int{p.l, p.r}
		} else {
			result[i] = notFound
		}
	}
	return result
}
