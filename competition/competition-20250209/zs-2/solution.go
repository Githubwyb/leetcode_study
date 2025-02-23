package main

func assignElements(groups []int, elements []int) []int {
	n := len(groups)
	res := make([]int, n, n)
	for i := range res {
		res[i] = -1
	}
	m := map[int]bool{}
	for i := range res {
		m[i] = true
	}
	for i, v := range elements {
		for j := range m {
			v1 := groups[j]
			if v1%v == 0 {
				res[j] = i
				delete(m, j)
			}
		}
	}
	return res
}
