package main

func minArraySum(nums []int, k int, op1 int, op2 int) int {
	n := len(nums)

	// 设置记忆化的map
	cacheMap := make([][][]int, n)
	for i := range cacheMap {
		tmp := make([][]int, op1+1)
		for j := range tmp {
			tmp1 := make([]int, op2+1)
			for k := range tmp1 {
				tmp1[k] = -1
			}
			tmp[j] = tmp1
		}
		cacheMap[i] = tmp
	}

	// 返回从i开始的sum
	var dfs func(i, op1, op2 int) int
	dfs = func(i, op1, op2 int) int {
		if i == n {
			return 0
		}
		ans := cacheMap[i][op1][op2]
		if ans != -1 {
			return ans
		}

		v := nums[i]
		// 找执行op1，op2，op1+op2的
		ans = v + dfs(i+1, op1, op2)
		if op1 > 0 {
			ans = min(ans, (v+1)/2+dfs(i+1, op1-1, op2))
		}
		if op2 > 0 && k > 0 && v >= k {
			ans = min(ans, v-k+dfs(i+1, op1, op2-1))
			if op1 > 0 {
				// 如果除完大于k则先除后减，否则先减后除
				tmp := dfs(i+1, op1-1, op2-1)
				if (v+1)/2 >= k {
					tmp += (v+1)/2 - k
				} else {
					tmp += (v - k + 1) / 2
				}
				ans = min(ans, tmp)
			}
		}
		cacheMap[i][op1][op2] = ans
		return ans
	}
	return dfs(0, op1, op2)
}
