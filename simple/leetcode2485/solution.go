package main

import "math"

func pivotInteger(n int) int {
	l, r := 1, n
	tmp := 0 // 加左边减右边
	// 两边遍历直到中间
	for l != r {
		if tmp > 0 {
			tmp -= r
			r--
		} else {
			tmp += l
			l++
		}
	}
	if tmp == 0 {
		return l
	}
	return -1
}

func pivotInteger1(n int) int {
	m := n * (n + 1) / 2
	for i := 1; i <= n; i++ {
		if i*i == m {
			return i
		}
	}
	return -1
}

func pivotInteger2(n int) int {
	m := n * (n + 1) / 2
	x := int(math.Sqrt(float64(m)))
	if x * x == m {
		return x
	}
	return -1
}
