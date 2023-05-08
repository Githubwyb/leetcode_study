package main

import "sort"

func oulerPrimes(mx int, primes *[]int) {
	flag := make([]bool, mx+1) // 标记数有没有被筛掉，false就是没有
	for i := 2; i < mx+1; i++ {
		if !flag[i] {
			// 数没有被比自己小的数筛掉，就代表是质数
			*primes = append(*primes, i)
		}
		for _, v := range *primes {
			if i*v > mx {
				break
			}
			// 每一个数都作为因子乘以比自己小的素数筛掉后面的数
			flag[i*v] = true
			if i%v == 0 {
				// 减少时间复杂度的关键算法
				// 12 = 2 * 3 * 2，i = 4时，只排了8就退出了，因为6会将12排除
				// 也就是，假设v可以整除i即i = kv，有某个数为x = mi = kmv
				//        那么存在一个数 i < km < x可以把x排掉，用i乘以所有的质数去排除就没什么意义了，提前退出减少时间复杂度
				break
			}
		}
	}
}

var primes []int = make([]int, 0, 168)

func init() {
	oulerPrimes(1000, &primes)
}

func primeSubOperation(nums []int) bool {
	n := len(nums)
	last := nums[n-1]
	for i := n - 2; i >= 0; i-- {
		v := nums[i]
		if v < last {
			last = v
			continue
		}
		for _, vp := range primes {
			if vp >= v {
				return false
			}
			if v-vp < last {
				v -= vp
				break
			}
		}
		if v >= last {
			return false
		}
		last = v
		continue
	}
	return true
}

func primeSubOperation1(nums []int) bool {
	pre := 0
	for _, v := range nums {
		// 每个数都尽可能小
		if v <= pre {
			return false
		}
		// 找到能尽可能接近pre的最大质数，反过来就是找v-pre之间的最大质数
		i := sort.SearchInts(primes, v-pre) - 1
		if i < 0 {
			pre = v
		} else {
			pre = v - primes[i]
		}
	}
	return true
}
