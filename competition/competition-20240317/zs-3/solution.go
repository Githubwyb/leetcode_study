package main

import (
	"math"
	"sort"
)

func minimumDeletions(word string, k int) int {
	cMap := map[byte]int{}
	for _, v := range word {
		cMap[byte(v)]++
	}

	n := len(cMap)
	cArr := make([]int, 0, n)
	for _, v := range cMap {
		cArr = append(cArr, v)
	}

	sort.Ints(cArr)

	// 每个v作为最小值
	res := math.MaxInt
	delAll := 0
	for i, v := range cArr {
		idx := sort.Search(n-i, func(j int) bool { return cArr[j+i] > v+k }) + i
		resTmp := delAll
		for j := idx; j < n; j++ {
			resTmp += cArr[j] - (v + k)
		}
		if resTmp < res {
			res = resTmp
		}
		delAll += v
	}
	return res
}
