package main

import "math"

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

func closestPrimes1(left int, right int) []int {
	min := math.MaxInt
	result := []int{-1, -1}
	flag := make([]bool, right+1)
	primeNums := []int{}
	p, n := 0, 0
	state := 0
	for i := 2; i <= right; i++ {
		if !flag[i] {
			primeNums = append(primeNums, i)
			if i >= left {
				switch state {
				case 0:
					p = i
					state = 1
				case 1:
					n = i
					result[0], result[1] = p, n
					min = n - p
					state = 2
				case 2:
					p, n = n, i
					if n-p < min {
						result[0], result[1] = p, n
						min = n - p
					}
				}
				if min <= 2 {
					return result
				}
			}
		}

		for _, p := range primeNums {
			if i*p > right {
				break
			}
			flag[i*p] = true
			if i%p == 0 {
				break
			}
		}
	}
	return result
}
