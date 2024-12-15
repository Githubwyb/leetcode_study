package main

import (
	"math"
	"sort"
)

func shortestBeautifulSubstring(s string, k int) string {
	n := len(s)
	idxArr := make([]int, 0, n)
	for i, v := range s {
		if v == '1' {
			idxArr = append(idxArr, i)
		}
	}
	ansArr := []string{""}
	minLen := math.MaxInt
	for i := 0; i < len(idxArr)-k+1; i++ {
		st, ed := idxArr[i], idxArr[i+k-1]
		if ed-st > minLen {
			continue
		}
		if ed-st < minLen {
			minLen = ed - st
			ansArr = ansArr[:0]
		}
		ansArr = append(ansArr, s[st:ed+1])
	}
	sort.Strings(ansArr)
	return ansArr[0]
}
