package main

func findMatrix(nums []int) [][]int {
	arrMap := make(map[int]int)
	for _, v := range nums {
		arrMap[v]++
	}
	res := make([][]int, 0, 1)
	for len(arrMap) > 0 {
		arr := make([]int, 0, 1)
		for k, v := range arrMap {
			arr = append(arr, k)
			if v == 1 {
				delete(arrMap, k)
			} else {
				arrMap[k]--
			}
		}
		res = append(res, arr)
	}
	return res
}
