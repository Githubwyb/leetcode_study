package main

func mgcd(a int, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func subarrayGCD(nums []int, k int) int {
	result := 0
	for i := range nums {
		// 从起点向后最大公约数设为tmp
		// 向后遍历查看有几个子数组
		for tmp, j := nums[i], i; j < len(nums); j++ {
			if nums[j]%k != 0 {
				// 数组中存在一个不是k的倍数的值，就不用继续找了
				break
			}

			// 倍数的最大公约数向后传递，如 gcd(8, 4) = 4 => gcd(8, 4, 6) = gcd(4, 6) = 2
			if tmp != k {
				tmp = mgcd(tmp, nums[j])
			}

			if tmp == k {
				result++
			}
		}
	}
	return result
}
