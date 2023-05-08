package main

import (
	"math"
)

func maxDivScore(nums []int, divisors []int) int {
	res, rc := math.MaxInt, -1
	for _, v := range divisors {
		cnt := 0
		for _, v1 := range nums {
			if v1%v == 0 {
				cnt++
			}
		}
		if cnt > rc || (rc == cnt && res > v) {
			res, rc = v, cnt
		}
	}
	return res
}
