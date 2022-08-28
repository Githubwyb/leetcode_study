package main

import (
	"fmt"
)

var kPrice = []int{1, 3, 7, 11, 13}

func getMinProduct(money, index int, prices []int, minCount int) int {
	if index < 0 {
		return -1
	}
	if kPrice[index] > money {
		return getMinProduct(money, index-1, prices, minCount)
	}
	// 如果存在一个更优结果，此结果给这里剩余的个数不足以完成money，不继续尝试
	if minCount > 0 && minCount*kPrice[index] < money {
		return -1
	}
	// 最大化使用价钱
	maxCount := money / kPrice[index]
	if maxCount > prices[index] {
		maxCount = prices[index]
	}
	// 把钱花出去，正好花完就返回当前商品用了多少
	leftMoney := money - kPrice[index]*maxCount
	// 不剩钱直接返回
	if leftMoney == 0 {
		return maxCount
	}
	// 剩的有钱，看是否存在更小方案，默认无方案
	for i := maxCount; i >= 0; i-- {
		leftMoney = money - kPrice[index]*i
		count := getMinProduct(leftMoney, index-1, prices, minCount-i)
		// 下一个价格没有方案，少买一个当前的再试一次
		if count == -1 {
			continue
		}
		checkResult := i + count
		if minCount < 0 {
			// 有一个方案，先记录
			minCount = checkResult
		} else if checkResult < minCount {
			minCount = checkResult
		}
	}
	return minCount
}

func main() {
	// var a1, a3, a7, a11, a13 int
	// money := 0
	// for {
	// n, _ := fmt.Scan(&a1, &a3, &a7, &a11, &a13)
	// if n == 0 {
	// 	break
	// }
	// n, _ = fmt.Scan(&money)
	// if n == 0 {
	// 	break
	// }

	fmt.Println(getMinProduct(30, 4, []int{1, 2, 3, 4, 5}, -1))
	// }
}
