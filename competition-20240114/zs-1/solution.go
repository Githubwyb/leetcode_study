package main

func maxFrequencyElements(nums []int) int {
	vMap := map[int]int{}
	maxFre := 0
	for _, v := range nums {
		vMap[v]++
		if vMap[v] > maxFre {
			maxFre = vMap[v]
		}
	}
	res := 0
	for _, v := range vMap {
		if v == maxFre {
			res += v
		}
	}
	return res
}
