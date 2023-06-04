package main

func rowAndMaximumOnes(mat [][]int) []int {
	res := []int{0, 0}
	for i, v := range mat {
		cnt := 0
		for _, v1 := range v {
			if v1 == 1 {
				cnt++
			}
		}
		if cnt > res[1] {
			res[0], res[1] = i, cnt
		}
	}
	return res
}
