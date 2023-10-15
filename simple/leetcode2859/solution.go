package main

import "math/bits"

func sumIndicesWithKSetBits(nums []int, k int) int {
	res := 0
	for i, x := range nums {
		if bits.OnesCount(uint(i)) == k {
			res += x
		}
	}
	return res
}
