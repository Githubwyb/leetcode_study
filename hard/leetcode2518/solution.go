package main

func countPartitions(nums []int, k int) int {
	const mod int = 1e9 + 7
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum < k*2 {
		return 0
	}

	ans := 1
	f := make([]int, k)
	f[0] = 1
	// 从0算到n，不需要记录已经算过的值，即计算i时，使用i-1的结果，不关心i-2的
	for _, v := range nums {
		ans = ans * 2 % mod
		// 要加上j和j-nums[i]，那么j肯定要大于等于nums[i]
		// j只需要计算到k-1即可
		for j := k - 1; j >= v; j-- {
			// f[i][j] = f[i-1][j] + f[i-1][j-nums[i]]
			// 对于当前i来说，f存储的是i-1的值，那么就是f[j] = f[j] + f[j-nums[i]]
			f[j] = (f[j] + f[j-v]) % mod
		}
	}
	for _, v := range f {
		ans -= v * 2
	}
	return (ans%mod + mod) % mod
}
