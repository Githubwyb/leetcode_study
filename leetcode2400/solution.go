package main

const maxSum int = 1e9 + 7

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getCount(sum int, pos int) int {
	dp := make([][]int, sum+1)
	for i := range dp {
		dp[i] = make([]int, pos+1)
		for j := range dp[i] {
			if i == j {
				dp[i][j] = 1
			} else if i == 1 {
				dp[i][j] = j
			} else {
				dp[i][j] = -1
			}
		}
	}

	for i := 1; i <= pos; i++ {
		for j := 1; j <= i && j <= sum; j++ {
			if j == 1 || i == j {
				continue
			}
			if dp[j][i] != -1 {
				continue
			}

			// 1 + dp[sum-1][sum] + dp[sum-1][sum+1]
			tmp := 1
			for k := 0; k < i-j; k++ {
				tmp += dp[j-1][j+k]
				tmp %= maxSum
			}
			dp[j][i] = tmp
		}
	}
	return dp[sum][pos]
}

func numberOfWays(startPos int, endPos int, k int) int {
	diff := abs(startPos - endPos)
	// 首先判断是否可达
	if k < diff {
		return 0
	}
	if k == diff {
		return 1
	}
	// 判断是否奇偶不同
	if (k % 2) != (diff % 2) {
		return 0
	}

	return getCount((k-diff)/2, k)
}

func numberOfWays1(startPos int, endPos int, k int) int {
	diff := abs(startPos - endPos)
	// 首先判断是否可达
	if k < diff {
		return 0
	}
	if k == diff {
		return 1
	}
	// 判断是否奇偶不同
	if (k % 2) != (diff % 2) {
		return 0
	}

	n := (k - diff) / 2

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}

	for i := 1; i < k+1; i++ {
		dp[1][i] = i
		// j最大不超过i
		for j := 2; j < n+1 && j <= i; j++ {
			dp[j][i] = (dp[j][i-1] + dp[j-1][i-1]) % maxSum
		}
	}

	return dp[n][k]
}
