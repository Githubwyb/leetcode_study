package main

import (
	"container/list"
	_ "fmt"
)

/********** 第一种解法，暴力+bfs **********/

type pointT struct {
	x int
	y int
}

var (
	r int // 行数
	c int // 列数
	// 按照上下左右的相对位置设定，用于后面方便找四周的点
	kRoundPoints = []pointT{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
	}
)

func shortestPathAllKeys(grid []string) int {
	r = len(grid)
	c = len(grid[0])

	// location['a'] is the (x, y) of a
	location := make(map[byte]pointT)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			ch := grid[i][j]
			if ch != '.' && ch != '#' {
				location[ch] = pointT{j, i}
			}
		}
	}
	keyNums := len(location) / 2
	// 枚举到所有的钥匙，题目条件只会从a-b、a-c、a-d、a-e、a-f几种情况
	alphabet := make([]byte, keyNums)
	for i := 0; i < keyNums; i++ {
		alphabet[i] = byte('a' + i)
	}

	res := -1
	permutation(alphabet, 0, func(str []byte) {
		ans := 0
		keymask := 0
		for i := 0; i < len(str); i++ {
			var src byte
			if i == 0 {
				src = '@'
			} else {
				src = alphabet[i-1]
			}
			tmp := bfs(location[src], location[alphabet[i]], grid, keymask)
			if tmp == -1 {
				return
			}
			ans += tmp
			keymask |= 1 << (alphabet[i] - 'a')
		}
		if res == -1 || ans < res {
			res = ans
		}
	})
	return res
}

// 全排列
func permutation(str []byte, index int, f func(str []byte)) {
	if len(str) == index {
		f(str)
		return
	}

	// 不交换的场景
	permutation(str, index+1, f)
	// index对应位置向后交换
	for i := index + 1; i < len(str); i++ {
		str[i], str[index] = str[index], str[i]
		permutation(str, index+1, f)
		str[i], str[index] = str[index], str[i]
	}
}

// 计算src到dst的最短路径长度，keymask为按位标记某个钥匙是否已找到
// 返回从src到dst的最短路径长度
func bfs(src pointT, dst pointT, grid []string, keymask int) int {
	// 减小计算量，走过的路不再走，记录一下哪里走过了
	seen := make([][]bool, len(grid))
	for i := range seen {
		seen[i] = make([]bool, len(grid[0]))
	}
	// 源地址记录走过了，注意x是第二维的坐标
	seen[src.y][src.x] = true

	// 使用层数作为步数
	curDepth := 0
	var queue list.List
	// 插入源地址，作为第一层，使用nil作为层间隔
	queue.PushBack(src)
	queue.PushBack(nil)
	// 队列一定含有一个层间隔，不在头就在尾，如果只剩一个层间隔，说明没路可走
	for queue.Len() > 1 {
		tmp := queue.Front().Value
		queue.Remove(queue.Front())
		if tmp == nil {
			// 找到层间隔，说明当前层遍历完了，步数加一准备下一层
			curDepth++
			// 当前层遍历完，队列剩余的都是下一层，加入一个层间隔
			queue.PushBack(nil)
			continue
		}

		// 判断当前点是不是目标点，如果是，说明走到了，返回步数
		tx, ty := tmp.(pointT).x, tmp.(pointT).y
		if tx == dst.x && ty == dst.y {
			return curDepth
		}
		// 不是目标点，从此点出发，向四周走一下
		for i := range kRoundPoints {
			px, py := tx+kRoundPoints[i].x, ty+kRoundPoints[i].y
			// 如果超出边界或者已经走过了或者碰到墙，就继续
			if py < 0 || py >= len(grid) || px < 0 || px >= len(grid[0]) || seen[py][px] || grid[py][px] == '#' {
				continue
			}
			ch := grid[py][px]
			// 如果是锁，看一下有没有钥匙，没有钥匙就跳过
			if (ch <= 'Z' && ch >= 'A') && ((1<<(ch-'A'))&keymask) == 0 {
				continue
			}
			// 这个点可以走，走上去，记录到队列中，作为下一层的起点
			seen[py][px] = true
			queue.PushBack(pointT{px, py})
		}
	}
	return -1
}

/********** 第二种解法，状态压缩+bfs **********/

type pointST struct {
	x     int
	y     int
	state int
}

func shortestPathAllKeys1(grid []string) int {
	// location['a'] is the (x, y) of a
	location := make(map[byte]pointT)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			ch := grid[i][j]
			if ch != '.' && ch != '#' {
				location[ch] = pointT{j, i}
			}
		}
	}
	keyNums := len(location) / 2
	// 枚举到所有的钥匙，题目条件只会从a-b、a-c、a-d、a-e、a-f几种情况
	alphabet := make([]byte, keyNums)
	for i := 0; i < keyNums; i++ {
		alphabet[i] = byte('a' + i)
	}
	finalState := (1 << keyNums) - 1

	// 将钥匙的持有状态作为判断成三维的bfs
	return bfsThree(location['@'], grid, finalState)
}

// 将钥匙的持有状态作为判断成三维的bfs
func bfsThree(src pointT, grid []string, finalState int) int {
	sx, sy := src.x, src.y
	srcP := pointST{sx, sy, 0}

	// 减小计算量，走过的路不再走，记录一下哪里走过了
	seen := make([][]map[int]bool, len(grid))
	for i := range seen {
		seen[i] = make([]map[int]bool, len(grid[0]))
		for j := range seen[i] {
			seen[i][j] = make(map[int]bool)
		}
	}
	seen[sy][sx][0] = true

	// 使用层数作为步数
	curDepth := 0
	var queue list.List
	// 插入源地址，作为第一层，使用nil作为层间隔
	queue.PushBack(srcP)
	queue.PushBack(nil)
	// 队列一定含有一个层间隔，不在头就在尾，如果只剩一个层间隔，说明没路可走
	for queue.Len() > 1 {
		tmp := queue.Front().Value
		queue.Remove(queue.Front())
		if tmp == nil {
			// 找到层间隔，说明当前层遍历完了，步数加一准备下一层
			curDepth++
			// 当前层遍历完，队列剩余的都是下一层，加入一个层间隔
			queue.PushBack(nil)
			continue
		}

		// 判断当前点是不是已经达成目标
		tx, ty, state := tmp.(pointST).x, tmp.(pointST).y, tmp.(pointST).state
		if state == finalState {
			return curDepth
		}
		// 不是目标点，从此点出发，向四周走一下
		for i := range kRoundPoints {
			px, py := tx+kRoundPoints[i].x, ty+kRoundPoints[i].y
			// 如果超出边界或者碰到墙，就继续
			if py < 0 || py >= len(grid) || px < 0 || px >= len(grid[0]) || grid[py][px] == '#' {
				continue
			}
			// 判断是否曾以相同状态要走过这个点
			tmpMap := seen[py][px]
			if _, ok := tmpMap[state]; ok {
				continue
			}
			ch := grid[py][px]
			// 如果是锁，看一下有没有钥匙，没有钥匙就跳过
			if (ch <= 'Z' && ch >= 'A') && ((1<<(ch-'A'))&state) == 0 {
				continue
			}
			// 可以踩上去，就记录踩上去前的状态
			seen[py][px][state] = true
			// 如果是钥匙，保存新的状态到队列中
			tmpState := state
			if ch <= 'z' && ch >= 'a' {
				tmpState |= (1 << (ch - 'a'))
			}
			// 记录到队列中，作为下一层的点
			queue.PushBack(pointST{px, py, tmpState})
		}
	}
	return -1
}
