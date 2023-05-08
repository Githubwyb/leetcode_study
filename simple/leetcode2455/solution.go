package main

func averageValue(nums []int) int {
	result := 0
	count := 0
	for _, v := range nums {
		if v%6 == 0 {
			result += v
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return (result / count)
}
