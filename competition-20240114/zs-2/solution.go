package main

import (
	"strings"
)

func beautifulIndices(s string, a string, b string, k int) []int {
	getAllIdxs := func(s, sub string) []int {
		res := []int{}
		idx := strings.Index(s, sub)
		last := 0
		for idx != -1 {
			res = append(res, idx+last)
			last += idx + 1
			idx = strings.Index(s[last:], sub)
		}
		return res
	}
	aIdxs := getAllIdxs(s, a)
	bIdxs := getAllIdxs(s, b)
	res := make([]int, 0, 1)
	n := len(bIdxs)
	if n == 0 {
		return res
	}
	idx := 0
	for _, i := range aIdxs {
		for i-bIdxs[idx] > k {
			idx++
			if idx == n {
				return res
			}
		}
		if i-bIdxs[idx] <= k && bIdxs[idx]-i <= k {
			res = append(res, i)
		}
	}
	return res
}
