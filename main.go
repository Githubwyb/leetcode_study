package main

import (
	"container/list"
	"fmt"
)

type lineT struct {
	left  int
	right int
}

// 返回是否可以到达所有点
func bfs(connections [][]int) bool {
	seen := make([]bool, len(connections))
	seen[0] = true
	var queue list.List
	queue.PushBack(0)
	for queue.Len() > 0 {
		point := queue.Front().Value.(int)
		queue.Remove(queue.Front())
		// 向所有点走一遍，能走到就设置seen并加入队列，已经走过就不用了
		for i, v := range connections[point] {
			// 从这里走不到或者已经走过了
			if v == 0 || seen[i] {
				continue
			}
			seen[i] = true
			queue.PushBack(i)
		}
	}
	for _, v := range seen {
		if !v {
			return true
		}
	}
	return false
}

func getCount(n int, lines []lineT) int {
	result := 0
	connections := make([][]int, n)
	for index := range connections {
		connections[index] = make([]int, n)
		connections[index][index] = 1
	}
	for _, v := range lines {
		connections[v.left][v.right]++
		connections[v.right][v.left]++
	}
	for i := range lines {
		connections[lines[i].left][lines[i].right]--
		connections[lines[i].right][lines[i].left]--
		for j := i + 1; j < len(lines); j++ {
			connections[lines[j].left][lines[j].right]--
			connections[lines[j].right][lines[j].left]--
			if bfs(connections) {
				result++
			}
			connections[lines[j].left][lines[j].right]++
			connections[lines[j].right][lines[j].left]++
		}
		connections[lines[i].left][lines[i].right]++
		connections[lines[i].right][lines[i].left]++
	}
	return result
}

func main() {
	fmt.Println(getCount(3, []lineT{{0, 1}, {1, 2}, {2, 0}}))
	fmt.Println(getCount(2, []lineT{{0, 0}, {0, 0}, {0, 1}}))
	fmt.Println(getCount(3, []lineT{{0, 1}, {0, 1}, {0, 1}, {1, 2}, {1, 2}, {1, 2}}))
	fmt.Println(getCount(4, []lineT{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {2, 3}}))
}
