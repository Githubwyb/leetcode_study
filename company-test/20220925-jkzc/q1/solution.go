package main

import (
	_ "fmt"
)

func getInsertLocation(arr []int, in int) int {
	// 从左开始返回第一个满足cmp的下标i，要求i右边的都要满足
	cmp := func(i int) bool {
		// 最后一个可插入的，就是第一个大于的
		return arr[i] > in
	}

	left, right := 0, len(arr)
	for left < right {
		// 奇数个数返回中间，偶数个数返回中间偏右的一个
		mid := (left + right) / 2
		// 此点满足，说明要么是这个，要么是左边的
		if cmp(mid) {
			right = mid
			continue
		}
		// 此点不满足，那么一定在右边
		left = mid + 1
	}
	// right一定满足，出来时left只能等于right
	return right
}
