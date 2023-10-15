package main

func longestAlternatingSubarray(nums []int, threshold int) int {
	n := len(nums)
	res := 0
	// 枚举左端点
	for l := 0; l < n && n-l > res; l++ {
		if nums[l]%2 != 0 || nums[l] > threshold {
			continue
		}

		r := l + 1
		for t := 1; r < n && nums[r] <= threshold && nums[r]%2 == t; r, t = r+1, t^1 {
		}
		if r-l > res {
			res = r - l
		}
		l = r - 1
	}
	return res
}
