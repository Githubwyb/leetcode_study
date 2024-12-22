package main

import "sort"

func minLength(s string, numOps int) int {
	n := len(s)
	// 二分法查找结果
	ans := sort.Search(n, func(mx int) bool {
		if mx == 0 {
			return false
		}
		cnt := 0
		if mx == 1 {
			// 最大为1的情况下，取1开头和0开头最小值
			for i, v := range s {
				if int(v&1) != i&1 {
					cnt++
				}
			}
			cnt = min(cnt, n-cnt)
		} else {
			for i := 0; i < n; {
				// 查找连续字串，分割为mx+1的次数
				st := i
				for i++; i < n && s[i] == s[i-1]; i++ {
				}
				l := i - st
				cnt += l / (mx + 1)
			}
		}
		return cnt <= numOps
	})
	return ans
}
