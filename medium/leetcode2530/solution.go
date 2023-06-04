package main

import (
	"container/heap"
	"sort"
)

type bhp struct{ sort.IntSlice }

func (l *bhp) Less(i, j int) bool { return l.IntSlice.Less(j, i) }
func (l *bhp) Push(x interface{}) { l.IntSlice = append(l.IntSlice, x.(int)) }
func (l *bhp) Pop() interface{} {
	n := len(l.IntSlice)
	x := l.IntSlice[n-1]
	l.IntSlice = l.IntSlice[:n-1]
	return x
}

func maxKelements(nums []int, k int) int64 {
	bigHeap := bhp{nums}
	heap.Init(&bigHeap)
	var result int64 = 0
	for i := 0; i < k; i++ {
		v := heap.Pop(&bigHeap).(int)
		result += int64(v)
		heap.Push(&bigHeap, (v+2)/3)
	}
	return result
}
