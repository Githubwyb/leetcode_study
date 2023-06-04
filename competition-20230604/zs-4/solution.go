package main

func count(num1 string, num2 string, min_sum int, max_sum int) int {
	var mod int = 1e9 + 7
	// 返回小于s的满足条件的数量
	getCount := func(s string) int {
		// 状态记忆数组，第一维是位数，第二维是状态（当前表示为前面数字的和），value是数量
		memo := make([][]int, len(s))
		for i := range memo {
			memo[i] = make([]int, max_sum+1)
			for j := range memo[i] {
				memo[i][j] = -1
			}
		}

		// p为当前要枚举的位，0是最高位，len(s)-1是最低位
		// sum是前面位数的和
		// limitUp代表前面的数位是否都到达上界
		var dfs func(p, sum int, limitUp bool) (res int)
		dfs = func(p, sum int, limitUp bool) (res int) {
			if sum > max_sum {
				return
			}
			if p == len(s) {
				// 到头了，sum必须大于等于最小sum
				if sum >= min_sum {
					return 1
				}
				return
			}

			if !limitUp {
				// 没到上界才能取状态，否则状态是假的
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

	ans := getCount(num2) - getCount(num1) + mod
	// 判断一下num1是否合法，上面直接减去了num1而不是num1-1，少算了num1
	sumNum1 := 0
	for _, c := range num1 {
		sumNum1 += int(c - '0')
	}
	if sumNum1 >= min_sum && sumNum1 <= max_sum {
		ans++
	}
	return ans % mod
}
