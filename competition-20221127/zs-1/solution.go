package main

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
	for i := 1; i <= n; i++ {
		if 2*i*i == n*n+n {
			return i
		}
	}
	return -1
}
