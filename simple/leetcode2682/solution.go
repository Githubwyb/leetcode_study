package main

func circularGameLosers(n int, k int) []int {
	fetchMap := make([]bool, n)
	for cur, i := 0, k; !fetchMap[cur]; i += k {
		fetchMap[cur] = true
		// 下一个接球的人
		cur = (cur + i) % n
	}
	res := make([]int, 0, n)
	for i, c := range fetchMap {
		if !c {
			res = append(res, i+1)
		}
	}
	return res
}
