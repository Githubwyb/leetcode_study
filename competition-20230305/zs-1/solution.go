package main

func passThePillow(n int, time int) int {
	tmp := time % (n + n - 2)
	if tmp >= n-1 {
		return n - (tmp - n + 1)
	}
	return tmp + 1
}
