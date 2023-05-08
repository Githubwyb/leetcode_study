package main

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

func findTheLongestBalancedSubstring(s string) int {
	res := 0
	pre, cnt := 0, 0 // pre为0的数量，cnt为当前字符连续的数量
	n := len(s)
	for i, v := range s {
		cnt++
		if i == n-1 || byte(v) != s[i+1] {
			if v == '0' {
				// 0到1变化，pre等于cnt
				pre = cnt
			} else {
				// 1到0变化，判断是否可以更新res
				res = max(res, min(pre, cnt)*2)
			}
			// 变化重新统计
			cnt = 0
		}
	}
	return res
}
