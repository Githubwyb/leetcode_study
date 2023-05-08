package main

import (
	"container/heap"
	. "leetcode/common"
)

type LittleHeap []int64

func (h *LittleHeap) Len() int { return len(*h) }

// less必须满足当Less(i, j)和Less(j, i)都为false，则两个索引对应的元素相等
// 为true，i向栈顶移动；为false，j向栈顶移动
func (h *LittleHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *LittleHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *LittleHeap) Push(x interface{}) {
	*h = append(*h, x.(int64))
}

func (h *LittleHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargestLevelSum(root *TreeNode, k int) int64 {
	q := make([]*TreeNode, 0, 1)
	tmp := make([]*TreeNode, 0, 1)
	bHeap := make(LittleHeap, 0, k)

	q = append(q, root)
	for len(q) > 0 {
		tmp, q = q, tmp
		var s int64 = 0
		q = q[:0]
		for _, node := range tmp {
			s += int64(node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		if len(bHeap) < k {
			heap.Push(&bHeap, s)
		} else if s > bHeap[0] {
			heap.Pop(&bHeap)
			heap.Push(&bHeap, s)
		}
	}
	if len(bHeap) < k {
		return -1
	}
	return bHeap[0]
}
