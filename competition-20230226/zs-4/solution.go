package main

import (
	"container/heap"
	"math"
	"sort"
)

type locationT struct {
	x, y     int
	distance int
}
type LittleHeap []locationT

func (h *LittleHeap) Len() int { return len(*h) }

// less必须满足当Less(i, j)和Less(j, i)都为false，则两个索引对应的元素相等
// 为true，i向栈顶移动；为false，j向栈顶移动
func (h *LittleHeap) Less(i, j int) bool { return (*h)[i].distance < (*h)[j].distance }
func (h *LittleHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *LittleHeap) Push(x interface{}) {
	*h = append(*h, x.(locationT))
}

func (h *LittleHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func minimumTime(grid [][]int) int {
	if grid[1][0] > 1 && grid[0][1] > 1 {
		return -1
	}
	var arrLoc [][]int = [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	m, n := len(grid), len(grid[0])
	var lHeap LittleHeap = make([]locationT, 0, m*n)
	lHeap = append(lHeap, locationT{
		x: 0, y: 0, distance: 0,
	})

	disMap := make([][]int, m)
	for i := range disMap {
		disMap[i] = make([]int, n)
		for j := range disMap[i] {
			disMap[i][j] = math.MaxInt
		}
	}
	// 一定能到，所以不用加判断条件
	// 不能到的上面干掉了
	for {
		tmp := heap.Pop(&lHeap).(locationT)
		if tmp.x == m-1 && tmp.y == n-1 {
			return tmp.distance
		}
		for _, v := range arrLoc {
			x, y := tmp.x+v[0], tmp.y+v[1]
			if x < 0 || x >= m || y < 0 || y >= n {
				continue
			}
			dis := max(tmp.distance+1, grid[x][y])
			dis += (dis - x - y) % 2
			if dis < disMap[x][y] {
				disMap[x][y] = dis
				heap.Push(&lHeap, locationT{
					x: x, y: y, distance: dis,
				})

			}
		}
	}
}

// 二分查找+BFS
func minimumTime1(grid [][]int) int {
	if grid[1][0] > 1 && grid[0][1] > 1 {
		return -1
	}
	var arrLoc [][]int = [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	m, n := len(grid), len(grid[0])
	// 用是否等于当前的endTime来标记是否走过
	seen := make([][]int, m)
	for i := range seen {
		seen[i] = make([]int, n)
	}
	type pointT struct {
		x, y int
	}
	q1 := make([]pointT, 0, m+n)
	q2 := make([]pointT, 0, m+n)
	endTime := sort.Search(1e5+m+n, func(endTime int) bool {
		if endTime < grid[m-1][n-1] {
			return false
		}
		q1 = q1[:0]
		q1 = append(q1, pointT{
			x: m - 1, y: n - 1,
		})
		seen[m-1][n-1] = endTime
		for t := endTime - 1; len(q1) > 0; t-- {
			q2 = q2[:0]
			for _, point := range q1 {
				for _, v := range arrLoc {
					x, y := point.x+v[0], point.y+v[1]
					if x < 0 || x >= m || y < 0 || y >= n || seen[x][y] == endTime || grid[x][y] > t {
						continue
					}
					if x == 0 && y == 0 {
						return true
					}
					seen[x][y] = endTime
					q2 = append(q2, pointT{
						x: x, y: y,
					})
				}
			}
			q1, q2 = q2, q1
		}
		return false
	})
	return endTime + (endTime+m+n)%2
}
