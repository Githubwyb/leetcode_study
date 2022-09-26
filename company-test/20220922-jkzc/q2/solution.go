package main

import (
	_ "fmt"
)

/*
小明有N块草莓蛋糕，小红有M块巧克力蛋糕(N+M<1000)，他们都觊觎着对方手里的蛋糕，最终他们决定来玩一个游戏。他们分别把蛋糕暂时地放到了桌子上，两个人轮流从桌子上拿去蛋糕，每次允许从草莓蛋糕或者巧克力蛋糕里拿取「最少」一块「最多」K块蛋糕(不允许即拿草莓又拿巧克力蛋糕)，拿到最后一块蛋糕的人将赢得所有的蛋糕。假设两个小朋友都足够聪明，先拿的人是赢还是输呢?
*/

// 检查第一个是否可以赢，true可以，false不可以
func checkIsFirstWin(n int, m int, k int) bool {
	// 取m和n中较大的那个（假设是m），如果满足m=a(k+1)+n，a为0到正无穷的整数。先拿的输了，否则先拿的就赢了，
	if m < 0 || n < 0 {
		return false
	}
	if m == 0 && n == 0 {
		return false
	}

	diff := m - n
	if diff < 0 {
		diff = -diff
	}

	return diff%(k+1) != 0
}
