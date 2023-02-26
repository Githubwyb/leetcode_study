package main

func findTheString(lcp [][]int) string {
	n := len(lcp)
	s := make([]byte, n)
	var c byte = 'a'
	for i := range lcp {
		if s[i] == 0 {
			if c > 'z' {
				return ""
			}
			s[i] = c
			c++
		}
		for j := range lcp {
			if i == j {
				// 判断对角线
				if lcp[i][j] != n-i {
					return ""
				}
				continue
			}
			if j < i {
				// 左下角用于判断是否对称
				if lcp[j][i] != lcp[i][j] {
					return ""
				}
				// 右上角用于拼接字符串，所以左下角的返回
				continue
			}
			if lcp[i][j] > (n - i) {
				return ""
			}
			if lcp[i][j] > 0 {
				s[j] = s[i]
			}
		}
	}
	// 动态规划验证矩阵，倒序验证
	for i := n - 2; i >= 0; i-- {
		for j := n - 1; j > i; j-- {
			if s[i] != s[j] {
				if lcp[i][j] != 0 {
					return ""
				}
			} else if j == n-1 {
				if lcp[i][j] != 1 {
					return ""
				}
			} else {
				if lcp[i][j] != lcp[i+1][j+1]+1 {
					return ""
				}
			}
		}
	}
	return string(s)
}
