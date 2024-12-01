package common

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minInts(nums ...int) int {
	ans := nums[0]
	for _, v := range nums[1:] {
		if ans > v {
			ans = v
		}
	}
	return ans
}

func maxInts(nums ...int) int {
	ans := nums[0]
	for _, v := range nums[1:] {
		ans = max(ans, v)
	}
	return ans
}

func bool2int(x bool) int {
	if x {
		return 1
	}
	return 0
}
