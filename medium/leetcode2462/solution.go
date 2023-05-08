package main

import (
	"container/heap"
	"sort"
)

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

// 用不着Push和Pop，随便实现一下
type lhp struct{ sort.IntSlice }

func (lhp) Push(interface{})     {}
func (lhp) Pop() (_ interface{}) { return }

func totalCost1(costs []int, k int, candidates int) int64 {
	var ans int64 = 0
	// 在原始数据上进行建堆
	n := len(costs)
	if candidates*2 < n {
		// 比n小就建两个堆
		pre := lhp{costs[:candidates]}
		suf := lhp{costs[n-candidates:]}
		heap.Init(&pre)
		heap.Init(&suf)
		for i, j := candidates, n-candidates-1; i <= j && k > 0; k-- {
			if pre.IntSlice[0] <= suf.IntSlice[0] {
				// 选前面的
				ans += int64(pre.IntSlice[0])
				// 将中间的赋值给前面的
				pre.IntSlice[0] = costs[i]
				heap.Fix(&pre, 0)
				i++
			} else {
				ans += int64(suf.IntSlice[0])
				suf.IntSlice[0] = costs[j]
				heap.Fix(&suf, 0)
				j--
			}
		}
		if k == 0 {
			return ans
		}
		// 取了几个后，重叠了，剩下的合并成一个数组
		costs = append(pre.IntSlice, suf.IntSlice...)
	}
	sort.Ints(costs)
	for _, c := range costs[:k] {
		ans += int64(c)
	}
	return ans
}
