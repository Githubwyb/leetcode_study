package main

func pivotInteger(n int) int {
	l, r := 1, n
	ls, rs := 0, 0
	for l != r {
		if ls > rs {
			rs += r
			r--
		} else {
			ls += l
			l++
		}
	}
	if ls == rs {
		return l
	}
	return -1
}
