package main

import (
	"container/list"
	"fmt"
	"sort"
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
	// fmt.Println(getCount(3, []lineT{{0, 1}, {1, 2}, {2, 0}}))
	// fmt.Println(getCount(2, []lineT{{0, 0}, {0, 0}, {0, 1}}))
	// fmt.Println(getCount(3, []lineT{{0, 1}, {0, 1}, {0, 1}, {1, 2}, {1, 2}, {1, 2}}))
	// fmt.Println(getCount(4, []lineT{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {2, 3}}))
	a := []int{1, 2, 2, 4, 4, 5}

	// 二分查找从最左边开始的第一个满足条件的索引
	fmt.Println(sort.Search(len(a), func(i int) bool { return a[i] >= 2 })) // 1
	// 二分查找从最左边开始第一个 func(i) <= 0 的索引，并返回对应索引是否 func(i) == 0
	fmt.Println(sort.Find(len(a), func(i int) int { return 3 - a[i] })) // 3, false
	// 基本类型可以直接调用，但是只支持升序的slice
	// 源码就是调用 sort.Search(len(a), func(i int) bool { return a[i] >= 2 }
	fmt.Println(sort.SearchInts(a, 2)) // 1
}
