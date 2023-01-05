package main

import (
	"math"
	"sort"
)

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	target := int(math.Sqrt(float64(num))) + 1
	for i := 2; i < target; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func closestPrimes(left int, right int) []int {
	state := 0
	p, n := 0, 0
	result := []int{-1, -1}
	min := math.MaxInt
	for i := left; i <= right; i++ {
		if !isPrime(i) {
			continue
		}
		switch state {
		case 0:
			// 未找到第一个质数
			state = 1
			p = i
		case 1:
			// 找到第一个质数，未找到第二个
			state = 2
			n = i
			result[0], result[1] = p, n
			min = n - p
		case 2:
			// 两个都找到了，找个最小的
			p, n = n, i
			if min > n-p {
				result[0], result[1] = p, n
				min = n - p
			}
		}
		if min <= 2 {
			return result
		}
	}
	return result
}

var primes []int = make([]int, 0, 78500)

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

func init() {
	oulerPrimes(1e6, &primes)
	primes = append(primes, 1e6+1) // 加一个边界防止越界
}

func closestPrimes1(left int, right int) []int {
	i := sort.SearchInts(primes, left)
	state := 0
	p, n := 0, 0
	result := []int{-1, -1}
	min := math.MaxInt
	for ; primes[i] <= right; i++ {
		switch state {
		case 0:
			// 未找到第一个质数
			state = 1
			p = primes[i]
		case 1:
			// 找到第一个质数，未找到第二个
			state = 2
			n = primes[i]
			result[0], result[1] = p, n
			min = n - p
		case 2:
			// 两个都找到了，找个最小的
			p, n = n, primes[i]
			if min > n-p {
				result[0], result[1] = p, n
				min = n - p
			}
		}
	}
	return result
}
