package main

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func leftRigthDifference(nums []int) []int {
	result := make([]int, len(nums))
	rsum := 0
	for _, v := range nums {
		rsum += v
	}
	lsum := 0
	for i, v := range nums {
		rsum -= v
		result[i] = abs(lsum - rsum)
		lsum += v
	}
	return result
}
