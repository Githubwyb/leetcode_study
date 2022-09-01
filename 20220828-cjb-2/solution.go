package main

import (
	"math"
)

var kPrice = []int{1, 3, 7, 11, 13}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type state struct {
	count int
	coins []int
}

func getMinCount(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	amountMap := make([]state, amount+1)
	amountMap[0].count = 0
	amountMap[0].coins = make([]int, len(coins))
	copy(amountMap[0].coins, coins)
	for i := 1; i <= amount; i++ {
		amountMap[i].count = math.MaxInt
		for j := range kPrice {
			v := kPrice[j]
			if i < kPrice[j] {
				continue
			}
			// 当硬币数不够的时候，跳过
			if amountMap[i-v].count == math.MaxInt || amountMap[i-v].coins[j] == 0 {
				continue
			}
			amountMap[i].count = min(amountMap[i].count-1, amountMap[i-v].count) + 1
			amountMap[i].coins = make([]int, len(coins))
			copy(amountMap[i].coins, amountMap[i-v].coins)
			amountMap[i].coins[j]--
		}
	}
	if amountMap[amount].count == math.MaxInt {
		return -1
	}
	return amountMap[amount].count
}
