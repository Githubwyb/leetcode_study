package main

import "container/heap"

type worker struct {
	cost  int
	index int
}

type littleHeap []worker

func (h *littleHeap) Len() int { return len(*h) }

// 为true，i向栈顶移动；为false，j向栈顶移动
func (h *littleHeap) Less(i, j int) bool {
	if (*h)[i].cost < (*h)[j].cost {
		return true
	} else if (*h)[i].cost == (*h)[j].cost && (*h)[i].index < (*h)[j].index {
		return true
	}
	return false
}

func (h *littleHeap) Pop() interface{} {
	item := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return item
}
func (h *littleHeap) Push(v interface{}) {
	*h = append(*h, v.(worker))
}
func (h *littleHeap) Swap(i, j int) {
	(*h)[i].cost, (*h)[j].cost = (*h)[j].cost, (*h)[i].cost
	(*h)[i].index, (*h)[j].index = (*h)[j].index, (*h)[i].index
}

func totalCost(costs []int, k int, candidates int) int64 {
	// 取左右指针为candidates之外的索引
	l, r := candidates, len(costs)-candidates-1
	pickHeap := make(littleHeap, 0)
	for i := 0; i < len(costs); i++ {
		if i == l && i <= r {
			i = r + 1
		}
		heap.Push(&pickHeap, worker{costs[i], i})
	}

	var result int64 = 0
	for i := 0; i < k; i++ {
		item := heap.Pop(&pickHeap).(worker)
		result += int64(item.cost)

		if l > r {
			continue
		}

		if item.index < l {
			heap.Push(&pickHeap, worker{costs[l], l})
			l++
		} else {
			heap.Push(&pickHeap, worker{costs[r], r})
			r--
		}
	}
	return result
}
