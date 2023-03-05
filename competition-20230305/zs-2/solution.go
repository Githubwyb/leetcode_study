package main

import (
	"container/heap"
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
	queue := list.New()
	bHeap := make(LittleHeap, 0, k)
	queue.PushBack(root)
	queue.PushBack(nil)
	var sum int64 = 0
	for queue.Len() > 0 {
		tmp := queue.Front()
		queue.Remove(tmp)
		if tmp.Value == nil {
			if len(bHeap) < k {
				heap.Push(&bHeap, sum)
			} else if sum > bHeap[0] {
				heap.Pop(&bHeap)
				heap.Push(&bHeap, sum)
			}
			sum = 0
			if queue.Len() > 0 {
				queue.PushBack(nil)
			}
			continue
		}
		node := tmp.Value.(*TreeNode)
		sum += int64(node.Val)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	if len(bHeap) < k {
		return -1
	}
	return bHeap[0]
}
