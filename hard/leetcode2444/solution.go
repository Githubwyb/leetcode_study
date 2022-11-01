package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func countSubarrays(nums []int, minK int, maxK int) int64 {
	s, minI, maxI := -1, -1, -1
	var result int64 = 0
	for e, v := range nums {
		if v < minK || v > maxK {
			s = e // 子数组不包含s
			continue
		}
		if v == minK {
			minI = e
		}
		if v == maxK {
			maxI = e
		}
		// 1 2 3 4 5
		// s   i   e
		result += int64(max(min(minI, maxI)-s, 0))
	}
	return result
}
