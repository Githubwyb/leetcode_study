package main

func reverse(num int) int {
	result := 0
	for ; num != 0; num /= 10 {
		result *= 10
		result += num % 10
	}
	return result
}

func countDistinctIntegers(nums []int) int {
	numMap := make(map[int]bool)
	for _, v := range nums {
		numMap[v] = true
		numMap[reverse(v)] = true
	}
	return len(numMap)
}
