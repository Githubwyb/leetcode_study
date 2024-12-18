package main

import "sort"

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Slice(happiness, func(i, j int) bool { return happiness[i] > happiness[j] })

	res := int64(0)
	for i := 0; i < k; i++ {
		if happiness[i] <= i {
			break
		}
		if happiness[i] > i {
			res += int64(happiness[i] - i)
		}
	}
	return res
}
