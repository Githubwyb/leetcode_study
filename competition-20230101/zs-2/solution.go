package main

func distinctPrimeFactors(nums []int) int {
	set := make(map[int]bool)
	for _, v := range nums {
		// 从2开始除，因为如果找到质数直接除就好，找到合数能整除，在质数时已经除了
		for i := 2; i <= v; i++ {
			if v%i == 0 {
				set[i] = true
				v, i = v/i, 1
			}
		}
	}
	return len(set)
}
