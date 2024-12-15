package main

func findPeaks(mountain []int) []int {
	res := make([]int, 0)
	st := 0
	for i, v := range mountain {
		if i == 0 {
			continue
		}
		if v > mountain[i-1] {
			st = 1
			continue
		}
		if v < mountain[i-1] && st == 1 {
			res = append(res, i-1)
		}
		st = 0
	}
	return res
}
