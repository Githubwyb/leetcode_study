package main

func gcd(a int, b int) int {
	// 当a为最大公约数时，计算后a = 0，b = a
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func maxScore(nums []int) int64 {
	// 先计算后缀，多一个数字好计算
	// 0和任何数字的最大公约数都是这个数本身
	// 1和任何数字的最大公倍数都是这个数本身
	n := len(nums)
	sufLcm := make([]int, n+1)
	sufGcd := make([]int, n+1)
	sufLcm[n] = 1
	for i := range nums {
		i = n - 1 - i
		v := nums[i]
		sufLcm[i] = lcm(sufLcm[i+1], v)
		sufGcd[i] = gcd(sufGcd[i+1], v)
	}

	// 先用一个都不排除做初值找最大
	ans := int64(sufLcm[0] * sufGcd[0])
	preGcd, preLcm := 0, 1
	// 一个一个排除，然后取前缀和后缀进行lcm/gcd找最大
	for i, v := range nums {
		tmp := int64(gcd(sufGcd[i+1], preGcd) * lcm(sufLcm[i+1], preLcm))
		if tmp > ans {
			ans = tmp
		}
		preGcd = gcd(preGcd, v)
		preLcm = lcm(preLcm, v)
	}
	return ans
}

func maxScore1(nums []int) int64 {
	ans := int64(0)

	// 爆破计算，对每一个元素都计算排除掉它的剩余算素的lcm和gcd
	// 取-1是满足不去除也要计算在内
	for i := -1; i < len(nums); i++ {
		lcmNum := 0
		gcdNum := 0
		for j, v := range nums {
			if j == i {
				continue
			}
			if lcmNum == 0 {
				lcmNum = v
				gcdNum = v
			} else {
				lcmNum = lcm(v, lcmNum)
				gcdNum = gcd(v, gcdNum)
			}
		}
		tmp := int64(lcmNum * gcdNum)
		if ans < tmp {
			ans = tmp
		}
	}
	return ans
}
