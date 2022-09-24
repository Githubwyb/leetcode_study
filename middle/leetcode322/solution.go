package main

import (
	"math"
	"sort"
)

func getMinCount(coins []int, amount int, amountMap []int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	// 存在记录，直接返回
	if amountMap[amount] != 0 {
		return amountMap[amount]
	}

	// 根据coins遍历，取所有方案最小的
	minCount := -1
	for i := len(coins); i > 0; i-- {
		if amount < coins[i-1] {
			continue
		}
		if amount == coins[i-1] {
			amountMap[amount] = 1
			return 1
		}
		tmp := getMinCount(coins, amount-coins[i-1], amountMap)
		// 没结果的，找下一个
		if tmp < 0 {
			continue
		}
		if minCount < 0 {
			minCount = tmp
			continue
		}
		if tmp < minCount {
			minCount = tmp
		}
	}
	if minCount >= 0 {
		amountMap[amount] = minCount + 1
		return minCount + 1
	}
	amountMap[amount] = -1
	return -1
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	amountMap := make([]int, amount+1)
	sort.Ints(coins)
	return getMinCount(coins, amount, amountMap)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func coinChange1(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	amountMap := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		amountMap[i] = math.MaxInt
		for _, v := range coins {
			if i < v {
				continue
			}
			amountMap[i] = min(amountMap[i]-1, amountMap[i-v]) + 1
		}
	}
	if amountMap[amount] == math.MaxInt {
		return -1
	}
	return amountMap[amount]
}
