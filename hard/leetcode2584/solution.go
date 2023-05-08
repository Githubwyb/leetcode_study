package main

import "math"

func getPrimeFactor(n int, handle func(prime int)) {
	max := int(math.Sqrt(float64(n))) + 1
	for i := 2; i < max && i < n; i++ {
		if n%i == 0 {
			handle(i)
			n /= i
			i = 1
		}
	}
	if n > 1 {
		handle(n)
	}
}

func findValidSplit(nums []int) int {
	left := make(map[int]int)
	right := make([]int, len(nums)) // 左端点为i的右端点最大值
	for i, v := range nums {
		getPrimeFactor(v, func(prime int) {
			if l, ok := left[prime]; ok {
				right[l] = i
			} else {
				left[prime] = i
			}
		})
	}
	maxR := 0
	for l, r := range right {
		if l > maxR {
			return maxR
		}
		if r > maxR {
			maxR = r
		}
	}
	return -1
}
