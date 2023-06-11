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

func count1(num1 string, num2 string, min_sum int, max_sum int) int {
	var mod int = 1e9 + 7

	// 返回能取到的最大值和最小值
	getMinMax := func(s int) (minNum, maxNum int) {
		minNum = 0
		maxNum = 9
		if s+minNum < min_sum {
			minNum = min_sum - s
		}
		if s+maxNum > max_sum {
			maxNum = max_sum - s
		}
		return
	}

	// 返回小于s的满足条件的数量
	getCount := func(s string) int {
		n := len(s)
		// 递推，从最后一位开始算，一个数组保存状态（从最高位到这里的和），只保存没有到达上限的，即可以从0到n个9的
		limitState := 0 // 保存每一位当前面所有位到达上限时的后续数量
		limitSum := 0
		for _, v := range s {
			limitSum += int(v - '0')
		}
		limitSum -= int(s[n-1] - '0')
		minNum, maxNum := getMinMax(limitSum)
		for d := minNum; d <= int(s[n-1]-'0') && d <= maxNum; d++ {
			limitState++
		}
		// 只有一位就可以直接返回了
		if n <= 1 {
			return limitState
		}

		state := make([]int, max_sum+1)
		// 最后一位的状态的初始值
		for s := range state {
			// 前面和加上9还小于min_sum的都是0
			if s+9 < min_sum {
				state[s] = 0
				continue
			}

			minNum, maxNum := getMinMax(s)
			state[s] = maxNum - minNum + 1
		}

		// 对可以取到0到n个9的状态递推到第2位
		for i := n - 2; i > 0; i-- {
			// 先计算此位到达上限的状态
			limitSum -= int(s[i] - '0')
			_, maxNum := getMinMax(limitSum)
			for d := 0; d < int(s[i]-'0') && d <= maxNum; d++ {
				limitState = (limitState + state[limitSum+d]) % mod
			}

			// 计算此位没有到达上限的状态
			curState := make([]int, max_sum+1)
			for s := range curState {
				_, maxNum := getMinMax(s)
				// 当前位的在此状态下的数量为低一位的state[s+d]的和
				for d := 0; d <= maxNum; d++ {
					curState[s] = (curState[s] + state[s+d]) % mod
				}
			}
			state = curState
		}

		// 第一位计算不超过上限的数量
		for d := 0; d < int(s[0]-'0') && d <= max_sum; d++ {
			limitState = (limitState + state[d]) % mod
		}
		return limitState
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

func count2(num1 string, num2 string, min_sum int, max_sum int) int {
	var mod int = 1e9 + 7

	// 返回能取到的最大值
	getMax := func(s int) int {
		maxNum := 9
		if s+maxNum > max_sum {
			maxNum = max_sum - s
		}
		return maxNum
	}
	getMin := func(s int) int {
		minNum := min_sum - s
		if minNum < 0 {
			minNum = 0
		}
		return minNum
	}

	// 返回小于s的满足条件的数量
	getCount := func(s string) int {
		n := len(s)
		// 最后一位要特殊处理，所以这里要对只有一位的情况单独计算
		if n == 1 {
			num := int(s[0] - '0')
			if num < min_sum {
				return 0
			}
			return num - min_sum + 1
		}
		// 记录前一位可以到达某个sum的统计数量，要求下一位可以从0算到9，也就是当前不能到上界
		preState := make([]int, max_sum+1)
		limitSum := 0
		// 遍历到除最后一位，最后一位要判断最小值
		for _, v := range s[:n-1] {
			num := int(v - '0')
			curState := make([]int, max_sum+1)
			// 到前一位已经得到的sum为st
			for st, c := range preState {
				if c == 0 {
					continue
				}
				maxNum := getMax(st)
				for d := 0; d <= maxNum; d++ {
					curState[d+st] = (curState[d+st] + c) % mod
				}
			}
			// 算一下上界的情况，当前位除了到达上界，下一位都可以按照0到9算，所以直接加到状态中
			maxNum := getMax(limitSum)
			for d := 0; d <= maxNum && d < num; d++ {
				curState[d+limitSum] = (curState[d+limitSum] + 1) % mod
			}
			limitSum += num
			preState = curState
		}

		// 遍历最后一位
		res := 0
		for st, c := range preState {
			if c == 0 {
				continue
			}
			minNum := getMin(st)
			maxNum := getMax(st)
			for d := minNum; d <= maxNum; d++ {
				res = (res + c) % mod
			}
		}
		// 算一下上界
		maxNum := getMax(limitSum)
		minNum := getMin(limitSum)
		for d := minNum; d <= maxNum && d <= int(s[n-1]-'0'); d++ {
			res = (res + 1) % mod
		}
		return res
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
