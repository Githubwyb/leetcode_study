package common

// 数位dp模板
func NumDP() {
	var mod int = 1e9 + 7
	max_sum := 10

	stateNum := max_sum
	// 返回小于s的满足条件的数量
	getCount := func(s string) int {
		// 状态记忆数组，第一维是位数，第二维是状态（状态根据具体情况来），value是从这一位向后满足状态的数量
		memo := make([][]int, len(s))
		for i := range memo {
			memo[i] = make([]int, stateNum+1)
			for j := range memo[i] {
				memo[i][j] = -1
			}
		}

		// p为当前要枚举的位，0是最高位，len(s)-1是最低位
		// sum是前面位数的和
		// limitUp代表前面的数位是否都到达上界
		var dfs func(p, sum int, limitUp bool) (res int)
		dfs = func(p, sum int, limitUp bool) (res int) {
			// 处理一些限制条件
			// TODO
			if p == len(s) {
				// 到最后一位了，满足条件返回1，不满足返回0
				// TODO
				if true {
					return 1
				}
				return
			}

			if !limitUp {
				// 没到上界才能取状态下的值，否则状态是假的
				tmp := memo[p][sum]
				if tmp >= 0 {
					return tmp
				}
				defer func() { memo[p][sum] = res }()
			}
			up := 9
			if limitUp {
				up = int(s[p] - '0')
			}
			for d := 0; d <= up; d++ {
				res = (res + dfs(p+1, sum+d, limitUp && d == int(s[p]-'0'))) % mod
			}
			return
		}
		return dfs(0, 0, true)
	}
	getCount("1234")
}
