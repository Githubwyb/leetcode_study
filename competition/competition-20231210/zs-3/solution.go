package main

func countSubarrays(nums []int, k int) int64 {
	max := 0
	idxs := []int{}
	for i, v := range nums {
		if v > max {
			max = v
			idxs = []int{i}
		} else if v == max {
			idxs = append(idxs, i)
		}
	}

	res := int64(0)
	n1 := len(idxs)
	n := len(nums)
	last := 0
	for i := 0; i <= n1-k; i++ {
		l := idxs[i]
		r := idxs[i+k-1]
		res += int64(l-last+1) * int64(n-r)
		last = l + 1
	}
	return res
}
