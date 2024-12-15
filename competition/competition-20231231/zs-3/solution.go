package main

import (
	"fmt"
	"sort"
)

func maximumLength(s string) int {
	aMap := map[byte][]int{}
	for i, v1 := range s {
		v := byte(v1)
		aMap[v] = append(aMap[v], i)
	}

	res := 0
	for _, v := range aMap {
		n := len(v)
		if n-2 <= res {
			continue
		}

		lenInts := make([]int, 0, n)
		pre := -2
		s := 0
		for _, v1 := range v {
			if v1-1 == pre || pre == -2 {
				s++
			} else {
				lenInts = append(lenInts, s)
				s = 1
			}
			pre = v1
		}
		if s != 0 {
			lenInts = append(lenInts, s)
		}
		fmt.Println(lenInts)
		maxNum := sort.Search(n-1, func(i int) bool {
			cnt := 0
			for _, v := range lenInts {
				if v >= i {
					cnt += v - i + 1
					if cnt >= 3 {
						return false
					}
				}
			}
			return true
		}) - 1
		if maxNum > res {
			res = maxNum
		}
	}
	if res == 0 {
		return -1
	}
	return res
}
