package main

func countAlternatingSubarrays(nums []int) int64 {
	// 从左到右找最长的交替子数组
	res := int64(0)
	n := 0
	pre := -1
	for _, v := range nums {
		if v == pre {
			res += int64(n * (n + 1) / 2)
			n = 1
		} else {
			n++
		}
		pre = v
	}
	res += int64(n * (n + 1) / 2)
	return res
}
