package main

func evenOddBit(n int) []int {
	res := []int{0, 0}
	isOdd := false
	for n > 0 {
		if n&0x01 == 1 {
			if isOdd {
				res[1]++
			} else {
				res[0]++
			}
		}
		n >>= 1
		isOdd = !isOdd
	}
	return res
}
