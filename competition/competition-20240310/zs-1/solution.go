package main

import "sort"

func minimumBoxes(nums []int, divisors []int) int {
	sort.Slice(divisors, func(i, j int) bool { return divisors[i] > divisors[j] })
	sum := 0
	for _, v := range nums {
		sum += v
	}
	for i, v := range divisors {
		sum -= v
		if sum <= 0 {
			return i + 1
		}
	}

	return len(divisors)
}
