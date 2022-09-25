package main

import (
	_ "fmt"
)

// 检查第一个是否可以赢，true可以，false不可以
func checkIsFirstWin(n int, m int) bool {
	if m < 0 || n < 0 {
		return false
	}
	if m == 0 && n == 0 {
		return false
	}

	less, bigger := m, n
	if m > n {
		less = n
		bigger = m
	}

	// 0 n 谁先拿谁赢
	if less == 0 {
		return true
	}

	// 1 1 谁先拿谁输，大于1，构造1就赢了
	if less == 1 {
		return bigger > 1
	}

	// 相等就输
	// 2 2，构造 1 2 和 0 2 都输了
	// 3 3，构造 2 3 对方构造 2 2 就赢了
	//      构造 1 3 对方构造 1 1 就赢了
	//      构造 0 3 对方直接赢了
	// 那么不等的情况，第一个拿的只需要构造相等就赢了
	return less != bigger
}
