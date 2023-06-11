package main

func semiOrderedPermutation(nums []int) int {
	// 先找到1和n的位置
	index1, indexn := 0, 0
	n := len(nums)
	for i, v := range nums {
		if v == 1 {
			index1 = i
		} else if v == n {
			indexn = i
		}
	}

	sum := (index1 - 0) + (n - 1 - indexn)
	if index1 > indexn {
		sum--
	}
	return sum
}
